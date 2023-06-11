package network

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	timeout = time.Second * 5
)

type (
	Response struct {
		Ping   time.Duration
		Status int
	}
)

func PingUrl(url string) *Response {
	ch := make(chan *Response)

	go func() {
		defer close(ch)
		timer := time.Now()
		c := &fasthttp.Client{}

		url = fmt.Sprintf("https://%v", url)
		statusCode, _, err := c.Get(nil, url)

		ch <- newResponse(statusCode, time.Since(timer))

		if err != nil {
			fmt.Printf("Got error while ping, %v\n", err)
			return
		}
	}()

	select {
	case <-time.After(timeout):
		return newResponse(504, timeout)
	case response := <-ch:
		return response
	}
}

func newResponse(statusCode int, ping time.Duration) *Response {
	return &Response{
		Ping:   ping,
		Status: statusCode,
	}
}
