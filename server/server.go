package server

import (
	"net/http"
	"github.com/xaviergodart/gydro/middlewares"
	"github.com/xaviergodart/gydro/models"
)

type Server struct {
	ReverseProxy *middlewares.ReverseProxy
}

func NewServer() *Server {
	apis := models.FindAllApis()
	reverseProxy := middlewares.NewReverseProxy(apis)
	return &Server{ReverseProxy: reverseProxy}
}

func (s *Server) ListenAndServe(addr string) {
	http.Handle("/", s.ReverseProxy)
	http.ListenAndServe(addr, nil)
}

