package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.NoContent(200)
	})
	e.Start(":8000")
}
