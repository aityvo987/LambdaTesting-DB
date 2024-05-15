package entity

import (
	"time"
)

type DeleteEmployeeRequest struct {
	RequestID   string       `json:"requestId" validate:"required"`
	RequestTime time.Time    `json:"requestTime"`
	EmployeeID  string       `json:"employeeId"`
	Data        EmployeeData `json:"data" validate:"required"`
}
