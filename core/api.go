package core

import (
	"log"
	"net/http"

	_ "github.com/jabernardo/tugon/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	Addr   string
	Router *Router
}

func New(router *Router, addr string) *API {
	return &API{
		Addr:   addr,
		Router: router,
	}
}

func (api *API) ListenAndServe() {
	api.Router.Mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	Logger().Info("Running at", "addr", api.Addr)
	err := http.ListenAndServe(api.Addr, &api.Router.Mux)

	if err != nil {
		log.Fatalln(err)
	}
}
