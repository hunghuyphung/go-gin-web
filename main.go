package main

import (
	"github.com/joho/godotenv"
	"go-gin-web/app/route"
	"go-gin-web/config"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := route.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		return
	}
}
