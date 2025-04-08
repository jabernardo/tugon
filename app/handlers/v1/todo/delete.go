package todo

import (
	"net/http"
	"strconv"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

// @Description   Delete specific TODO item
// @Tags          todo
// @Accept        json
// @Produce       json
// @Param         id path int true "TODO ID"
// @Success       200 {object} core.SuccessResponse
// @Failure       400 {object} core.FailureResponse
// @Failure       404 {object} core.FailureResponse
// @Failure       417 {object} core.FailureResponse
// @Router        /v1/todo/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")

	id, err := strconv.Atoi(idPath)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid ID").Write(w)
		return
	}

	todoRepo := repositories.NewTodoRepository()
	res, err := todoRepo.Delete(id)

	if err != nil {
		core.NewFailureResponse(http.StatusBadRequest, "Invalid ID").Write(w)
		return
	}

	count, err := res.RowsAffected()

	if err != nil {
		core.NewFailureResponse(http.StatusExpectationFailed, "Could not get count of affected rows").Write(w)
		return
	}

	core.NewSuccessResponse(count).Write(w, nil)
}
