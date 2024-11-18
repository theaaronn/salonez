package main

import (
	"Salonez/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "/static")

	e.GET("/", handlers.IndexHandler)
	e.GET("/halls", handlers.HallList)

	e.GET("/login", handlers.ShowLoginHandler)
	e.POST("/login", handlers.ValidateLogin)

	e.POST("/signup", handlers.SignUpHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:6969"))
}
