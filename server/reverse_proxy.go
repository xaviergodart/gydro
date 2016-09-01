package server

import (
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/stream"
	"github.com/vulcand/oxy/cbreaker"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	stream *stream.Streamer
}

func NewReverseProxy(backends []string) *ReverseProxy {
	fwd, _ := forward.New()
	cb, _ := cbreaker.New(fwd, `NetworkErrorRatio() > 0.5`)
	lb, _ := roundrobin.New(cb)
	for _, backend := range backends {
		target, _ := url.Parse(backend)
		lb.UpsertServer(target)
	}
	stream, _ := stream.New(lb, stream.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
	return &ReverseProxy{stream: stream}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.stream.ServeHTTP(w, r)
}
