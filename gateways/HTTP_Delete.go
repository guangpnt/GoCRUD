package gateways

import (
	"crud/dao"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(dao.DeleteAnimalReq)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	err := sv.Service.DeleteAnimal(ctx, &dao.DeleteAnimalReq{
		Species: r.Species,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, map[string]string{
		"message": "Deleted",
	})
}
