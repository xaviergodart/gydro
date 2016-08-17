package server

import (
	"log"
	"math/rand"
	"net/url"
	"net/http"
	"net/http/httputil"
	"github.com/xaviergodart/gydro/models"
)

type Proxy struct {
	Proxies map[string][]*httputil.ReverseProxy
}

func NewProxy() *Proxy {
	apis := models.FindAllApis()
	conf := make(map[string][]*httputil.ReverseProxy)

	// Loading routing and backend configuration from database
	for _, api := range apis {
		var targets []*httputil.ReverseProxy
		for _, backend := range api.Backends {
			target, _ := url.Parse(backend)
			targets = append(targets, httputil.NewSingleHostReverseProxy(target))
		}
		conf[api.Route] = targets
	}

	return &Proxy{Proxies: conf}
}

func (p *Proxy) ListenAndServe(addr string) {
	http.Handle("/", p)
	http.ListenAndServe(addr, nil)
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	if backends, ok := p.Proxies[r.URL.Path]; ok {
		log.Println("proxy: custom")
		log.Println(ok)
		//random load balancing between backends
		backends[rand.Intn(len(backends))].ServeHTTP(w, r)
		return
	}

	log.Println("proxy: default")
	http.NotFoundHandler().ServeHTTP(w, r)
}
