package server

import (
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/stream"
	"log"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	stream *stream.Streamer
}

func NewReverseProxy(backends []string) *ReverseProxy {
	fwd, _ := forward.New()
	lb, _ := roundrobin.New(fwd)
	for _, backend := range backends {
		target, _ := url.Parse(backend)
		lb.UpsertServer(target)
	}
	stream, _ := stream.New(lb, stream.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
	return &ReverseProxy{stream: stream}
}

// TODO : Add rewrite handler like rw, _:= rewrite.New(fwd)

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	r.URL.Path = "/a/data/"
	r2, _ := http.NewRequest(r.Method, r.URL.String(), r.Body)
	log.Println(r2.URL.Path)
	rp.stream.ServeHTTP(w, r2)
}
