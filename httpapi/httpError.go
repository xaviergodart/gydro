package httpapi

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

func NewHttpError(w http.ResponseWriter, err string) http.Error {
	http.Error(w, HttpErrors[err].getJSON(), HttpErrors[err].Code)
	return
}

func (e *HttpError) getJSON() string {
	jsonError, _ := json.Marshal(e)
	return string(jsonError[:])
}
