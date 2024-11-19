package db

import "database/sql"

func CheckCredentials(email, password, userType string) (bool, error) {
	db := GetDb()
	var (
		row           *sql.Row
		returnedEmail string
	)
	switch userType {
	case "1":
		row = db.QueryRow("select Correo from Usuario where Contrasena = ?", password)
	case "2":
		row = db.QueryRow("select Correo from Propietario where Contrasena = ?", password)
	case "3":
		row = db.QueryRow("select Correo from Administrador where Contrasena = ?", password)
	}
	err := row.Scan(&returnedEmail)
	if err != nil {
		return false, err
	}
	if returnedEmail != email {
		return false, nil
	}
	return true, nil
}
