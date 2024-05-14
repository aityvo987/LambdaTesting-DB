package main

import (
	"context"
	handler "lambda-test/internal/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/employee", handler.GetEmployee)
	r.GET("/employees", handler.GetEmployees)
	r.POST("/employee", handler.CreateEmployee)

	r.DELETE("/employee", handler.DeleteEmployee)
	r.PUT("/employee", handler.UpdateEmployee)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
