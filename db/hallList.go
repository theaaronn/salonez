package db

import (
	"Salonez/models"
	"fmt"
)

func GetHallList() ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
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
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hall = models.Hall{}
		var calle, numero, colonia, ciudad, estado string
		var cp int
		err := rows.Scan(&hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &calle, &numero, &colonia, &cp, &ciudad, &estado)
		if err != nil {
			return nil, err
		}
		hall.Direccion = fmt.Sprintf("%s %s, %s, %d, %s, %s", calle, numero, colonia, cp, ciudad, estado)
		halls = append(halls, hall)
	}
	return halls, nil
}
