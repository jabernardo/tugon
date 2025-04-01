package core

import (
	"log"
	"net/http"
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
	log.Println("[api] running at", addr)
	err := http.ListenAndServe(addr, &api.mux)

	if err != nil {
		log.Fatalln(err)
	}
}
