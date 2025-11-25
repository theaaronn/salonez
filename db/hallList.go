package db

import (
	"Salonez/models"

	"github.com/lib/pq"
)

func GetAllHalls() ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			h.id,
			h.imgs_path,
			h.nombre, 
			h.numero_telefono, 
			h.capacidad, 
			h.precio,
			h.direccion
		FROM halls h
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hall models.Hall
		var imgsPaths pq.StringArray

		err := rows.Scan(&hall.Id, &imgsPaths, &hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &hall.Direccion)
		if err != nil {
			return nil, err
		}
		hall.ImgsPath = []string(imgsPaths)
		halls = append(halls, hall)
	}
	return halls, nil
}

func SearchHalls(param string) ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
	defer db.Close()

	query := `
		SELECT
			h.id,
			h.imgs_path,
			h.nombre, 
			h.numero_telefono, 
			h.capacidad, 
			h.precio,
			h.direccion
		FROM halls h
		WHERE h.nombre ILIKE $1
	`
	rows, err := db.Query(query, "%"+param+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hall models.Hall
		var imgsPaths pq.StringArray

		err := rows.Scan(&hall.Id, &imgsPaths, &hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &hall.Direccion)
		if err != nil {
			return nil, err
		}
		hall.ImgsPath = []string(imgsPaths)
		halls = append(halls, hall)
	}
	return halls, nil
}

func GetHallsByOwner(ownerId int) ([]models.Hall, error) {
	halls := make([]models.Hall, 0)
	db := GetDb()
	defer db.Close()

	query := `
		SELECT 
			h.id,
			h.imgs_path,
			h.nombre, 
			h.numero_telefono, 
			h.capacidad, 
			h.precio,
			h.direccion
		FROM halls h
		WHERE h.propietario_id = $1
	`
	rows, err := db.Query(query, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hall models.Hall
		var imgsPaths pq.StringArray

		err := rows.Scan(&hall.Id, &imgsPaths, &hall.Nombre, &hall.NumeroTelefono, &hall.Capacidad, &hall.Precio, &hall.Direccion)
		if err != nil {
			return nil, err
		}
		hall.ImgsPath = []string(imgsPaths)
		halls = append(halls, hall)
	}
	return halls, nil
}

func CreateHall(hall models.Hall, ownerId int) error {
	db := GetDb()
	defer db.Close()

	query := `
		INSERT INTO halls (imgs_path, nombre, direccion, capacidad, numero_telefono, precio, propietario_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := db.Exec(query, pq.Array(hall.ImgsPath), hall.Nombre, hall.Direccion, hall.Capacidad, hall.NumeroTelefono, hall.Precio, ownerId)
	return err
}
