# Generate Swagger Step

1. Run `go install github.com/swaggo/swag/cmd/swag@latest`
2. Run `swag init` then you'll see directory `docs`
3. Run `go get -u github.com/swaggo/echo-swagger`
4. Import `github.com/swaggo/echo-swagger` and add `_ "[module]/docs"`
5. add snippet `e.GET("/swagger/*", echoSwagger.WrapHandler)`
6. add Swagger comment then run `swag init`
7. `go run main.go`
8. go to `http://localhost:1323/swagger/index.html`