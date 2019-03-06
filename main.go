package main

import (
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.NoContent(200)
	})
	e.GET("/crash", func(c echo.Context) error {
		go func() {
			time.Sleep(5 * time.Second)
			panic("crashed")
		}()
		return c.NoContent(500)
	})
	e.Start(":8000")
}
