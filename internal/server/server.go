package server

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/mssola/useragent"
	"github.com/valyala/fasthttp"
	"go-domain-watcher/internal/stats"
	domainStorage "go-domain-watcher/internal/storage"
)

type (
	Server struct {
		router *fasthttprouter.Router
	}
)

func NewServer() *Server {
	domainStorage.GetStorage().Watch()

	return &Server{
		router: fasthttprouter.New(),
	}
}

func (s *Server) Listen() error {
	router := fasthttprouter.New()

	router.GET(string(specialDomainEndpoint), HandleSpecial)
	router.GET(string(minPingEndpoint), HandleMinPing)
	router.GET(string(maxPingEndpoint), HandleMaxPing)

	router.GET(string(statsEndpoint), HandleStats)

	err := fasthttp.ListenAndServe(":54", router.Handler)
	if err != nil {
		return err
	}

	return nil
}

func updateStats(endpoint stats.Endpoint, ctx *fasthttp.RequestCtx) {
	if statsService == nil {
		statsService = stats.NewService()
	}

	ua := useragent.New(string(ctx.Request.Header.UserAgent()))

	statsService.AddStatisticsRow(endpoint, stats.NewRow(endpoint, ua))
}
