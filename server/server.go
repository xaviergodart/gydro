package server

import (
	"github.com/xaviergodart/gydro/middlewares"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

func ListenAndServe(addr string) {
	apis := models.FindAllApis()
	reverseProxy := NewReverseProxy(apis)
	http.Handle("/",
		middlewares.KeyAuth(
			reverseProxy,
		))
	http.ListenAndServe(addr, nil)
}
