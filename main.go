package main

import (
	//"net/http"

	asset "github.com/K-Kaotsane/gofarm/constant"
	"github.com/K-Kaotsane/gofarm/router"
	"github.com/K-Kaotsane/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Renderer = asset.LoadTemplate()

	asset.LoadAssets(app)
	//app.Start(":3000")
	/*app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "My first web")
	})*/

	router.LoadAllRouters(app)
	server.SetServer(app)
}
