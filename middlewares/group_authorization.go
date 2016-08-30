package middlewares

import (
	"github.com/xaviergodart/gydro/errors"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

func GroupAuthorization(apiGroups []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		consumer := r.Context().Value("consumer").(*models.Consumer)
		if authorize(apiGroups, consumer.Groups) == false {
			errors.NewHttpError(w, "ErrorNoRequiredGroup")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func authorize(apiGroups, consumerGroups []string) bool {
	if len(apiGroups) == 0 {
		return true
	}

	for _, cgroup := range consumerGroups {
		for _, agroup := range apiGroups {
			if cgroup == agroup {
				return true
			}
		}
	}

	return false
}
