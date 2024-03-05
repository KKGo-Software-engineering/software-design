package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", body),
		StatusCode: 200,
	}, nil
}
