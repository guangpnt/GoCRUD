package gateways

import (
	"crud/dao"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) Edit(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(dao.GetAnimalRes)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	err := sv.Service.EditAnimal(ctx, &dao.GetAnimalRes{
		ID:      r.ID,
		Species: r.Species,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, map[string]string{
		"message": "Edit complete",
	})
}
