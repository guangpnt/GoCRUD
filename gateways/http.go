package gateways

import (
	"crud/services"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	Service services.IServicesCRUD
}

func NewHTTP(g *echo.Group, sv services.IServicesCRUD) {
	h := &HTTP{
		Service: sv,
	}

	g.GET("", h.GetAll)
	g.POST("", h.Add)
	g.PUT("", h.Edit)
	g.DELETE("", h.Delete)
}
