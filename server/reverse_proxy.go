package server

import (
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/buffer"
	"github.com/vulcand/oxy/cbreaker"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	handler http.Handler
}

func NewReverseProxy(backends []string) *ReverseProxy {
	fwd, _ := forward.New()
	lb, _ := roundrobin.New(fwd)
	for _, backend := range backends {
		target, _ := url.Parse(backend)
		lb.UpsertServer(target)
	}
	buff, _ := buffer.New(lb, buffer.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
	cb, _ := cbreaker.New(buff, `NetworkErrorRatio() > 0.5`)
	return &ReverseProxy{handler: cb}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.handler.ServeHTTP(w, r)
}
