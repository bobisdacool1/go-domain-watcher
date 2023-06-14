package server

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"go-domain-watcher/internal/stats"
	domainStorage "go-domain-watcher/internal/storage"
)

type (
	Handler func(ctx *fasthttp.RequestCtx)
)

var (
	statsService *stats.Service

	endpoints = []stats.Endpoint{specialDomainEndpoint, minPingEndpoint, maxPingEndpoint, statsEndpoint}
)

const (
	specialDomainEndpoint stats.Endpoint = "/special/:domain"
	minPingEndpoint       stats.Endpoint = "/min"
	maxPingEndpoint       stats.Endpoint = "/max"
	statsEndpoint         stats.Endpoint = "/stats"
)

func HandleSpecial(ctx *fasthttp.RequestCtx) {
	defer updateStats(ctx, specialDomainEndpoint)

	storage := domainStorage.GetStorage()
	availableDomains := storage.GetAvailableDomains()
	domains := storage.GetDomains()

	url := fmt.Sprintf("%v", ctx.UserValue("domain"))
	if url == "" {
		fmt.Fprint(ctx, "No domain provided :(")
		return
	}

	for _, domain := range availableDomains {
		if domain.Url == url {
			fmt.Fprintf(ctx, "Domain found, url: %v, ping: %v", domain.Url, domain.PingTime)
			return
		}
	}

	for _, domain := range domains {
		if domain.Url == url {
			fmt.Fprintf(ctx, "Domain found, not working, url: %v, ping: %v, status: %v", domain.Url, domain.PingTime, domain.Status)
			return
		}
	}
	fmt.Fprint(ctx, "No domain found :(")
}

func HandleMinPing(ctx *fasthttp.RequestCtx) {
	defer updateStats(ctx, minPingEndpoint)

	storage := domainStorage.GetStorage()
	availableDomains := storage.GetAvailableDomains()
	minPingIndex := 0

	for i, domain := range availableDomains {
		if domain.PingTime < availableDomains[minPingIndex].PingTime {
			minPingIndex = i
		}
	}
	fmt.Fprintf(ctx, "Domain found, url: %v, ping: %v", availableDomains[minPingIndex].Url, availableDomains[minPingIndex].PingTime)
}

func HandleMaxPing(ctx *fasthttp.RequestCtx) {
	defer updateStats(ctx, maxPingEndpoint)

	storage := domainStorage.GetStorage()
	availableDomains := storage.GetAvailableDomains()
	maxPingIndex := 0

	for i, domain := range availableDomains {
		if domain.PingTime > availableDomains[maxPingIndex].PingTime {
			maxPingIndex = i
		}
	}
	fmt.Fprintf(ctx, "Domain found, url: %v, ping: %v", availableDomains[maxPingIndex].Url, availableDomains[maxPingIndex].PingTime)
}

func HandleStats(ctx *fasthttp.RequestCtx) {
	defer updateStats(ctx, statsEndpoint)
	if statsService == nil {
		fmt.Fprint(ctx, "No stats yet :(")
		return
	}

	for _, endpoint := range endpoints {
		rows := statsService.GetByEndpoint(endpoint)
		fmt.Fprintf(ctx, "Visits for %v:%v\n", endpoint, len(rows))
	}

}
