package main

import (
	"fmt"

	"github.com/jabernardo/tugon/app/repositories"
	"github.com/jabernardo/tugon/core"
	"github.com/joho/godotenv"
)

func setupRepositories() {
	db := core.GetDBInstance()

	repositories.SetupTodoRepository(db)

	todo := repositories.NewTodoRepository()

	todo.Create(&repositories.Todo{Title: "Hello", Description: "Just Greet"})
}

func bootstrap() {
	err := godotenv.Load()

	if err != nil {
		core.Logger().Warn("Could not load `.env`", "err", err)

	}
	setupRepositories()
}
