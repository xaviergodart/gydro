package server

import (
	"log"
	"net/url"
	"net/http"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/stream"
	"github.com/xaviergodart/gydro/models"
)

type Server struct {
	EntryPoints map[string]*stream.Streamer
}

func NewServer() *Server {
	apis := models.FindAllApis()
	conf := make(map[string]*stream.Streamer)

	// Loading routing and backend configuration from database
	for _, api := range apis {
		fwd, _ := forward.New()
		lb, _ := roundrobin.New(fwd)
		for _, backend := range api.Backends {
			target, _ := url.Parse(backend)
			lb.UpsertServer(target)
		}
		stream, _ := stream.New(lb, stream.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
		conf[api.Route] = stream
	}

	return &Server{EntryPoints: conf}
}

func (s *Server) ListenAndServe(addr string) {
	http.Handle("/", s)
	http.ListenAndServe(addr, nil)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GydroProxy", "GydroProxy")
	log.Println(r.URL.Path)
	if backends, ok := s.EntryPoints[r.URL.Path]; ok {
		log.Println("proxy: custom")
		log.Println(ok)
		backends.ServeHTTP(w, r)
		return
	}

	log.Println("proxy: default")
	http.NotFoundHandler().ServeHTTP(w, r)
}
