package server

import (
	"github.com/bmizerany/pat"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/stream"
	"github.com/xaviergodart/gydro/models"
	"log"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	mux *pat.PatternServeMux
}

func NewReverseProxy(apis []*models.Api) *ReverseProxy {
	// Loading routing and backend configuration
	mux := pat.New()
	for _, api := range apis {
		fwd, _ := forward.New()
		lb, _ := roundrobin.New(fwd)
		for _, backend := range api.Backends {
			target, _ := url.Parse(backend)
			lb.UpsertServer(target)
		}
		stream, _ := stream.New(lb, stream.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
		mux.Get(api.Route, stream)

	}

	return &ReverseProxy{mux: mux}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	rp.mux.ServeHTTP(w, r)
}
