package handlers

import (
	"Salonez/db"
	"Salonez/static/views"
	"Salonez/utils"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	halls, err := db.GetAllHalls()
	if err != nil {
		return utils.Render(c, 505, views.InternalServerError())
	}
	content := views.HallList(halls)
	return utils.Render(c, 200, views.Layout("Salonez", content, c))
}
