package endpoint

import (
	"context"
	"example.com/web-crawler-golang/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CrawlerServiceMock struct {
}

var expectedUid string
var expectedCrawledUrls []string
var expectedErr error

func (c CrawlerServiceMock) GenerateCrawlerRequestId(url string) (string, error) {
	return expectedUid, expectedErr
}

func (c CrawlerServiceMock) FetchCrawledUrlById(uid string) ([]string, error)  {
	return expectedCrawledUrls, expectedErr
}

func Test_submitTask(t *testing.T){
	t.Run("Success Case", func(t *testing.T) {
		s := &CrawlerServiceMock{}
		expectedUid = "1234567"
		expectedErr = nil

		e := submitTask(s)

		request := domain.FetchUrl{Url: "http://dummy.com"}

		response, err := e(context.Background(), request)
		assert.NoError(t, err)

		expectedResponse := domain.TaskResponse{TaskId: expectedUid}
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Failure case", func(t *testing.T) {
		s := &CrawlerServiceMock{}
		expectedUid = "1234567"
		expectedErr = domain.ErrMissingArgument

		e := submitTask(s)

		request := domain.FetchUrl{Url: "http://dummy.com"}
		_, err := e(context.Background(), request)
		assert.Equal(t, expectedErr, err)
	})
}

func Test_getCrawledUrlsById(t *testing.T) {

	t.Run("Success Case", func(t *testing.T) {
		s := &CrawlerServiceMock{}
		expectedCrawledUrls = []string{"http://dummy.com/abc"}
		expectedErr = nil

		e := getCrawledUrlsById(s)

		uid := "1234567"

		response, err := e(context.Background(), uid)
		assert.NoError(t, err)

		expectedResponse := []string{"http://dummy.com/abc"}
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Failure Case", func(t *testing.T) {
		s := &CrawlerServiceMock{}
		expectedCrawledUrls = nil
		expectedErr = domain.ErrNotProcessed

		e := getCrawledUrlsById(s)

		uid := "1234567"

		_, err := e(context.Background(), uid)
		assert.Equal(t, expectedErr, err)
	})
}