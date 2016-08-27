package middlewares

import (
	"github.com/xaviergodart/gydro/errors"
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/ratelimiter"
	"net/http"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		consumer := r.Context().Value("consumer").(*models.Consumer)
		if ratelimiter.IsExceeded(consumer.GetId(), consumer.RateLimit) {
			errors.NewHttpError(w, "ErrorQuotaLimitReached")
			return
		}
		next.ServeHTTP(w, r)
	})
}
