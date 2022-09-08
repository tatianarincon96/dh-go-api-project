package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
	"github.com/tatianarincon96/dh-go-api-project/internal/employee"
	"github.com/tatianarincon96/dh-go-api-project/pkg/store"
	"github.com/tatianarincon96/dh-go-api-project/pkg/web"
	"os"
	"strconv"
)

type employeeHandler struct {
	s employee.Service
}

// NewEmployeeHandler crea un nuevo controller de employeeos
func NewEmployeeHandler(s employee.Service) *employeeHandler {
	return &employeeHandler{
		s: s,
	}
}

func InitEmployeeHandler() *employeeHandler {
	// Inicializa employeeo
	list := store.LoadEmployees("../employees.json")
	repo := employee.NewRepository(list)
	service := employee.NewService(repo)
	return NewEmployeeHandler(service)
}

// GetAll obtiene todos los employeeos
func (h *employeeHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, _ := h.s.GetAll()
		web.Success(ctx, 200, employees)
	}
}

// GetByID obtiene un employeeo por su id
func (h *employeeHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		p, err := h.s.GetById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("employee not found"))
			return
		}
		web.Success(ctx, 200, p)
	}
}

// Post crear un employeeo nuevo
func (h *employeeHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		// Sigo con la ejecución - Validación de token ok!
		var newEmployee domain.Employee
		err := ctx.ShouldBindJSON(&newEmployee)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid employee"))
			return
		}

		p, err := h.s.Create(newEmployee)
		if err != nil {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		web.Success(ctx, 201, p)
	}
}

// Post crear un employeeo nuevo
func (h *employeeHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		var newEmployee domain.Employee
		err := ctx.ShouldBindJSON(&newEmployee)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid employee"))
			return
		}

		p, err := h.s.Update(newEmployee)
		if err != nil {
			web.Failure(ctx, 400, errors.New(err.Error()))
			return
		}
		web.Success(ctx, 201, p)
	}
}

// DeleteById elimina un employeeo por su id
func (h *employeeHandler) DeleteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = h.s.DeleteByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("employee not found"))
			return
		}
		web.Success(ctx, 200, "Employee deleted")
	}
}

func (h *employeeHandler) UpdateField() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Valido TOKEN
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}

		var newEmployee domain.Employee
		err := ctx.ShouldBindJSON(&newEmployee)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid employee"))
			return
		}
		p, err := h.s.UpdateField(newEmployee)
		if err != nil {
			web.Failure(ctx, 404, errors.New("employee not found to update"))
			return
		}
		web.Success(ctx, 200, p)
	}
}
