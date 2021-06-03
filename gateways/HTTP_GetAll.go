package gateways

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) GetAll(c echo.Context) error {
	res, err := sv.Service.GetAllAnimal(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, res)
}
