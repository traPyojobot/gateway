// ひとりごと関連のGatewayの動作
package gateway

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traPyojobot/gateway/pkg/generators/monologue"
	"github.com/traPyojobot/gateway/pkg/model"
)

func GetMonologue(c echo.Context) error {
	m := &model.Message{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := monologue.Create(m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
