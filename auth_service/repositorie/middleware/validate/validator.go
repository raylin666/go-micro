package validate

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/raylin666/go-micro-protoc/errors"
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
					return nil, errors.ErrorDataValidateError(err.Error())
				}
			}
			return handler(ctx, req)
		}
	}
}
