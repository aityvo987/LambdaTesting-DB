package handler

import (
	"context"
	"encoding/json"
	"errors"
	entity "lambda-test/internal/entity"
	pkg "lambda-test/internal/pkg"
	usecase "lambda-test/internal/usecase"
	"log"
	"net/http"

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
		log.Printf("Get Employees - Cannot read Request: %s\n", err)
		err = errors.New(pkg.ReqError.Code())
		resp.Body = pkg.Response(entity.AppResponse{
			Error: err,
		})
		return resp, nil
	}
	log.Printf("Get Employees - Calling to Usecase\n")
	respGet, err := usecase.GetEmployees(tranReq)
	if err != nil {
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
		Data:      respGet.Employees,
	})
	return resp, nil

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
		log.Printf("Get Employee - Cannot read Request: %s\n", err)
		err = errors.New(pkg.ReqError.Code())
		resp.Body = pkg.Response(entity.AppResponse{
			Error: err,
		})
		return resp, nil
	}
	log.Printf("Get Employee - Calling to Usecase\n")
	respGen, err := usecase.GetEmployee(tranReq)
	if err != nil {
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
	return resp, nil

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
		log.Printf("Create Employee - Cannot read Request: %s\n", err)
		err = errors.New(pkg.ReqError.Code())
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	log.Printf("Create Employee - Calling to Usecase\n")
	err = usecase.CreateEmployee(tranReq)
	if err != nil {
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
	})
	return resp, nil

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
		log.Printf("Update Employee - Cannot read Request: %s\n", err)
		err = errors.New(pkg.ReqError.Code())
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	log.Printf("Update Employee - Calling to Usecase\n")
	err = usecase.UpdateEmployee(tranReq)
	if err != nil {
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
	})
	return resp, nil

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
		log.Printf("Delete Employee - Cannot read Request: %s\n", err)
		err = errors.New(pkg.ReqError.Code())
		resp.Body = pkg.Response(entity.AppResponse{
			Error:     err,
			RequestID: tranReq.RequestID,
		})
		return resp, nil
	}
	log.Printf("Delete Employee - Calling to Usecase\n")
	err = usecase.DeleteEmployee(tranReq)
	if err != nil {
		log.Printf("Delete Employee - ERROR In Get in Usecase: %s\n", err)
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
	})
	return resp, nil

}
