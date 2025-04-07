package v1

import (
	"net/http"

	"github.com/jabernardo/tugon/core"
)

// Hello World!
//
// @Description   A simple greeting earthlings!
// @Produce       plain
// @Success       200   {object} core.SuccessResponse "Hello World!"
//
// @Router        /v1/hello [get]
func Hello(w http.ResponseWriter, r *http.Request) {
	core.NewSuccessResponse("Hello World!").Write(w, nil)
}
