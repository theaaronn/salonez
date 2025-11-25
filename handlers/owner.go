package handlers

import (
	"Salonez/db"
	"Salonez/models"
	"Salonez/static/views"
	"Salonez/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

// OwnerDashboard muestra los salones del propietario y sus reservaciones
func OwnerDashboard(c echo.Context) error {
	email, _ := c.Cookie("email")
	if email == nil {
		return c.Redirect(302, "/login")
	}

	userId, err := db.GetUserIdByEmail(email.Value)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	// Obtener salones del propietario
	halls, err := db.GetHallsByOwner(userId)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	// Obtener reservaciones de todos los salones del propietario
	allReservations := make([]db.ReservationData, 0)
	for _, hall := range halls {
		reservations, err := db.GetReservationsByHall(int(hall.Id))
		if err != nil {
			continue
		}
		allReservations = append(allReservations, reservations...)
	}

	return utils.Render(c, 200, views.PropietaryHome(halls, allReservations, c))
}

// CreateHall crea un nuevo salón
func CreateHall(c echo.Context) error {
	email, _ := c.Cookie("email")
	if email == nil {
		return c.Redirect(302, "/login")
	}

	userId, err := db.GetUserIdByEmail(email.Value)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	hall := models.Hall{
		Nombre:         c.FormValue("nombre"),
		Direccion:      c.FormValue("direccion"),
		Capacidad:      c.FormValue("capacidad"),
		NumeroTelefono: c.FormValue("telefono"),
		ImgsPath:       []string{}, // Por ahora sin imágenes
	}

	precio, err := strconv.ParseFloat(c.FormValue("precio"), 64)
	if err != nil {
		precio = 0
	}
	hall.Precio = precio

	err = db.CreateHall(hall, userId)
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}

	return c.Redirect(302, "/dashboard/owner")
}
