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
	_, err := c.Cookie("pass")
	if err != nil {
		return utils.Render(c, 200, views.LoginSignup())
	}
	userType, err := c.Cookie("type")
	if err != nil {
		return utils.Render(c, 200, views.LoginSignup())
	}
	switch userType.Value {
	case "1":
		return utils.Render(c, 200, views.UserMenu())
	case "2":
		return utils.Render(c, 200, views.PropietaryMenu())
	case "3":
		return utils.Render(c, 200, views.AdminMenu())
	default:
		return utils.Render(c, 200, views.LoginSignup())
	}
}

func ValidateLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	userType := c.FormValue("userType")

	existent, err := db.CheckCredentials(email, password, userType)
	if err != nil || !existent {
		c.Response().Header().Set("HX-Retarget", "#login-error")
		return utils.Render(c, 422, views.LoginErr())
	}

	// TODO: Generate session passCookie
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

	halls, err := db.GetAllHalls()
	if err != nil {
		return utils.Render(c, 500, views.InternalServerError())
	}
	return utils.Render(c, 200, views.HallList(halls))
}
