package constant

import (
	"github.com/labstack/echo/v4"
)

func LoadAssets(app *echo.Echo) {
	//Load asset
	app.Static("static", "repository/assets")

}
