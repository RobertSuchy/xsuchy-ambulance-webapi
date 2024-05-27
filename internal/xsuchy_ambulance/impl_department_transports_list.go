package xsuchy_ambulance

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "github.com/RobertSuchy/xsuchy-ambulance-webapi/internal/db_service"
)

// CreateTransport - Saves new transport into the list
func (this *implDepartmentTransportsListAPI) CreateTransport(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_transport")
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

  db, ok := value.(db_service.DbService[Transport])
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

  transport := Transport{}
  err := ctx.BindJSON(&transport)
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

  if transport.Id == "" {
      transport.Id = uuid.New().String()
  }

  err = db.CreateDocument(ctx, transport.Id, &transport)
  switch err {
  case nil:
      ctx.JSON(
          http.StatusCreated,
          transport,
      )
  case db_service.ErrConflict:
      ctx.JSON(
          http.StatusConflict,
          gin.H{
              "status":  "Conflict",
              "message": "Transport already exists",
              "error":   err.Error(),
          },
      )
  default:
      ctx.JSON(
          http.StatusBadGateway,
          gin.H{
              "status":  "Bad Gateway",
              "message": "Failed to create transport in database",
              "error":   err.Error(),
          },
      )
  }
}

// DeleteTransport - Deletes specific transport
func (this *implDepartmentTransportsListAPI) DeleteTransport(ctx *gin.Context) {
    value, exists := ctx.Get("db_service_transport")
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

    db, ok := value.(db_service.DbService[Transport])
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

    transportID := ctx.Param("transportId")
    if transportID == "" {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Transport ID is required",
                "error":   "transport ID not provided",
            })
        return
    }

    err := db.DeleteDocument(ctx, transportID)
    switch err {
    case nil:
        ctx.JSON(
            http.StatusNoContent,
            nil,
        )
    case db_service.ErrNotFound:
        ctx.JSON(
            http.StatusNotFound,
            gin.H{
                "status":  "Not Found",
                "message": "Transport not found",
                "error":   err.Error(),
            },
        )
    default:
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to delete transport from database",
                "error":   err.Error(),
            },
        )
    }	
}

// GetTransport - Provides details about a specific transport
func (this *implDepartmentTransportsListAPI) GetTransport(ctx *gin.Context) {
	value, exists := ctx.Get("db_service_transport")
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

    db, ok := value.(db_service.DbService[Transport])
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

    transportID := ctx.Param("transportId")
    if transportID == "" {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Transport ID is required",
                "error":   "transport ID not provided",
            })
        return
    }

    transport, err := db.FindDocument(ctx, transportID)
    if err != nil {
        if err == db_service.ErrNotFound {
            ctx.JSON(
                http.StatusNotFound,
                gin.H{
                    "status":  "Not Found",
                    "message": "Transport not found",
                    "error":   err.Error(),
                })
        } else {
            ctx.JSON(
                http.StatusBadGateway,
                gin.H{
                    "status":  "Bad Gateway",
                    "message": "Failed to fetch transport from database",
                    "error":   err.Error(),
                })
        }
        return
    }

    ctx.JSON(http.StatusOK, transport)
}

// GetTransportsList - Provides the department transports list
func (this *implDepartmentTransportsListAPI) GetTransportsList(ctx *gin.Context) {
    value, exists := ctx.Get("db_service_transport")
    if (!exists) {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db not found",
                "error":   "db not found",
            })
        return
    }

    db, ok := value.(db_service.DbService[Transport])
    if (!ok) {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db context is not of required type",
                "error":   "cannot cast db context to db_service.DbService",
            })
        return
    }

    departmentID := ctx.Param("departmentId")
    if (departmentID == "") {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Department ID is required",
                "error":   "department ID not provided",
            })
        return
    }

    transports, err := db.GetAllDocuments(ctx)
    if (err != nil) {
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to fetch transports from database",
                "error":   err.Error(),
            },
        )
        return
    }

    var filteredTransports []Transport
    for _, transport := range transports {
        if transport.FromDepartmentId == departmentID || transport.ToDepartmentId == departmentID {
            filteredTransports = append(filteredTransports, *transport)
        }
    }

    ctx.JSON(http.StatusOK, filteredTransports)	
}

// UpdateTransport - Updates specific transport
func (this *implDepartmentTransportsListAPI) UpdateTransport(ctx *gin.Context) {
    value, exists := ctx.Get("db_service_transport")
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

    db, ok := value.(db_service.DbService[Transport])
    if (!ok) {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db context is not of required type",
                "error":   "cannot cast db context to db_service.DbService",
            })
        return
    }

    transportID := ctx.Param("transportId")
    if (transportID == "") {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Transport ID is required",
                "error":   "transport ID not provided",
            })
        return
    }

    transport := Transport{}
    err := ctx.BindJSON(&transport)
    if (err != nil) {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Invalid request body",
                "error":   err.Error(),
            })
        return
    }

    if transport.Id != transportID {
        ctx.JSON(
            http.StatusForbidden,
            gin.H{
                "status":  "Forbidden",
                "message": "Transport ID in the path and request body do not match",
            })
        return
    }

    err = db.UpdateDocument(ctx, transportID, &transport)
    switch err {
    case nil:
        ctx.JSON(
            http.StatusOK,
            transport,
        )
    case db_service.ErrNotFound:
        ctx.JSON(
            http.StatusNotFound,
            gin.H{
                "status":  "Not Found",
                "message": "Transport not found",
                "error":   err.Error(),
            },
        )
    default:
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to update transport in database",
                "error":   err.Error(),
            },
        )
    }
}
