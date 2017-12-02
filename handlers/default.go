package handlers

import (
	"encoding/json"
	"net/http"
)

type DefaultHandler struct {
}

// asJSON formats a response as JSON
func (d *DefaultHandler) asJSON(writer *http.ResponseWriter, data interface{}) {
	json.NewEncoder(*writer).Encode(data)
}
