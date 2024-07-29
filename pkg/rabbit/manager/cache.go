package manager

import (
	"context"
	"fmt"
	"github.com/aide-family/moon/pkg/client/rpc"
	"github.com/aide-family/moon/pkg/runtime/watch"
	"k8s.io/apimachinery/pkg/conversion"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	mcache "github.com/aide-family/moon/pkg/runtime/cache"
	"github.com/aide-family/moon/pkg/runtime/client"

	"k8s.io/client-go/tools/cache"
)

type innerCache struct {
	cacheMap *innerCacheMap
	scheme   *runtime.Scheme
}

func newInnerCache(cli *rpc.RESTClient, scheme *runtime.Scheme, resync time.Duration) *innerCache {
	return &innerCache{
		cacheMap: newInnerCacheMap(cli, scheme, resync),
		scheme:   scheme,
	}
}

func (m *innerCache) Get(context context.Context, key string, out client.Object) error {
	kind, err := m.scheme.ObjectKind(out)
	if err != nil {
		return err
	}

	started, cacheEntry, err := m.cacheMap.Get(context, kind, out)
	if err != nil {
		return err
	}

	if !started {
		return fmt.Errorf("cache not started, %s", kind)
	}
	return cacheEntry.Reader.Get(context, key, out)
}

func (m *innerCache) List(context context.Context, out client.ObjectList, opts *api.ListOptions) error {
	kind, err := m.scheme.ObjectKind(out)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(kind, "List") {
		return fmt.Errorf("list must be suffix with List, %s", kind)
	}
	kind = kind[:len(kind)-4]
	started, cacheEntry, err := m.cacheMap.Get(context, kind, out)
	if err != nil {
		return err
	}

	if !started {
		return fmt.Errorf("cache not started, %s", kind)
	}
	return cacheEntry.Reader.List(context, out, opts)
}

func (m *innerCache) WaitForCacheSync(ctx context.Context) bool {
	if !m.cacheMap.waitForStarted(ctx) {
		return false
	}
	return cache.WaitForCacheSync(ctx.Done(), m.cacheMap.HasSyncedFuncs()...)
}

type innerCacheMap struct {
	cli *rpc.RESTClient

	scheme *runtime.Scheme

	kindCacheMap map[string]*MapEntry

	resync time.Duration

	// mu guards access to the map
	mu sync.RWMutex

	ctx context.Context
	// start is true if the informers have been started
	started bool

	startWait chan struct{}
}

func newInnerCacheMap(cli *rpc.RESTClient, scheme *runtime.Scheme, resync time.Duration) *innerCacheMap {
	return &innerCacheMap{
		cli:          cli,
		scheme:       scheme,
		kindCacheMap: make(map[string]*MapEntry),
		resync:       0,
		mu:           sync.RWMutex{},
		startWait:    make(chan struct{}),
	}
}

func (ic *innerCacheMap) Start(ctx context.Context) {
	func() {
		ic.mu.Lock()
		defer ic.mu.Unlock()

		// Set the stop channel so it can be passed to informers that are added later
		ic.ctx = ctx

		for _, cacheEntry := range ic.kindCacheMap {
			go cacheEntry.Informer.Run(ctx)
		}

		// Set started to true so we immediately start any informers added later.
		ic.started = true
		close(ic.startWait)
	}()
	<-ctx.Done()
}

func (ic *innerCacheMap) HasSyncedFuncs() []cache.InformerSynced {
	ic.mu.RLock()
	defer ic.mu.RUnlock()
	syncedFuncs := make([]cache.InformerSynced, 0, len(ic.kindCacheMap))
	for _, informer := range ic.kindCacheMap {
		syncedFuncs = append(syncedFuncs, informer.Informer.HasSynced)
	}
	return syncedFuncs
}

func (ic *innerCacheMap) waitForStarted(ctx context.Context) bool {
	select {
	case <-ic.startWait:
		return true
	case <-ctx.Done():
		return false
	}
}

