package Content

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// will deleted after service has handler
type ServiceContent struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type RouterContent struct {
}

func (SC *ServiceContent) WelcomeContent(res http.ResponseWriter, req *http.Request) {
	SC.Message = "Welcome to content"
	SC.Status = 200
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(SC)
}

func (RC *RouterContent) AddSignHandler(r *mux.Router) *mux.Router {
	var serviceContent ServiceContent
	r.HandleFunc("/content", serviceContent.WelcomeContent).Methods("GET")
	return r
}
