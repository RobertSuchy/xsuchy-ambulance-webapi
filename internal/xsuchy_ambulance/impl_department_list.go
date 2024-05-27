package xsuchy_ambulance

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "github.com/RobertSuchy/xsuchy-ambulance-webapi/internal/db_service"
)

// CreateDepartment - Creates a new department
func (this *implDepartmentListAPI) CreateDepartment(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
    if !exists {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db not found",
                "error":   "db not found",
            })
        return
    }

    db, ok := value.(db_service.DbService[Department])
    if !ok {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db context is not of required type",
                "error":   "cannot cast db context to db_service.DbService",
            })
        return
    }

    department := Department{}
    err := ctx.BindJSON(&department)
    if err != nil {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Invalid request body",
                "error":   err.Error(),
            })
        return
    }

    if department.Id == "" {
        department.Id = uuid.New().String()
    }

    err = db.CreateDocument(ctx, department.Id, &department)

    switch err {
    case nil:
        ctx.JSON(
            http.StatusCreated,
            department,
        )
    case db_service.ErrConflict:
        ctx.JSON(
            http.StatusConflict,
            gin.H{
                "status":  "Conflict",
                "message": "Department already exists",
                "error":   err.Error(),
            },
        )
    default:
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to create department in database",
                "error":   err.Error(),
            },
        )
    }
}

// GetAllDepartments - Provides the list of all departments
func (this *implDepartmentListAPI) GetAllDepartments(ctx *gin.Context) {
    updateDepartmentFunc(ctx, func(
        ctx *gin.Context,
        department *Department,
    ) (updatedDepartment *Department, responseContent interface{}, status int) {
        value, exists := ctx.Get("db_service")
        if !exists {
            return nil, gin.H{
                "status":  "Internal Server Error",
                "message": "db_service not found",
                "error":   "db_service not found",
            }, http.StatusInternalServerError
        }

        db, ok := value.(db_service.DbService[Department])
        if !ok {
            return nil, gin.H{
                "status":  "Internal Server Error",
                "message": "db_service context is not of type db_service.DbService",
                "error":   "cannot cast db_service context to db_service.DbService",
            }, http.StatusInternalServerError
        }

        departments, err := db.GetAllDocuments(ctx)
        if err != nil {
            return nil, gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to retrieve departments from database",
                "error":   err.Error(),
            }, http.StatusBadGateway
        }

        if departments == nil {
            departments = []Department{}
        }

        return nil, departments, http.StatusOK
    })
}

// GetDepartment - Provides details about a specific department
func (this *implDepartmentListAPI) GetDepartment(ctx *gin.Context) {
    updateDepartmentFunc(ctx, func(
        ctx *gin.Context,
        department *Department,
    ) (updatedDepartment *Department, responseContent interface{}, status int) {
        return nil, department, http.StatusOK
    })
}
