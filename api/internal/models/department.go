package models

import "github.com/lazyprogrammerP/cc-hack/internal/sqlc"

type GetAllDepartmentsResponse struct {
	Departments []sqlc.Department `json:"departments"`
}

type CreateDepartmentRequest struct {
	Names []string `json:"names"`
}

type CreateDepartmentResponse struct {
	Departments []sqlc.Department `json:"departments"`
}
