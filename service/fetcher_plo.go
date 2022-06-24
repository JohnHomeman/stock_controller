package service

import (
	"bytes"
	"stock_controller/glob"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/xerrors"
)

var PloController *FetcherPlo

type FetcherPlo struct {
	Code string
	Url  string
}

func NewFetcherPlo() *FetcherPlo {
	return &FetcherPlo{
		Code: glob.YamlC.FetcherploCfg.Code,
		Url:  glob.YamlC.FetcherploCfg.Url,
	}
}

func (f *FetcherPlo) Parse(bodyBytes []byte) (proxies []mongodb.Proxy, err error) {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		err = xerrors.Errorf("error getting document from body: %w", err)
		return
	}

	doc.Find(".table").Find("ul").Each(func(i int, s *goquery.Selection) {
		item := mongodb.Proxy{SourceCode: f.Code}

		temp := strings.Split(s.Find(".proxy").Find("script").Text(), "'")
		url := strings.Split(glob.DecodeBase64(temp[1]), ":")
		item.IPAddress = url[0]
		item.Port = url[1]

		if https := s.Find(".https").Text(); https == "HTTPS" {
			item.HTTPS = true
		} else {
			item.HTTPS = false
		}

		if anonymous := s.Find(".type").Text(); anonymous == "Elite" {
			item.Anonymous = true
		} else {
			item.Anonymous = false
		}

		temp = strings.Split(s.Find(".country-city").Find(".name").Text(), " ")
		item.CountryCode = strings.TrimSpace(temp[0])
		item.Country = strings.TrimSpace(temp[1])
		if item.CountryCode != "--" && item.Country != "Unknown" {
			if glob.IsInt(item.Port) {
				proxies = append(proxies, item)
			}

		}
	})
	return
}
