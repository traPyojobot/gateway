package gateway

import "github.com/labstack/echo/v4"

func SetUp(e *echo.Echo) {
	api := e.Group("/api")

	apiReply := api.Group("/reply")
	apiReply.POST("", GetReply)

	apiMonologue := api.Group("/monologue")
	apiMonologue.GET("", GetMonologue)
}
