package handlers

import (
	"Salonez/db"
	"Salonez/static/views"
	"Salonez/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ShowLoginHandler(c echo.Context) error {
	// Verificar si ya tiene sesión activa
	emailCookie, err := c.Cookie("email")
	if err == nil && emailCookie.Value != "" {
		userTypeCookie, _ := c.Cookie("type")
		// Si ya está autenticado, redirigir al dashboard correspondiente
		switch userTypeCookie.Value {
		case "3":
			return c.Redirect(302, "/dashboard/admin")
		case "2":
			return c.Redirect(302, "/dashboard/owner")
		case "1":
			return c.Redirect(302, "/dashboard/user")
		}
	}
	// Mostrar página de login sin layout (solo el formulario)
	return utils.Render(c, 200, views.LoginSignup())
}

func ValidateLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	userType, err := db.CheckCredentials(email, password)
	if err != nil || userType == "" {
		c.Response().Header().Set("HX-Retarget", "#login-error")
		return utils.Render(c, 422, views.LoginErr())
	}

	// Guardar cookies de sesión
	passCookie := new(http.Cookie)
	passCookie.Name = "pass"
	passCookie.Value = password
	passCookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(passCookie)

	userTypeCookie := new(http.Cookie)
	userTypeCookie.Name = "type"
	userTypeCookie.Value = userType
	userTypeCookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(userTypeCookie)

	emailCookie := new(http.Cookie)
	emailCookie.Name = "email"
	emailCookie.Value = email
	emailCookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(emailCookie)

	// Redirigir según tipo de usuario
	switch userType {
	case "3": // Admin
		return c.Redirect(302, "/dashboard/admin")
	case "2": // Owner
		return c.Redirect(302, "/dashboard/owner")
	case "1": // User
		return c.Redirect(302, "/dashboard/user")
	default:
		return c.Redirect(302, "/")
	}
}

func Logout(c echo.Context) error {
	// Borrar todas las cookies
	passCookie := new(http.Cookie)
	passCookie.Name = "pass"
	passCookie.Value = ""
	passCookie.MaxAge = -1
	c.SetCookie(passCookie)

	userTypeCookie := new(http.Cookie)
	userTypeCookie.Name = "type"
	userTypeCookie.Value = ""
	userTypeCookie.MaxAge = -1
	c.SetCookie(userTypeCookie)

	emailCookie := new(http.Cookie)
	emailCookie.Name = "email"
	emailCookie.Value = ""
	emailCookie.MaxAge = -1
	c.SetCookie(emailCookie)

	return c.Redirect(302, "/login")
}
