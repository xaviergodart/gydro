package errors

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var HttpErrors map[string]*HttpError = map[string]*HttpError{
	"ErrorApiKeyMandatory":   &HttpError{Code: http.StatusUnauthorized, Message: "apikey is mandatory"},
	"ErrorApiKeyInvalid":     &HttpError{Code: http.StatusUnauthorized, Message: "given apikey is invalid"},
	"ErrorQuotaLimitReached": &HttpError{Code: http.StatusForbidden, Message: "quota limit reached"},
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
