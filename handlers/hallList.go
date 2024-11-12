package handlers

import (
	"Salonez/db"
	"Salonez/models"
	"Salonez/static/views"
	"Salonez/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func HallListHandler(c echo.Context) error {
	halls := make([]models.Hall, 0)
	db := db.GetDb()
	rows, err := db.Query(`
	SELECT 
			s.Nombre, 
			s.NumeroTelefono, 
			s.Capacidad, 
			s.Precio, 
			u.Calle, 
			u.Numero, 
			u.Colonia, 
			u.CP, 
			u.Ciudad, 
			u.Estado 
			FROM 
			Salon s
			JOIN 
			Ubicacion u ON s.idUbicacion = u.idUbicacion
	`)
	if err != nil {
		log.Fatal(err)
		return utils.Render(c, 500, views.InternalServerError())
	}
	defer rows.Close()

	for rows.Next() {
		var hall = models.Hall{}
		var calle, numero, colonia, ciudad, estado string
		var cp int
		err := rows.Scan(&hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &calle, &numero, &colonia, &cp, &ciudad, &estado)
		if err != nil {
			log.Fatal(err)
			return utils.Render(c, 500, views.InternalServerError())
		}
		hall.Direccion = fmt.Sprintf("%s %s, %s, %d, %s, %s", calle, numero, colonia, cp, ciudad, estado)
		halls = append(halls, hall)
	}
	return utils.Render(c, 200, views.Layout("MainPage",views.HallList(halls)))
}
