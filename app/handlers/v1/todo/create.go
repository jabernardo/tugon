package todo

import (
	"encoding/json"
	"net/http"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

// @Description   Create a TODO item
// @Tags          todo
// @Accept        json
// @Produce       json
// @Param         data body repositories.Todo true "Todo Object"
// @Success       200 {object} WrappedCreateItem
// @Failure       400 {object} core.FailureResponse
// @Router        /v1/todo/ [put]
func Create(w http.ResponseWriter, r *http.Request) {
	todoRepo := repositories.NewTodoRepository()
	var todo repositories.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid payload").Write(w)
		return
	}

	res, err := todoRepo.Create(&todo)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Insert failed").Write(w)
		return
	}

	id, _ := res.LastInsertId()

	core.NewSuccessResponse(id).Write(w, nil)
}
