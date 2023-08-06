package interceptor

import (
	"context"

	"github.com/bufbuild/connect-go"
)

type validator interface {
	Validate() error
}

func NewValidationInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// do nothing
				return next(ctx, req)
			}

			msg := req.Any()
			validator, ok := msg.(validator)
			if !ok {
				return next(ctx, req)
			}

			if err := validator.Validate(); err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, err)
			}

			return next(ctx, req)
		}
	}
}
