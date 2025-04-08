package todo

import (
	"net/http"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

// @Description   Get all TODO items
// @Tags          todo
// @Produce       json
// @Success       200   {object} WrappedGetAllResponse
// @Router        /v1/todo/all [get]
func GetAll(w http.ResponseWriter, r *http.Request) {
	todoRepo := repositories.NewTodoRepository()
	todoItems := todoRepo.GetAll()

	core.NewSuccessResponse(todoItems).Write(w, nil)
}
