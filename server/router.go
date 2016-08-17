package server

import (
	"github.com/bmizerany/pat"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

type Router struct {
	mux *pat.PatternServeMux
}

func NewRouter(apis []*models.Api) *Router {
	// Loading routing and backend configuration
	mux := pat.New()
	for _, api := range apis {
		mux.Get(api.Route, NewReverseProxy(api.Backends))
	}

	return &Router{mux: mux}
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.mux.ServeHTTP(w, r)
}
