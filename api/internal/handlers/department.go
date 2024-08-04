package handlers

import (
	"net/http"

	"github.com/lazyprogrammerP/cc-hack/internal/models"
	"github.com/lazyprogrammerP/cc-hack/internal/services"
	"github.com/lazyprogrammerP/cc-hack/internal/sqlc"
	"github.com/lazyprogrammerP/cc-hack/pkg/utils"
	"github.com/rs/zerolog/log"
)

type DepartmentHandler struct {
	service *services.DepartmentService
}

func NewDepartmentHandler(queries *sqlc.Queries) *DepartmentHandler {
	departmentService := services.NewDepartmentService(queries)
	return &DepartmentHandler{service: departmentService}
}

func (h *DepartmentHandler) GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	departments, err := h.service.GetAllDepartments(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get all departments")
		utils.EncodeResponse(w, http.StatusInternalServerError, &models.Error{Error: err.Error()})
		return
	}

	utils.EncodeResponse(w, http.StatusOK, &models.GetAllDepartmentsResponse{Departments: departments})
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	payload, err := utils.DecodeRequest[models.CreateDepartmentRequest](r)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode the request")
		utils.EncodeResponse(w, http.StatusBadRequest, &models.Error{Error: err.Error()})
		return
	}

	departments, err := h.service.CreateDepartment(r.Context(), payload.Names)
	if err != nil {
		log.Error().Err(err).Msg("failed to create the department")
		utils.EncodeResponse(w, http.StatusInternalServerError, &models.Error{Error: err.Error()})
		return
	}

	utils.EncodeResponse(w, http.StatusOK, &models.CreateDepartmentResponse{Departments: departments})
}
