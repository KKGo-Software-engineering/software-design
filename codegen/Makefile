install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

create-petstore:
	mkdir petstore

generate:
	oapi-codegen -package petstore openapi.yaml > petstore/petstore.gen.go
