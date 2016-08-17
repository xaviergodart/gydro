package middlewares

import (
	"github.com/xaviergodart/gydro/httpapi"
	"encoding/json"
	"net/http"
)

func KeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyget := r.URL.Query().Get("apikey")
		keyheader := r.Header.Get("apikey")
		if keyget == "" && keyheader == "" {
			httpapi.NewHttpError(w, "ErrorApiKeyMandatory")
			return
		}
		next.ServeHTTP(w, r)
	})
}
