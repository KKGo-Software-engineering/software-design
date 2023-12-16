package pet

import (
	"kbtg-bootcamp-petstore/petstore"

	"github.com/labstack/echo/v4"
)

type ServerWrapper struct {
	Handler petstore.ServerInterface
}

func NewServerWrapper() *ServerWrapper {
	return &ServerWrapper{}
}

// (PUT /pet)
func (w *ServerWrapper) UpdatePet(ctx echo.Context) error {
	return ctx.String(200, "Updated")
}

// (POST /pet)
func (w *ServerWrapper) AddPet(ctx echo.Context) error {
	return ctx.String(200, "Added")
}
