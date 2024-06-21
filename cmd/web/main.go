package main

import (
	"io"
	"log"
	"text/template"

	r "github.com/briandoesdev/caller-lookup/routes"
	home "github.com/briandoesdev/caller-lookup/routes/home"
	lookup "github.com/briandoesdev/caller-lookup/routes/lookup"

	"github.com/briandoesdev/caller-lookup/config"
	"github.com/briandoesdev/caller-lookup/services/openai"
	"github.com/briandoesdev/caller-lookup/services/twilio"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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

	// create server
	e := echo.New()
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	// register app level middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// register routes
	builder := r.NewRouteBuilder(e)
	builder.Register(home.Route)
	builder.Register(lookup.Route)

	e.Logger.Fatal(e.Start(config.Server.Host + ":" + config.Server.Port))
}
