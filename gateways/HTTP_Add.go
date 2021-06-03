package gateways

import (
	"crud/dao"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) Add(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(dao.CreateAnimalReq)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	_, err := sv.Service.AddAnimal(ctx, &dao.CreateAnimalReq{
		Species: r.Species,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(201, "Added")
}
