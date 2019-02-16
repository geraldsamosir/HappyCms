package Server

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Content"
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Static"
	"github.com/gorilla/mux"
)

type MainRouter struct {
}

func (Mr *MainRouter) Router() {
	r := mux.NewRouter().StrictSlash(true)
	api := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	//list api
	var ContentRouter Content.RouterContent
	r.PathPrefix("/api").Handler(negroni.New(
		negroni.Wrap(api),
	))
	ContentRouter.AddSignHandler(api)
	// list of router
	var StaticRouter Static.RouterStatic
	StaticRouter.AddSignHandler(r)

	fmt.Println("your server run in port")
	http.ListenAndServe("0.0.0.0:5000", r)
}
