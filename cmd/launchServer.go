package main

import (
	"Salonez/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.Static("/assets", "static/assets")

	// Rutas públicas
	e.GET("/", handlers.IndexHandler)
	e.GET("/halls", handlers.HallList)

	// Autenticación
	e.GET("/login", handlers.ShowLoginHandler)
	e.POST("/login", handlers.ValidateLogin)
	e.POST("/signup", handlers.SignUpHandler)
	e.GET("/logout", handlers.Logout)

	// Dashboard Admin
	e.GET("/dashboard/admin", handlers.AdminDashboard)

	// Dashboard Owner
	e.GET("/dashboard/owner", handlers.OwnerDashboard)
	e.POST("/dashboard/owner/create-hall", handlers.CreateHall)

	// Dashboard User
	e.GET("/dashboard/user", handlers.UserDashboard)
	e.POST("/dashboard/user/reserve", handlers.ReserveHall)
	e.POST("/dashboard/user/cancel", handlers.CancelReservation)
	e.POST("/dashboard/user/pay", handlers.ProcessPayment)

	e.Logger.Fatal(e.Start("127.0.0.1:6969"))
}
