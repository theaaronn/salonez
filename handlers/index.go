package handlers

import (
	"Salonez/static/views"
	"Salonez/utils"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	
	var content templ.Component = views.Index()

	return utils.Render(c, 200, views.Layout("Salonez", content))
}
