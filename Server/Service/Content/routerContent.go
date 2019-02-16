package Content

import (
	"github.com/gorilla/mux"
)

type RouterContent struct {
}

func (RC *RouterContent) AddSignHandler(r *mux.Router) *mux.Router {
	var serviceContent ServiceContent
	r.HandleFunc("/content", serviceContent.WelcomeContent).Methods("GET")
	return r
}
