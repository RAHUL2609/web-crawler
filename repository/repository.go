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

var ScannedUrls domain.UrlTemplate

func init() {
	urlTemplateJson, err := os.Open("resources/urls.json")
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

func CrawlUrl(url string) []string {
	if derivedUrls , ok := ScannedUrls.UrlMap[url]; ok {
		return derivedUrls
	}
	return nil
}
