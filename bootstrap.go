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

	for _, t := range todo.GetAll() {
		fmt.Println(t.Id, t.Title, t.Description)
	}
}

func bootstrap() {
	err := godotenv.Load()

	if err != nil {
		core.GetLoggerInstance().Warn("Could not load `.env`", "err", err)

	}
	setupRepositories()
}
