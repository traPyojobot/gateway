package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traPyojobot/gateway/cmd/gateway"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	gateway.SetUp(e)
	err := proxy.LoadModel()
	if err != nil {
		log.Fatal(err)
	}
	e.Logger.Fatal(e.Start(":8010"))
}
