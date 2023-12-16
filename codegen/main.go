package main

import (
	"kbtg-bootcamp-petstore/pet"
	"kbtg-bootcamp-petstore/petstore"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api := pet.NewServerWrapper()
	petstore.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":1323"))
}
