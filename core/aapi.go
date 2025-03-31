package core

import (
	"log"
	"net/http"
)

type API struct {
	Version string
	Mux     http.ServeMux
}

func New(version string) *API {
	return &API{
		Version: version,
		Mux:     *http.NewServeMux(),
	}
}

func (api *API) Use(router *Router) {
	for key, val := range router.GetRoutes() {
		api.Mux.Handle(key, val)
	}
}

func (api *API) ListenAndServe(addr string) {
	err := http.ListenAndServe(addr, &api.Mux)

	if err != nil {
		log.Fatalln(err)
	}
}
