package Utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message    string   `json:"message" validate:"required"`
	Status     int      `json:"status" validate:"required,number"`
	Validation []string `json:validation`
}

func (r *Response) ResponseJSON(responseBody interface{}, res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response, _ := json.Marshal(responseBody)
	res.Write(response)
}
