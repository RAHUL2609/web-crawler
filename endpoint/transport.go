package endpoint

import (
	"context"
	"example.com/web-crawler-golang/middleware"
	"example.com/web-crawler-golang/serveroption"
	service2 "example.com/web-crawler-golang/service"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	kitlogrus "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	"example.com/web-crawler-golang/domain"
	mhttp "example.com/web-crawler-golang/http"
)

// SubmitTask returns a handler for the submitting the task.
func SubmitTask (
	s service2.CrawlerServiceInterface,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
	}

	/*mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
		makeValidationMiddleware(),
	)*/

	mw := endpoint.Chain(
		makeValidationMiddleware(),
	)
	endpointHandler := kithttp.NewServer(
		mw(submitTask(s)),
		decodeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

// MakeGetReadingsHandler returns a handler for the Readings service.
func GetCrawledUrlsById (
	s service2.CrawlerServiceInterface,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		makeValidationUid(),
	)

	endpointHandler := kithttp.NewServer(
		mw(getCrawledUrlsById(s)),
		decodeSmartMeterIdFromRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.FetchUrl
	err := mhttp.DecodeRequest(ctx, r, &request)
	return request, err
}

func decodeSmartMeterIdFromRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return strings.Split(r.URL.Path, "/")[3], nil
}
