# Code generation

## Instructions

1. https://editor.swagger.io/
2. Run `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
   - Read more: https://github.com/deepmap/oapi-codegen/tree/master/examples/petstore-expanded
3. Make directory `mkdir petstore` run `oapi-codegen -package petstore openapi.yaml > petstore/petstore.gen.go`
4. Create `petstore.go` file and add the [following code](#petstorego)
5. Update `main.go` file to the [following code](#maingo)
6. Try to add edit `openapi.yaml` file and then generate again

### petstore.go

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

### main.go

```go
api := NewServerWrapper()
petstore.RegisterHandlers(e, api)
```

## Plugins

- VS Code: OpenAPI (Swagger) Editor
- GoLand: OpenAPI Specification by Jetbrains
