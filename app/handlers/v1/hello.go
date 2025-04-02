package v1

import (
	"fmt"
	"net/http"
)

// Hello World!
//
// @Description   A simple greeting earthlings!
// @Produce       plain
// @Success       200   {string} string "Hello World!"
//
// @Router        /v1/hello [get]
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	w.WriteHeader(200)
}
