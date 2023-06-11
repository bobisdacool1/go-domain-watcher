package stats

import (
	"github.com/mssola/useragent"
	"time"
)

type (
	Endpoint string

	Row struct {
		requestTime time.Time
		userAgent   *useragent.UserAgent
		endpoint    Endpoint
	}

	Service struct {
		stats map[Endpoint][]*Row
	}
)

func (s *Service) AddStatisticsRow(endpoint Endpoint, row *Row) {
	s.stats[endpoint] = append(s.stats[endpoint], row)
}

func NewRow(endpoint Endpoint, ua *useragent.UserAgent) *Row {
	return &Row{
		requestTime: time.Now(),
		userAgent:   ua,
		endpoint:    endpoint,
	}
}

func NewService() *Service {
	return &Service{
		stats: make(map[Endpoint][]*Row),
	}
}

func (s *Service) GetByEndpoint(endpoint Endpoint) []*Row {
	return s.stats[endpoint]
}
