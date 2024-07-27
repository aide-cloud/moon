package cache

import (
	"context"
	"fmt"
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/client"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/client-go/tools/cache"
	"reflect"
	"time"
)

type Options struct {
	Scheme *runtime.Scheme

	Resync *time.Duration
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
