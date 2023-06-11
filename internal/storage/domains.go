package storage

import (
	"sync"
	"time"
	"url-accessibility-checker/internal/network"
)

type (
	Domain struct {
		Url      string
		PingTime time.Duration
		Status   int
	}

	DomainsStorage struct {
		availableDomains []*Domain
		domains          []*Domain
		updateInterval   time.Duration
	}
)

var (
	storage    *DomainsStorage
	domainUrls = []string{
		"google.com", "youtube.com", "facebook.com", "baidu.com", "wikipedia.org", "qq.com", "taobao.com", "yahoo.com",
		"tmall.com", "amazon.com", "google.co.in", "twitter.com", "sohu.com", "jd.com", "live.com", "instagram.com",
		"sina.com.cn", "weibo.com", "google.co.jp", "reddit.com", "vk.com", "360.cn", "login.tmall.com", "blogspot.com",
		"yandex.ru", "google.com.hk", "netflix.com", "linkedin.com", "pornhub.com", "google.com.br", "twitch.tv",
		"pages.tmall.com", "csdn.net", "yahoo.co.jp", "mail.ru", "aliexpress.com", "alipay.com", "office.com",
		"google.fr", "google.ru", "google.co.uk", "microsoftonline.com", "google.de", "ebay.com", "microsoft.com",
		"livejasmin.com", "t.co", "bing.com", "xvideos.com", "google.ca",
	}
)

func newDomainsStorage() *DomainsStorage {
	domains := make([]*Domain, len(domainUrls))
	for i, domainUrl := range domainUrls {
		domains[i] = &Domain{
			Url: domainUrl,
		}
	}

	return &DomainsStorage{
		domains:        domains,
		updateInterval: time.Second * 30,
	}
}

func (s *DomainsStorage) updateDomains() {
	wg := sync.WaitGroup{}

	for _, d := range s.domains {
		d := d
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp := network.PingUrl(d.Url)
			d.PingTime = resp.Ping
			d.Status = resp.Status

			if resp.Status == 200 {
				s.availableDomains = append(s.availableDomains, d)
			}
		}()
	}
	wg.Wait()
}

func (s *DomainsStorage) Watch() {
	go func() {
		for {
			s.updateDomains()
			time.Sleep(s.updateInterval)
		}
	}()
}

func (s *DomainsStorage) GetAvailableDomains() []*Domain {
	return s.availableDomains
}

func (s *DomainsStorage) GetDomains() []*Domain {
	return s.domains
}

func GetStorage() *DomainsStorage {
	if storage == nil {
		storage = newDomainsStorage()
	}

	return storage
}
