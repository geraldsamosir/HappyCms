package Utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message    string      `json:"message" validate:"required"`
	Data       interface{} `json:"data"`
	Status     int         `json:"status" validate:"required,number"`
	Validation interface{} `json:validation`
}

type ValidationType struct {
	ValidationBody     []string `json:validationBody`
	ValidationQueryUrl []string `json:validationQueryUrl`
}

func (r *Response) ResponseJSON(responseBody interface{}, res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response, _ := json.Marshal(responseBody)
	res.Write(response)
}
