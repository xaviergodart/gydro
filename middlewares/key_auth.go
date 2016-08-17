package middlewares

import (
	"github.com/xaviergodart/gydro/errors"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

func KeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyget := r.URL.Query().Get("apikey")
		keyheader := r.Header.Get("apikey")
		switch {
			case keyget == "" && keyheader == "":
				errors.NewHttpError(w, "ErrorApiKeyMandatory")
				return
			case keyget != "":
				consumer := models.FindConsumerByApiKey(keyget)
				if consumer == nil {
					errors.NewHttpError(w, "ErrorApiKeyInvalid")
					return
				}
			case keyheader != "":
				consumer := models.FindConsumerByApiKey(keyheader)
				if consumer == nil {
					errors.NewHttpError(w, "ErrorApiKeyInvalid")
					return
				}
		}
		next.ServeHTTP(w, r)
	})
}
