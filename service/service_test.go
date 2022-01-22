package service

import (
	"example.com/web-crawler-golang/domain"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCrawlerService_GenerateCrawlerRequestId(t *testing.T) {
		t.Run("Success Case", func(t *testing.T) {
			cs := CrawlerService {
				lock: sync.Mutex{},
				task: make(map[string]*requestParser),
			}
			url := "https//dummy.com/abc"
			got, err := cs.GenerateCrawlerRequestId(url)
			if err == nil {
				assert.Equal(t, cs.task[got].status, INITIATED)
			} else {
				t.Errorf("Got error while making GenerateCrawlerRequestId, %v", err)
			}
		})
}

func TestCrawlerService_FetchCrawledUrlById(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		cs := CrawlerService {
			lock: sync.Mutex{},
			task: make(map[string]*requestParser),
		}

		cs.task["1234567"] = &requestParser{status:COMPLETED, childUrls: map[string]bool{
			"https//dummy.com/tiger":true,
		}}
		got, err := cs.FetchCrawledUrlById("1234567")
		if err == nil {
			assert.Equal(t, got[0], "https//dummy.com/tiger")
		} else {
			t.Errorf("Got error while making GenerateCrawlerRequestId, %v", err)
		}
	})

	t.Run("Failure Case", func(t *testing.T) {
		cs := CrawlerService {
			lock: sync.Mutex{},
			task: make(map[string]*requestParser),
		}

		cs.task["1234567"] = &requestParser{status:FAILED}
		_, err := cs.FetchCrawledUrlById("1234567")
		if err != nil {
		 	assert.Equal(t, domain.ErrNotProcessed, err)
		}
	})
}

func TestNewCrawlerService_Crawler(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		reqParser := requestParser{status:COMPLETED, childUrls: map[string]bool{}, requestUrl:"https//dummy.com"}
		reqParser.Crawler("https//dummy.com", MockRespository{})
		assert.Contains(t, reqParser.childUrls, "https//dummy.com/tiger")
	})
}


type MockRespository struct {
}

func(mr MockRespository) CrawlUrl(url string) []string {
	return []string{"https//dummy.com/tiger"}
}

func Contains(arr []string, str string) bool {
	for _,val := range arr {
		if val == str {
			return true
		}
	}
	return false
}