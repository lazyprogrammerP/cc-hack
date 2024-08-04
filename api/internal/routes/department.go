package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/lazyprogrammerP/cc-hack/internal/handlers"
	"github.com/lazyprogrammerP/cc-hack/internal/sqlc"
)

func RegisterDepartmentRoutes(queries *sqlc.Queries) chi.Router {
	r := chi.NewRouter()

	departmentHandler := handlers.NewDepartmentHandler(queries)

	r.Get("/", departmentHandler.GetAllDepartments)
	r.Post("/", departmentHandler.CreateDepartment)

	return r
}
