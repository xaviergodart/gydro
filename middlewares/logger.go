package middlewares

import (
	"github.com/gorilla/handlers"
	"os"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
