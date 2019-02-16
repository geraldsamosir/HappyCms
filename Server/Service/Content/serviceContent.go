package Content

import (
	"encoding/json"
	"net/http"
)

type ServiceContent struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (SC *ServiceContent) WelcomeContent(res http.ResponseWriter, req *http.Request) {
	SC.Message = "Welcome to content"
	SC.Status = http.StatusOK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(SC)
}
