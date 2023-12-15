package main

import (
	"github.com/labstack/echo/v4"
	"kbtg-bootcamp-petstore/petstore"
)

func main() {
	e := echo.New()
	api := NewServerWrapper()
	e.POST("/pet", api.AddPet)
	e.PUT("/pet", api.UpdatePet)
	petstore.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":1323"))
}
