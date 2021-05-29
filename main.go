package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traPyojobot/gateway/cmd/gateway"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	gateway.SetUp(e)
	e.Logger.Fatal(e.Start(":8010"))
}
