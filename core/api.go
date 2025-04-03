package core

import (
	"log"
	"net/http"

	_ "github.com/jabernardo/tugon/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	version string
	mux     http.ServeMux
}

func New(version string) *API {
	return &API{
		version: version,
		mux:     *http.NewServeMux(),
	}
}

func (api *API) Use(router *Router) {
	for key, val := range router.GetRoutes() {
		api.mux.Handle(key, val)
	}
}

func (api *API) ListenAndServe(addr string) {
	api.mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	GetLoggerInstance().Info("Running at", "addr", addr)
	err := http.ListenAndServe(addr, &api.mux)

	if err != nil {
		log.Fatalln(err)
	}
}
