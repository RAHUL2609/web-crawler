package endpoint

import (
	"context"
	"example.com/web-crawler-golang/domain"
	service2 "example.com/web-crawler-golang/service"
	"github.com/go-kit/kit/endpoint"
)

func submitTask(s service2.CrawlerServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqUrl := request.(domain.FetchUrl)
		uid, err := s.GenerateCrawlerRequestId(reqUrl.Url)
		if err != nil {
			return nil, err
		}
		return domain.TaskResponse{TaskId:uid}, nil
	}
}
func getCrawledUrlsById(s service2.CrawlerServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uid := request.(string)
		fetchedUrls, err := s.FetchCrawledUrlById(uid)
		if err != nil {
			return nil, err
		}
		return fetchedUrls, nil
	}
}
