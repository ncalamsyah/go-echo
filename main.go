package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/go-echo/cmd/api/handlers"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostHandler)
	e.GET("/post/:id", handlers.PostDetailHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
