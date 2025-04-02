package main

import (
	"log"
	"os"

	"github.com/jabernardo/tugon/app"
	"github.com/jabernardo/tugon/core"

	"github.com/joho/godotenv"
)

// @title         Tugon
// @version       1.0
// @description   This is a simple REST API for the Boiler Plate API

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("[aapi] could not load .env", err)
	}

	api := core.New("1.0")

	api.Use(app.GetRouter())

	addr := ":5000"

	if addrEnv, ok := os.LookupEnv("ADDR"); ok {
		addr = addrEnv
	}

	api.ListenAndServe(addr)
}