func (ic *innerCacheMap) Get(ctx context.Context, gvk string, obj runtime.Object) (bool, *MapEntry, error) {
	// Return the informer if it is found
	entry, started, ok := func() (*MapEntry, bool, bool) {
		ic.mu.RLock()
		defer ic.mu.RUnlock()
		entry, ok := ic.kindCacheMap[gvk]
		return entry, ic.started, ok
	}()

	if !ok {
		var err error
		if entry, started, err = ic.addInformerToMap(gvk, obj); err != nil {
			return started, nil, err
		}
	}

	if started && !entry.Informer.HasSynced() {
		// Wait for it to sync before returning the Informer so that folks don't read from a stale cache.
		if !cache.WaitForCacheSync(ctx.Done(), entry.Informer.HasSynced) {
			return started, nil, fmt.Errorf("failed waiting for %T Informer to sync", obj)
		}
	}

	return started, entry, nil
}

func (ic *innerCacheMap) addInformerToMap(kind string, obj runtime.Object) (*MapEntry, bool, error) {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	if entry, ok := ic.kindCacheMap[kind]; ok {
		return entry, ic.started, nil
	}

	lw := NewListWatchWithClient(ic.cli, kind, ic.scheme, &api.ListOptions{})
	indexer := cache.NewIndexer(KeyFunc, nil)
	entry := &MapEntry{
		Informer: newSharedIndexInformer(lw, obj, ic.resync, indexer),
		Reader:   Reader{indexer: indexer, kind: kind},
	}
	ic.kindCacheMap[kind] = entry

	if ic.started {
		go entry.Informer.Run(ic.ctx)
	}
	return entry, ic.started, nil
}

func KeyFunc(obj interface{}) (string, error) {
	data := obj.(client.Object)
	return data.GetName(), nil
}

type MapEntry struct {
	Informer *sharedIndexInformer
	Reader   Reader
}

type sharedIndexInformer struct {
	lw         mcache.ListerWatcher
	reSync     time.Duration
	objectType runtime.Object
	indexer    cache.Indexer
	sync       bool
}

func newSharedIndexInformer(lw mcache.ListerWatcher, object runtime.Object, resync time.Duration, indexer cache.Indexer) *sharedIndexInformer {
	return &sharedIndexInformer{
		lw:         lw,
		indexer:    indexer,
		reSync:     resync,
		objectType: object,
	}
}

func (s *sharedIndexInformer) Run(ctx context.Context) {
	s.reBuildCacheData()
	watchCh, err := s.lw.Watch(&api.ListOptions{})
	if err != nil {
		return
	}

	s.sync = true

	if s.reSync != 0 {
		ticker := time.NewTicker(s.reSync)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				watchCh.Stop()
				return
			case <-ticker.C:
				s.reBuildCacheData()
			case data := <-watchCh.ResultChan():
				switch data.Type {
				case watch.Added:
					err = s.indexer.Add(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				case watch.Deleted:
					err = s.indexer.Delete(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				case watch.Modified:
					err = s.indexer.Update(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				default:
					// TODO handler unknown type
				}
			}
		}
	} else {
		for {
			select {
			case <-ctx.Done():
				watchCh.Stop()
				return
			case data := <-watchCh.ResultChan():
				switch data.Type {
				case watch.Added:
					err = s.indexer.Add(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				case watch.Deleted:
					err = s.indexer.Delete(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				case watch.Modified:
					err = s.indexer.Update(data.Object)
					if err != nil {
						// TODO handler error
						continue
					}
				default:
					// TODO handler unknown type
				}
			}
		}
	}
}

func (s *sharedIndexInformer) reBuildCacheData() {
	list, err := s.lw.List(&api.ListOptions{})
	if err != nil {
		return
	}
	val := reflect.ValueOf(list)
	// Find the Items field
	field := val.FieldByName("Items")
	if !field.IsValid() {
		return
	}

	// Check if the Items field is an array or slice
	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		// Iterate over an array or slice
		data := make([]any, 0, field.Len())
		for i := 0; i < field.Len(); i++ {
			data[i] = field.Index(i).Interface()
		}
		err = s.indexer.Replace(data, "")
		if err != nil {
			return
		}
	}
}

func (s *sharedIndexInformer) HasSynced() bool {
	return s.sync
}

var _ client.Reader = &Reader{}

type Reader struct {
	indexer cache.Indexer
	kind    string
}

func (c *Reader) Get(_ context.Context, key string, out client.Object) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("object %s not found, kind: %s", key, c.kind)
	}

	if _, isObj := obj.(runtime.Object); !isObj {
		return fmt.Errorf("cache contained %T, which is not an Object", obj)
	}
	// 进行深度复制，以避免改变缓存
	obj = obj.(runtime.Object).DeepCopyObject()

	outVal := reflect.ValueOf(out)
	objVal := reflect.ValueOf(obj)
	// 判断输出的对象类型是否相同
	if !objVal.Type().AssignableTo(outVal.Type()) {
		return fmt.Errorf("cache had type %s, but %s was asked for", objVal.Type(), outVal.Type())
	}
	reflect.Indirect(outVal).Set(reflect.Indirect(objVal))
	out.GetObjectKind().SetKind(c.kind)
	return nil
}

