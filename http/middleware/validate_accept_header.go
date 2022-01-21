package middleware

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"example.com/web-crawler-golang/http"
	"example.com/web-crawler-golang/http/contextkeys"
)

// MakeAcceptHeaderValidationMiddleware returns a middleware to validate a request's accept header.
func MakeAcceptHeaderValidationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			acceptHeaderValue := ctx.Value(contextkeys.AcceptHeader).(string)
			if acceptHeaderValue != http.MimeJSON {
				return nil, fmt.Errorf("%w: '%s' not allowed", http.ErrInvalidAcceptHeader, acceptHeaderValue)
			}
			return next(ctx, request)
		}
	}
}
