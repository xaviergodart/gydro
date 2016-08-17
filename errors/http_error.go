package errors

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var HttpErrors map[string]*HttpError = map[string]*HttpError {
	"ErrorApiKeyMandatory": &HttpError{Code: 403, Message: "apikey is mandatory"},
}

func NewHttpError(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(HttpErrors[err].Code)
	w.Write(HttpErrors[err].getJSON())
}

func (e *HttpError) getJSON() []byte {
	jsonError, _ := json.Marshal(e)
	return jsonError
}
