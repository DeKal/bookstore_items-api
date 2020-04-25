package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/DeKal/bookstore_utils-go/errors"
)

// WriteJSONResponse write response in json  with status
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// WriteReponseError write response with given error
func WriteReponseError(w http.ResponseWriter, err *errors.RestError) {
	WriteJSONResponse(w, err.Status, err)
}
