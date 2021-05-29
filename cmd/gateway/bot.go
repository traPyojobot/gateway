package gateway

import (
	"github.com/labstack/echo/v4"
	"github.com/traPyojobot/gateway/pkg/model"
)

func PostReply(c echo.Context) error {
	var m *model.Message
	if err := c.Bind(m); err != nil {
		return echo.ErrBadRequest
	}

}
