package todo

import (
	"net/http"
	"strconv"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

// @Description   Get specific TODO item
// @Tags          todo
// @Accept        json
// @Produce       json
// @Param         id path int true "TODO ID"
// @Success       200 {object} WrappedGetItem
// @Failure       400 {object} core.FailureResponse
// @Failure       404 {object} core.FailureResponse
// @Router        /v1/todo/{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")

	id, err := strconv.Atoi(idPath)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid ID").Write(w)
		return
	}

	todoRepo := repositories.NewTodoRepository()
	results := todoRepo.Get(id)

	if results == nil {
		core.NewFailureResponse(http.StatusNotFound, "Not Found").Write(w)
		return
	}

	core.NewSuccessResponse(*results).Write(w, nil)
}
