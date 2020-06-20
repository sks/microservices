package httputil

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse writes a json response
func WriteJSONResponse(w http.ResponseWriter, val interface{}) {
	err := json.NewEncoder(w).Encode(val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
