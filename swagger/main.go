package main

import (
	"net/http"

	_ "kbtg-boocamp-cache/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const COUNTER_CACHE_KEY = "counter"

// Pet model info
// @Description Pet information
type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HTTPError model info
// @Description HTTPError information
type HTTPError struct {
	BusinessErrorCode int    `json:"businessErrorCode"`
	Description       string `json:"description"`
}

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/pets", allPet)
	e.Logger.Fatal(e.Start(":1323"))
}

// ListPet godoc
// @Summary      List of all pet
// @Description  list pets
// @Tags         pet
// @Accept       json
// @Produce      json
// @Success      200  {array}   Pet
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /pets [get]
func allPet(c echo.Context) error {
	return c.JSON(http.StatusOK, []Pet{
		{ID: 1, Name: "ChaNom"},
		{ID: 2, Name: "Mali"},
	})
}
