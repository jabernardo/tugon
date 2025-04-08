package todo

import (
	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
)

type WrappedGetAllResponse struct {
	core.SuccessResponse
	Data []repositories.Todo
}

type WrappedGetItem struct {
	core.SuccessResponse
	Data repositories.Todo
}

type WrappedCreateItem struct {
	core.SuccessResponse
	Data int `json:"data"`
}
