package handlers

import (
	"Salonez/static/views"
	"Salonez/utils"

	"github.com/labstack/echo/v4"
)

func ShowLoginHandler(c echo.Context) error {
	return utils.Render(c, 200, views.Login())
}
