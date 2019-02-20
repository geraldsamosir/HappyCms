package Utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
}

func (r *Response) ResponseJSON(responseBody interface{}, res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response, _ := json.Marshal(responseBody)
	res.Write(response)
}
