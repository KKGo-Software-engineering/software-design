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

```
4. Update `main.go` file to the following code:

```go
	myAPI := serverWrapper{}
	e.POST("/pet", myAPI.AddPet)
	petstore.RegisterHandlers(e, &myAPI)
```

## Plugins

- VS Code: OpenAPI (Swagger) Editor
- GoLand: OpenAPI Specification by Jetbrains
