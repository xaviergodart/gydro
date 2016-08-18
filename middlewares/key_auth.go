package middlewares

import (
	"github.com/xaviergodart/gydro/errors"
	"github.com/xaviergodart/gydro/models"
	"strconv"
	"net/http"
)

var (
	KeyParam string = "apikey"
)

func KeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyget := r.URL.Query().Get(KeyParam)
		keyheader := r.Header.Get(KeyParam)
		var consumer *models.Consumer
		switch {
			case keyget == "" && keyheader == "":
				errors.NewHttpError(w, "ErrorApiKeyMandatory")
				return
			case keyget != "":
				consumer = models.FindConsumerByApiKey(keyget)
				if consumer == nil {
					errors.NewHttpError(w, "ErrorApiKeyInvalid")
					return
				}
			case keyheader != "":
				consumer = models.FindConsumerByApiKey(keyheader)
				if consumer == nil {
					errors.NewHttpError(w, "ErrorApiKeyInvalid")
					return
				}
		}
		r.Header.Set("X-Consumer-ID", strconv.Itoa(consumer.GetId()))
		r.Header.Set("X-Consumer-Custom-ID", consumer.CustomId)
		r.Header.Set("X-Consumer-Username", consumer.Username)
		next.ServeHTTP(w, r)
	})
}
