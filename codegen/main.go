package main

import (
	"kbtg-bootcamp-petstore/petstore"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	myAPI := serverWrapper{}
	petstore.RegisterHandlers(e, &myAPI)
	e.Logger.Fatal(e.Start(":1323"))
}
