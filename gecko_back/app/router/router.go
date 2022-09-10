package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Router() (*echo.Echo, error) {
	e := echo.New()

	//err := crossLogging.Watcher(e)
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}

	v1 := e.Group("/api/v1")
	v1.Use()
	{
		v1.GET("/ping", func(c echo.Context) error {
			c.String(http.StatusOK, "pong")
			return nil
		})
	}

	return e, nil

}
