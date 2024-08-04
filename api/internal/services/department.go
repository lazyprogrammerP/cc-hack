package services

import (
	"context"

	"github.com/lazyprogrammerP/cc-hack/internal/sqlc"
)

type DepartmentService struct {
	queries *sqlc.Queries
}

func NewDepartmentService(queries *sqlc.Queries) *DepartmentService {
	return &DepartmentService{queries: queries}
}

func (s *DepartmentService) GetAllDepartments(ctx context.Context) ([]sqlc.Department, error) {
	departments, err := s.queries.GetAllDepartments(ctx)
	if err != nil {
		return nil, err
	}

	return departments, nil

}

func (s *DepartmentService) CreateDepartment(ctx context.Context, names []string) ([]sqlc.Department, error) {
	departments, err := s.queries.CreateDepartment(ctx, names)
	if err != nil {
		return nil, err
	}

	return departments, nil
}
