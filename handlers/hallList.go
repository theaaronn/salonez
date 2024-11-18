package handlers

import (
	"Salonez/db"
	"Salonez/static/views"
	"Salonez/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func HallList(c echo.Context) error {
	search := c.QueryParam("search")
	if search == "" {
		halls, err := db.GetAllHalls()
		if err != nil {
			return utils.Render(c, 505, views.InternalServerError())
		}
		return utils.Render(c, 200, views.HallList(halls))
	} else {
		halls, err := db.SearchHalls(search)
		if err != nil {
			fmt.Println(err.Error())
			return utils.Render(c, 505, views.InternalServerError())
		}
		return utils.Render(c, 200, views.HallList(halls))
	}
}
