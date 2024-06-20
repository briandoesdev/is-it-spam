package main

import (
	"log"

	"github.com/briandoesdev/caller-lookup/config"
	"github.com/briandoesdev/caller-lookup/services/openai"
	"github.com/briandoesdev/caller-lookup/services/twilio"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}
	log.Printf("Loaded env variables.")

	// load config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("error loading config: ", err)
	}
	log.Printf("Loaded config.")

	// initialize services
	twilio.InitService(config.Twilio)
	openai.InitService(config.OpenAI)
	log.Printf("Loaded services.")

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//e.Logger.Fatal(e.Start(config.Server.Host + ":" + config.Server.Port))
}
