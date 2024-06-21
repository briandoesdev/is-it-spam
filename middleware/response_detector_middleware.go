package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func DetectResponseFormat(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		host := c.Request().Host
		isHtmx := c.Request().Header.Get("HX-Request") == "true"
		if strings.Contains(host, "api") {
			c.Request().Header.Set("X-Response-Format", "api")
		} else if isHtmx {
			c.Request().Header.Set("X-Response-Format", "htmx")
		} else {
			c.Request().Header.Set("X-Response-Format", "html")
		}

		return next(c)
	}
}
