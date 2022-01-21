package endpoint

import (
	"context"
	"example.com/web-crawler-golang/domain"
	"github.com/go-kit/kit/endpoint"
)

func makeValidationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			_, ok := req.(domain.FetchUrl)
			if !ok {
				return nil, domain.ErrInvalidMessageType
			}
			// Validate actual URl
			return next(ctx, req)
		}
	}
}

func makeValidationUid() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			_, ok := req.(string)
			if !ok {
				return nil, domain.ErrInvalidMessageType
			}
			// Validate actual URl
			return next(ctx, req)
		}
	}
}
