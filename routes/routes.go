package routes

import (
	"github.com/labstack/echo/v4"
)

type RouteBuilder struct {
	e *echo.Echo
}

type ResponseWrapper struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type ErrorPayload struct {
	Message string `json:"message"`
}

func NewRouteBuilder(e *echo.Echo) *RouteBuilder {
	return &RouteBuilder{e}
}

func (rb *RouteBuilder) Register(route func(e *echo.Echo)) {
	route(rb.e)
}

func IsHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
