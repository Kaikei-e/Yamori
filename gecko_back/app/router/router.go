package router

import (
	"github.com/labstack/echo/v4"
)

func Router() (*echo.Echo, error) {
	e := echo.New()

	//err := crossLogging.Watcher(e)
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}

	e.Group("/api/v1")
	e.Use()
	{

	}

	return e, nil

}
