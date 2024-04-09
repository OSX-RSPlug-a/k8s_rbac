package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Device struct {

	UUID string `json:"UUID"`

	Mac string `json:"mac"`

	Firmware string `json:"firmware"`
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/about", about)

	e.Logger.Fatal(e.Start(":8080"))
}

func about(c echo.Context) error {
	d := Device{UUID: "b0e42fe7-31a5-4894-a441-007e5256afea", Mac: "5F-33-CC-1F-43-82", Firmware: "2.1.6"}

	return c.JSON(http.StatusOK, d)
}
