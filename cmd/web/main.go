package main

import (
	"log"
	"syscall"

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
		syscall.Exit(1)
	}

	// load config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("error loading config: ", err)
		syscall.Exit(1)
	}
	log.Printf("Loaded config.")

	// initialize services
	twilio.InitService(config.Twilio.AccountSid, config.Twilio.AuthToken)
	openai.InitService(config.OpenAI.ApiKey, config.OpenAI.Model)
	log.Printf(openai.GenerateCompletions("My name is Brian. Repeat it back to me."))

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(config.Server.Host + ":" + config.Server.Port))
}
