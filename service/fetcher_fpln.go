package service

import (
	"bytes"
	"stock_controller/glob"
	"stock_controller/model"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/xerrors"
)

var FplnController *FetcherFpln

type FetcherFpln struct {
	Code string
	Url  string
}

func NewFetcherFpln() *FetcherFpln {
	return &FetcherFpln{
		Code: glob.YamlC.FetcherfplnCfg.Code,
		Url:  glob.YamlC.FetcherfplnCfg.Url,
	}
}

func (f *FetcherFpln) Parse(bodyBytes []byte) (proxies []model.Proxy, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		err = xerrors.Errorf("error getting document from body: %w", err)
		return
	}

	doc.Find(".table-striped tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		item := model.Proxy{SourceCode: f.Code}

		if https := s.Find("td").Eq(6).Text(); https == "yes" {
			item.HTTPS = true
		} else {
			item.HTTPS = false
		}

		item.IPAddress = s.Find("td").Eq(0).Text()
		item.Port = s.Find("td").Eq(1).Text()

		if anonymous := s.Find("td").Eq(4).Text(); anonymous == "anonymous" {
			item.Anonymous = true
		} else {
			item.Anonymous = false
		}

		item.CountryCode = strings.TrimSpace(s.Find("td").Eq(2).Text())
		item.Country = strings.TrimSpace(s.Find("td").Eq(3).Text())

		if item.CountryCode != "--" && item.Country != "Unknown" {
			if glob.IsInt(item.Port) {
				proxies = append(proxies, item)
			}

		}
	})
	return
}
