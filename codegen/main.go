package main

import (
	"github.com/labstack/echo/v4"
	"kbtg-bootcamp-petstore/petstore"
)

func main() {
	e := echo.New()
	myAPI := serverWrapper{}
	e.POST("/pet", myAPI.AddPet)
	petstore.RegisterHandlers(e, &myAPI)
	e.Logger.Fatal(e.Start(":1323"))
}
