package db

import (
	"Salonez/models"
	"fmt"
)

func GetAllHalls() ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
	rows, err := db.Query(`
	SELECT 
			s.idSalon,
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
		var (
			hall                                   models.Hall
			calle, numero, colonia, ciudad, estado string
			cp                                     int
		)
		err := rows.Scan(&hall.Id, &hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &calle, &numero, &colonia, &cp, &ciudad, &estado)
		if err != nil {
			return nil, err
		}
		hall.Direccion = fmt.Sprintf("%s %s, %s, %d, %s, %s", calle, numero, colonia, cp, ciudad, estado)
		halls = append(halls, hall)
	}
	return halls, nil
}

func SearchHalls(param string) ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
	query := `
	SELECT
		s.idSalon,
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
		WHERE s.Nombre LIKE ?
	`
	rows, err := db.Query(query, "%"+param+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			hall                                   models.Hall
			calle, numero, colonia, ciudad, estado string
			cp                                     int
		)
		err := rows.Scan(&hall.Id, &hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &calle, &numero, &colonia, &cp, &ciudad, &estado)
		if err != nil {
			return nil, err
		}
		hall.Direccion = fmt.Sprintf("%s %s, %s, %d, %s, %s", calle, numero, colonia, cp, ciudad, estado)
		halls = append(halls, hall)
	}
	return halls, nil
}
