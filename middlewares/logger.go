package middlewares

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func Logger(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
