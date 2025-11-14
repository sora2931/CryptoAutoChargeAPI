package main

import (
	"coin/internal/utils"
	"coin/internal/server"
)

func main() {
	e := server.InitEcho()
	err := utils.LoadConfig("configs.json")
	if err != nil {
		panic(err)
	}

	e.Start(":8080")
}