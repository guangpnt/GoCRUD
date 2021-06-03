package main

import (
	"crud/database/datasources"
	"crud/database/repository"
	"crud/gateways"
	"crud/services"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := datasources.NewMongoDB()
	repo := repository.NewRepository(db)
	sv := services.NewServiceCRUD(repo)

	routG := e.Group("/")

	gateways.NewHTTP(routG, sv)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
