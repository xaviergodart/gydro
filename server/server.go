package server

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	Proxies map[string]*httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	if reverse, err := p.Proxies[r.URL.Path]; err == false {
		log.Println("proxy: custom")
		reverse.ServeHTTP(w, r)
		return
	}

	log.Println("proxy: default")
	http.NotFoundHandler().ServeHTTP(w, r)
}
