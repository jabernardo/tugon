package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

// @Description   Update a TODO item
// @Tags          todo
// @Accept        json
// @Produce       json
// @Param         id path int true "TODO ID"
// @Param         data body repositories.Todo true "Todo Object"
// @Success       200 {object} core.SuccessResponse
// @Failure       400 {object} core.FailureResponse
// @Router        /v1/todo/{id} [patch]
func Update(w http.ResponseWriter, r *http.Request) {
	todoRepo := repositories.NewTodoRepository()
	idPath := r.PathValue("id")
	var todo repositories.Todo

	id, err := strconv.Atoi(idPath)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid ID").Write(w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid payload").Write(w)
		return
	}

	res, err := todoRepo.Update(id, todo.Title, todo.Description)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Update failed").Write(w)
		return
	}

	count, err := res.RowsAffected()

	if err != nil {
		core.NewFailureResponse(http.StatusExpectationFailed, "Could not get count of affected rows").Write(w)
		return
	}

	core.NewSuccessResponse(count).Write(w, nil)
}
