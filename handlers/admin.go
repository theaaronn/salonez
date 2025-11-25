package handlers

import (
	"Salonez/db"
	"Salonez/static/views"
	"Salonez/utils"

	"github.com/labstack/echo/v4"
)

// AdminDashboard muestra todas las reservaciones del sistema
func AdminDashboard(c echo.Context) error {
	reservations, err := db.GetAllReservations()
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return utils.Render(c, 200, views.AdminHome(reservations, c))
}
