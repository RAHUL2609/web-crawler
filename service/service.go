package service

import (
	"example.com/web-crawler-golang/domain"
	"example.com/web-crawler-golang/repository"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
	"sync"
)

type CrawlerServiceInterface interface {
	GenerateCrawlerRequestId(url string) (string, error)
	FetchCrawledUrlById(uid string) ([]string, error)
}

type CrawlerService struct {
	lock sync.Mutex
	task map[string]*requestParser
}

type requestParser struct {
	childUrls 		map[string]bool
	processed 		*bool
	requestUrl 		string
	requestHostName string
	status			Status
}

type Status int

const (
	INITIATED Status = iota
	PROCESSING
	FAILED
	COMPLETED
)

func NewCrawlerService(
) CrawlerServiceInterface {
	return CrawlerService {
		lock: sync.Mutex{},
		task: make(map[string]*requestParser),
	}
}

func (cs CrawlerService) GenerateCrawlerRequestId(uri string) (string, error) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	uid := uuid.New().String()

	// Fetch host Name
	URL, err :=url.Parse(uri)
	if err != nil {
		log.Errorf("error while parse url : %s", uri)
		return "", domain.ErrNotProcessed
	}

	cs.task[uid] = &requestParser{
		childUrls:  make(map[string]bool),
		requestUrl: uri,
		requestHostName: URL.Hostname(),
		status:    INITIATED,
	}
	go cs.processCrawling(uid)
	return uid, nil
}

func (cs CrawlerService) FetchCrawledUrlById(uid string) ([]string, error) {
	if task, ok := cs.task[uid]; ok {
		if task.status == COMPLETED {
			return covertMapToArray(task.childUrls), nil
		} else {
			return nil, domain.ErrNotProcessed
		}
	}
	return nil, domain.ErrNotFound
}

func (cs CrawlerService) processCrawling(uid string) {
	if reqToProcess, ok := cs.task[uid]; ok && reqToProcess.status == INITIATED {
		reqToProcess.status = PROCESSING
		reqToProcess.Crawler(reqToProcess.requestUrl, repository.Respository{})
		reqToProcess.status = COMPLETED
	}
}

func (parser requestParser) Crawler(parentUrl string, provider repository.RespositoryInterface) {
	for _, url := range provider.CrawlUrl(parentUrl) {
		if _, ok := parser.childUrls[url]; !ok && parser.isSameDomain(url) {
			parser.childUrls[url] = true
			parser.Crawler(url, provider)
		}
	}
}

func (parser requestParser) isSameDomain(uri string) bool {
	rawUrl, err := url.Parse(uri)
	if err != nil {
		log.Errorf("error while parsing url : %s", uri)
		parser.status = FAILED
		return false
	}
	if strings.EqualFold(rawUrl.Hostname(), parser.requestHostName) {
		return true
	}
	return false
}


func covertMapToArray(mp map[string]bool) []string {
	derivedUrls := make([]string, 0)
	for key, _ := range mp {
		derivedUrls = append(derivedUrls, key)
	}
	return derivedUrls
}

