package http

import (
	context "github.com/K-Kaotsane/gofarm/controller/context/pages"
	"github.com/labstack/echo/v4"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", context.IndexContext)

}
