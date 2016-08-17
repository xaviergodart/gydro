package server

import (
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	Proxies map[string]*httputil.ReverseProxy
}

func NewProxy() *Proxy {
	conf := make(map[string]*httputil.ReverseProxy)
	target, _ := url.Parse("https://google.fr/")
	conf["/test"] = httputil.NewSingleHostReverseProxy(target)

	return &Proxy{Proxies: conf}
}

func (p *Proxy) ListenAndServe(addr string) {
	http.Handle("/", p)
	http.ListenAndServe(addr, nil)
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	if reverse, ok := p.Proxies[r.URL.Path]; ok {
		log.Println("proxy: custom")
		log.Println(ok)
		reverse.ServeHTTP(w, r)
		return
	}

	log.Println("proxy: default")
	http.NotFoundHandler().ServeHTTP(w, r)
}
