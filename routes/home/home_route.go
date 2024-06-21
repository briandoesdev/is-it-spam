package home

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	//routeGroup := e.Group("/")
	e.GET("/", home)
}

func home(c echo.Context) error {
	return c.Render(200, "home.html", nil)
}
