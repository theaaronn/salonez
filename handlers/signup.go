package handlers

import (
	"Salonez/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(c echo.Context) error {
	// Potencialmente mover esta lógica a su auth correspondiente
	email := c.FormValue("email")
	name := c.FormValue("name")
	password := c.FormValue("password")
	userType := c.FormValue("userType")

	db := db.GetDb()
	verifyEmailQuery := fmt.Sprintf("select email from usuario where email = %v", email)
	result := db.QueryRow(verifyEmailQuery)
	
	if result.Scan() == sql.ErrNoRows {
		// ! Mal
		return c.HTML(401, "<p>Email already taken</p>")
	}
	if userType == "propietary" {
		insertQuery := fmt.Sprintf("insert into Propietario set (Nombre, Correo, Contrasena) values (%v, %v, %v", name, email, password)
		_, err := db.Exec(insertQuery)
		if err != nil {
			// ! Mal
			return c.HTML(505, "<p>Internal server error, try again</p>")
		}
	} else if userType == "user" {
		insertQuery := fmt.Sprintf("insert into Usuario set (Nombre, Correo, Contrasena) values (%v, %v, %v", name, email, password)
		_, err := db.Exec(insertQuery)
		if err != nil {
			// ! Mal
			return c.HTML(505, "<p>Internal server error, try again</p>")
		}
	}
	c.SetCookie(&http.Cookie{
		Name: "user",
		Value: "<something>",
	})
	return c.Redirect(200, "/index")
}