func (c *Reader) List(_ context.Context, out client.ObjectList, opts *api.ListOptions) error {
	objs := c.indexer.List()
	runtimeObjs := make([]runtime.Object, 0, len(objs))
	for _, item := range objs {
		obj, isObj := item.(runtime.Object)
		if !isObj {
			return fmt.Errorf("cache contained %T, which is not an Object", obj)
		}
		outObj := obj.DeepCopyObject()
		outObj.GetObjectKind().SetKind(c.kind)
		runtimeObjs = append(runtimeObjs, outObj)
	}
	return SetList(out, runtimeObjs)
}

var objectSliceType = reflect.TypeOf([]runtime.Object{})

// SetList sets the given list object's Items member have the elements given in
// objects.
// Returns an error if list is not a List type (does not have an Items member),
// or if any of the objects are not of the right type.
func SetList(list runtime.Object, objects []runtime.Object) error {
	itemsPtr, err := getItemsPtr(list)
	if err != nil {
		return err
	}
	items, err := conversion.EnforcePtr(itemsPtr)
	if err != nil {
		return err
	}
	if items.Type() == objectSliceType {
		items.Set(reflect.ValueOf(objects))
		return nil
	}
	slice := reflect.MakeSlice(items.Type(), len(objects), len(objects))
	for i := range objects {
		dest := slice.Index(i)

		// check to see if you're directly assignable
		if reflect.TypeOf(objects[i]).AssignableTo(dest.Type()) {
			dest.Set(reflect.ValueOf(objects[i]))
			continue
		}

		src, err := conversion.EnforcePtr(objects[i])
		if err != nil {
			return err
		}
		if src.Type().AssignableTo(dest.Type()) {
			dest.Set(src)
		} else if src.Type().ConvertibleTo(dest.Type()) {
			dest.Set(src.Convert(dest.Type()))
		} else {
			return fmt.Errorf("item[%d]: can't assign or convert %v into %v", i, src.Type(), dest.Type())
		}
	}
	items.Set(slice)
	return nil
}

// getItemsPtr returns a pointer to the list object's Items member or an error.
func getItemsPtr(list runtime.Object) (interface{}, error) {
	v, err := conversion.EnforcePtr(list)
	if err != nil {
		return nil, err
	}

	items := v.FieldByName("Items")
	if !items.IsValid() {
		return nil, fmt.Errorf("missing Items field in list object")
	}
	switch items.Kind() {
	case reflect.Interface, reflect.Ptr:
		target := reflect.TypeOf(items.Interface()).Elem()
		if target.Kind() != reflect.Slice {
			return nil, fmt.Errorf("items field in list object is not a pointer to a slice")
		}
		return items.Interface(), nil
	case reflect.Slice:
		return items.Addr().Interface(), nil
	default:
		return nil, fmt.Errorf("items field in list object is not a pointer to a slice")
	}
}
