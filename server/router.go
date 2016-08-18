package server

import (
	router "github.com/gorilla/mux"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

type Router struct {
	mux *router.Router
}

func NewRouter(apis []*models.Api) *Router {
	// Loading routing and backend configuration
	mux := router.NewRouter()
	for _, api := range apis {
		mux.PathPrefix(api.Route).Handler(NewReverseProxy(api.Backends))
	}
	return &Router{mux: mux}
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.mux.ServeHTTP(w, r)
}
