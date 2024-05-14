package handler

import (
	"context"
	"encoding/json"
	"net/http"

	entity "lambda-test/internal/entity"
	pkg "lambda-test/internal/pkg"
	usecase "lambda-test/internal/usecase"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	GetEmployees(c *gin.Context)
}

func GetEmployees(ctx context.Context,
	req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		resp = events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
		tranReq = entity.GetEmployeesRequest{}
		err     error
	)

	err = json.Unmarshal([]byte(req.Body), &tranReq)
	if err != nil {
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}

	respGen, err := usecase.GetEmployees(tranReq)
	if err != nil {
		log.Printf("Get Employee - ERROR In Get in Usecase: %s\n", err)
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	resp.StatusCode = http.StatusOK
	resp.Body = pkg.Response(entity.AppResponse{
		Error:     err,
		RequestID: tranReq.RequestID,
		Data:      respGen,
	})
	return resp, err

}

func GetEmployee(ctx context.Context,
	req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		resp = events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
		tranReq = entity.GetEmployeeRequest{}
		err     error
	)

	err = json.Unmarshal([]byte(req.Body), &tranReq)
	if err != nil {
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}

	respGen, err := usecase.GetEmployee(tranReq)
	if err != nil {
		log.Printf("Get Employee - ERROR In Get in Usecase: %s\n", err)
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	resp.StatusCode = http.StatusOK
	resp.Body = pkg.Response(entity.AppResponse{
		Error:     err,
		RequestID: tranReq.RequestID,
		Data:      respGen,
	})
	return resp, err

}

func CreateEmployee(ctx context.Context,
	req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		resp = events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
		tranReq = entity.CreateEmployeeRequest{}
		err     error
	)

	err = json.Unmarshal([]byte(req.Body), &tranReq)
	if err != nil {
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}

	respGen, err := usecase.CreateEmployee(tranReq)
	if err != nil {
		log.Printf("Get Employee - ERROR In Get in Usecase: %s\n", err)
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	resp.StatusCode = http.StatusOK
	resp.Body = pkg.Response(entity.AppResponse{
		Error:     err,
		RequestID: tranReq.RequestID,
		Data:      respGen,
	})
	return resp, err

}

func UpdateEmployee(ctx context.Context,
	req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		resp = events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
		tranReq = entity.UpdateEmployeeRequest{}
		err     error
	)

	err = json.Unmarshal([]byte(req.Body), &tranReq)
	if err != nil {
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}

	respGen, err := usecase.UpdateEmployee(tranReq)
	if err != nil {
		log.Printf("Get Employee - ERROR In Get in Usecase: %s\n", err)
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	resp.StatusCode = http.StatusOK
	resp.Body = pkg.Response(entity.AppResponse{
		Error:     err,
		RequestID: tranReq.RequestID,
		Data:      respGen,
	})
	return resp, err

}

func DeleteEmployee(ctx context.Context,
	req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		resp = events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
		tranReq = entity.DeleteEmployeeRequest{}
		err     error
	)

	err = json.Unmarshal([]byte(req.Body), &tranReq)
	if err != nil {
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}

	respGen, err := usecase.DeleteEmployee(tranReq)
	if err != nil {
		log.Printf("Get Employee - ERROR In Get in Usecase: %s\n", err)
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	resp.StatusCode = http.StatusOK
	resp.Body = pkg.Response(entity.AppResponse{
		Error:     err,
		RequestID: tranReq.RequestID,
		Data:      respGen,
	})
	return resp, err

}
