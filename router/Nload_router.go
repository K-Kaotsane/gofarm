package router

import (
	path "github.com/K-Kaotsane/gofarm/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouters(app *echo.Echo) {
	path.IndexRouter(app)

}
