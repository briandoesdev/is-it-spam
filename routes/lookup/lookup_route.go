package lookup

import (
	"fmt"

	"github.com/briandoesdev/caller-lookup/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	mw "github.com/briandoesdev/caller-lookup/middleware"
	ai "github.com/briandoesdev/caller-lookup/services/openai"
	"github.com/briandoesdev/caller-lookup/services/twilio"
)

type Data struct {
	PhoneNumber string `json:"phone_number"`
	Summary     string `json:"summary"`
}

func Route(e *echo.Echo) {
	// create a new group for the lookup routes
	routeGroup := e.Group("/lookup")

	// register group level middleware
	routeGroup.Use(middleware.CORS())
	routeGroup.Use(mw.DetectResponseFormat)

	// register routes
	routeGroup.POST("", getNumberSummary)
}

func getNumberSummary(c echo.Context) error {
	var sum string
	number := c.FormValue("number")

	if number == "" {
		return c.JSON(400, &routes.ErrorPayload{Message: "phone number is required"})
	}

	// twilio service lookup
	t, err := twilio.Lookup(number)
	if err != nil {
		return c.JSON(500, &routes.ErrorPayload{Message: err.Error()})
	}

	// openai service summarization
	sum, err = ai.GenerateCompletions(t)
	if err != nil {
		return c.JSON(500, &routes.ErrorPayload{Message: err.Error()})
	}

	f := c.Request().Header.Get("X-Response-Format")
	switch f {
	case "api":
		return c.JSON(200, &Data{PhoneNumber: number, Summary: sum})
	case "htmx":
		return c.HTML(200, fmt.Sprintf("<p id=\"summary\" name=\"summary\">%s</p>", sum))
	default:
		return c.Render(200, "main.html", map[string]interface{}{"sum": sum})
	}
}
