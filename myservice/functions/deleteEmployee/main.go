package main

import (
	handler "lambda-test/internal/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.DeleteEmployee)
}
