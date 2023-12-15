package main

import (
	"kbtg-bootcamp-petstore/petstore"

	"github.com/labstack/echo/v4"
)

type serverWrapper struct {
	Handler petstore.ServerInterface
}

// (PUT /pet)
func (w *serverWrapper) UpdatePet(ctx echo.Context) error {
	return ctx.String(200, "Updated")
}

// (POST /pet)
func (w *serverWrapper) AddPet(ctx echo.Context) error {
	return ctx.String(200, "Added")
}
