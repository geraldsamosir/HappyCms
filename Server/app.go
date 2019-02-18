package Server

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Config"
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Service/Content"
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Service/Static"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	var c Config.ConfigService
	var db Config.MonggodbConnector	
	c.ServiceConf()
	a :=  db.Connect()
	fmt.Println(a)
	fmt.Println("your server run in port " + viper.GetString("port"))
	http.ListenAndServe("0.0.0.0:"+viper.GetString("port"), r)
}
