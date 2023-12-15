# Code generation

## Instructions

1. Run `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
2. Make directory `mkdir petstore` run `oapi-codegen -package petstore openapi.yaml > petstore/petstore.gen.go`
3. Create `petstore.go` file and add the following code:

```go
package main

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

```

4. Update `main.go` file to the following code:

```go
api := NewServerWrapper()
e.POST("/pet", api.AddPet)
e.PUT("/pet", api.UpdatePet)
petstore.RegisterHandlers(e, api)
```

## Plugins

- VS Code: OpenAPI (Swagger) Editor
- GoLand: OpenAPI Specification by Jetbrains
