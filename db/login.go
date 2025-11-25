package db

func CheckCredentials(email, password string) (string, error) {
	db := GetDb()
	defer db.Close()

	var (
		returnedEmail string
		userType      string
	)

	// Buscar usuario y obtener su tipo
	query := "SELECT correo, tipo_usuario FROM users WHERE correo = $1"
	row := db.QueryRow(query, email)
	err := row.Scan(&returnedEmail, &userType)
	if err != nil {
		return "", err
	}

	// Para demo, aceptar passwords simples
	if returnedEmail == email && password != "" {
		// Convertir tipo de usuario a número
		switch userType {
		case "usuario":
			return "1", nil
		case "propietario":
			return "2", nil
		case "admin":
			return "3", nil
		}
	}

	return "", nil
}
