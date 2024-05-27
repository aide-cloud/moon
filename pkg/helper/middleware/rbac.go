package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"

	"github.com/aide-cloud/moon/api/merr"
)

type CheckRbacFun func(ctx context.Context, operation string) (bool, error)

func Rbac(check CheckRbacFun) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			operation := GetOperation(ctx)
			// 判断该用户在该资源是否有权限
			has, err := check(ctx, operation)
			if err != nil {
				return nil, merr.ErrorNotification("权限校验失败")
			}
			if !has {
				return nil, merr.ErrorModal("请联系管理员分配权限")
			}

			return handler(ctx, req)
		}
	}
}