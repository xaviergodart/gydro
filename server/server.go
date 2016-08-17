package server

import (
	"github.com/xaviergodart/gydro/middlewares"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

func ListenAndServe(addr string) {
	apis := models.FindAllApis()
	router := NewRouter(apis)
	http.Handle("/",
		middlewares.KeyAuth(
			router,
		))
	http.ListenAndServe(addr, nil)
}
