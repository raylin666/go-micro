package validate

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	v1 "mt/api/v1"
)

type validator interface {
	Validate() error
}

// Validator is a validator middleware.
func Validator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					return nil, v1.ErrorDataValidateError(err.Error()).WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}
