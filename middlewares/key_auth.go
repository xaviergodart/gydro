package middlewares

import (
	"github.com/xaviergodart/gydro/errors"
	"github.com/xaviergodart/gydro/models"
	"strconv"
	"net/http"
)

func KeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyget := r.URL.Query().Get("apikey")
		keyheader := r.Header.Get("apikey")
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
				removeParameter("apikey", r)
			case keyheader != "":
				consumer = models.FindConsumerByApiKey(keyheader)
				if consumer == nil {
					errors.NewHttpError(w, "ErrorApiKeyInvalid")
					return
				}
				r.Header.Del("apikey")
		}
		r.Header.Set("X-GYDRO-ConsumerId", strconv.Itoa(consumer.GetId()))
		r.Header.Set("X-GYDRO-ConsumerCustomId", consumer.CustomId)
		next.ServeHTTP(w, r)
	})
}

// Remove parameter from query
func removeParameter(param string, r *http.Request) {
	rParams := r.URL.Query()
	rParams.Del(param)
	r.RequestURI = r.URL.Path + "?" + rParams.Encode()
}
