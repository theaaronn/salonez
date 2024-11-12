package main

import (
	"Salonez/auth"
	"Salonez/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "/static")

	e.GET("/", handlers.HallListHandler)
	e.GET("/login", handlers.ShowLoginHandler)
	e.POST("/login", auth.AuthHandler)

	e.POST("/signup", handlers.SignUpHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:6969"))
}
