package Static

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "client/public"
)

type RouterStatic struct {
}

func (RS *RouterStatic) AddSignHandler(r *mux.Router) *mux.Router {
	r.PathPrefix("/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir(STATIC_DIR))))
	return r
}
