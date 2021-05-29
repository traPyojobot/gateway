package gateway

import "github.com/labstack/echo/v4"

func SetUp(e *echo.Group) {
	api := e.Group("/api")

	apiReply := api.Group("/reply")
	apiReply.GET("", PostReply)
}
