package middleware

import (
	"context"

	"github.com/aide-family/moon/pkg/util/types"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// SourceType 获取请求头中的Source-Type  sourceType System Team
func SourceType() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				sourceCode := tr.RequestHeader().Get("Source-Type")
				if types.TextIsNull(sourceCode) {
					sourceCode = "Team"
				}
				source := &SourceTypeInfo{
					SourceCode: sourceCode,
				}
				ctx = context.WithValue(ctx, SourceTypeInfo{}, source)
			}
			return handler(ctx, req)
		}
	}
}

// SourceTypeInfo Request header source
type SourceTypeInfo struct {
	SourceCode string
}

func (s *SourceTypeInfo) SetSourceType(sourceCode string) {
	s.SourceCode = sourceCode
}

func (s *SourceTypeInfo) GetSourceCode() string {
	if types.IsNil(s.SourceCode) {
		return ""
	}
	return s.SourceCode
}

func ParseSourceTypeInfo(ctx context.Context) (*SourceTypeInfo, bool) {
	sourceTypeInfo, ok := ctx.Value(SourceTypeInfo{}).(*SourceTypeInfo)
	if ok {
		return sourceTypeInfo, true
	}
	return nil, false
}
