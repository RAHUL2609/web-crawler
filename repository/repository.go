package repository

import (
	"encoding/json"
	"example.com/web-crawler-golang/domain"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type RespositoryInterface interface {
	CrawlUrl(url string) []string
}

type Respository struct {
}

var ScannedUrls domain.UrlTemplate

func InitLoadUrl() {
	urlTemplateJson, err := os.Open("urls.json")
	if err != nil {
		log.Fatalf("Error while loading json file")
	}
	defer urlTemplateJson.Close()
	jsonBytes, err := ioutil.ReadAll(urlTemplateJson)
	if err != nil {
		log.Fatalf("Error while reading parsing files from json")
	}
	err = json.Unmarshal(jsonBytes, &ScannedUrls)
	if err != nil {
		log.Fatalf("Error while unmarshalling")
	}
}

func (r Respository) CrawlUrl(url string) []string {
	if derivedUrls , ok := ScannedUrls.UrlMap[url]; ok {
		return derivedUrls
	}
	return nil
}
