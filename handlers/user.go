package handlers

import (
	"Salonez/db"
	"Salonez/static/views"
	"Salonez/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserDashboard muestra salones disponibles y reservaciones del usuario
func UserDashboard(c echo.Context) error {
	email, _ := c.Cookie("email")
	if email == nil {
		return c.Redirect(302, "/login")
	}

	userId, err := db.GetUserIdByEmail(email.Value)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	// Obtener todos los salones
	halls, err := db.GetAllHalls()
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	// Obtener reservaciones del usuario
	reservations, err := db.GetReservationsByUser(userId)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return utils.Render(c, 200, views.UserHome(halls, reservations, c))
}

// ReserveHall crea una nueva reservación
func ReserveHall(c echo.Context) error {
	email, _ := c.Cookie("email")
	if email == nil {
		return c.Redirect(302, "/login")
	}

	userId, err := db.GetUserIdByEmail(email.Value)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	hallId, err := strconv.Atoi(c.FormValue("hall_id"))
	if err != nil {
		return utils.Render(c, 400, views.InternalServerError())
	}

	date := c.FormValue("date")
	time := c.FormValue("time")
	percentage := c.FormValue("percentage")

	percentageFloat := 0.0
	if percentage != "" {
		percentageFloat, err = strconv.ParseFloat(percentage, 64)
		if err != nil || percentageFloat < 0 || percentageFloat > 100 {
			return utils.Render(c, 400, views.InternalServerError())
		}
	}

	err = db.CreateReservation(hallId, userId, date, time, percentageFloat)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return c.Redirect(302, "/dashboard/user")
}

// ProcessPayment procesa un pago adicional en una reservación
func ProcessPayment(c echo.Context) error {
	reservationId, err := strconv.Atoi(c.FormValue("reservation_id"))
	if err != nil {
		return utils.Render(c, 400, views.InternalServerError())
	}

	additionalPercentage := c.FormValue("additional_percentage")
	additionalFloat, err := strconv.ParseFloat(additionalPercentage, 64)
	if err != nil || additionalFloat <= 0 || additionalFloat > 100 {
		return utils.Render(c, 400, views.InternalServerError())
	}

	err = db.ProcessPayment(reservationId, additionalFloat)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return c.Redirect(302, "/dashboard/user")
}

// CancelReservation cancela una reservación existente
func CancelReservation(c echo.Context) error {
	reservationId, err := strconv.Atoi(c.FormValue("reservation_id"))
	if err != nil {
		return utils.Render(c, 400, views.InternalServerError())
	}

	err = db.CancelReservation(reservationId)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return c.Redirect(302, "/dashboard/user")
}
