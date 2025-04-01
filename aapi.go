package main

import (
	"log"
	"os"

	"github.com/jabernardo/aapi/app"
	"github.com/jabernardo/aapi/core"

	"github.com/joho/godotenv"
)

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
