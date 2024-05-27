package xsuchy_ambulance

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// CreateTransport - Saves new transport into the list
func (this *implDepartmentTransportsListAPI) CreateTransport(ctx *gin.Context) {
	updateDepartmentFunc(ctx, func(c *gin.Context, department *Department) (*Department, interface{}, int) {
		var transport Transport

		if err := c.ShouldBindJSON(&transport); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if transport.PatientId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Patient ID is required",
			}, http.StatusBadRequest
		}

		if transport.Id == "" || transport.Id == "@new" {
			transport.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(department.TransportsList, func(existing Transport) bool {
			return transport.Id == existing.Id || transport.PatientId == existing.PatientId
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Transport entry already exists",
			}, http.StatusConflict
		}

		department.TransportsList = append(department.TransportsList, transport)

		entryIndx := slices.IndexFunc(department.TransportsList, func(existing Transport) bool {
			return transport.Id == existing.Id
		})
		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save transport entry",
			}, http.StatusInternalServerError
		}
		return department, department.TransportsList[entryIndx], http.StatusOK
	})
}

// DeleteTransport - Deletes specific transport
func (this *implDepartmentTransportsListAPI) DeleteTransport(ctx *gin.Context) {
	updateDepartmentFunc(ctx, func(c *gin.Context, department *Department) (*Department, interface{}, int) {
		transportId := ctx.Param("transportId")

		if transportId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Transport ID is required",
			}, http.StatusBadRequest
		}

		transportIndx := slices.IndexFunc(department.TransportsList, func(existing Transport) bool {
			return transportId == existing.Id
		})

		if transportIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Transport not found",
			}, http.StatusNotFound
		}

		department.TransportsList = append(department.TransportsList[:transportIndx], department.TransportsList[transportIndx+1:]...)
		return department, nil, http.StatusNoContent
	})
}

// GetTransport - Provides details about a specific transport
func (this *implDepartmentTransportsListAPI) GetTransport(ctx *gin.Context) {
	updateDepartmentFunc(ctx, func(c *gin.Context, department *Department) (*Department, interface{}, int) {
		transportId := ctx.Param("transportId")

		if transportId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Transport ID is required",
			}, http.StatusBadRequest
		}

		transportIndx := slices.IndexFunc(department.TransportsList, func(existing Transport) bool {
			return transportId == existing.Id
		})

		if transportIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Transport not found",
			}, http.StatusNotFound
		}

		return nil, department.TransportsList[transportIndx], http.StatusOK
	})
}

// GetTransportsList - Provides the department transports list
func (this *implDepartmentTransportsListAPI) GetTransportsList(ctx *gin.Context) {
	updateDepartmentFunc(ctx, func(c *gin.Context, department *Department) (*Department, interface{}, int) {
		result := department.TransportsList
		if result == nil {
			result = []Transport{}
		}

		return nil, result, http.StatusOK
	})
}

// UpdateTransport - Updates specific transport
func (this *implDepartmentTransportsListAPI) UpdateTransport(ctx *gin.Context) {
	updateDepartmentFunc(ctx, func(c *gin.Context, department *Department) (*Department, interface{}, int) {
		var transport Transport

		if err := c.ShouldBindJSON(&transport); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		transportId := ctx.Param("transportId")

		if transportId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Transport ID is required",
			}, http.StatusBadRequest
		}

		transportIndx := slices.IndexFunc(department.TransportsList, func(existing Transport) bool {
			return transportId == existing.Id
		})

		if transportIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Transport not found",
			}, http.StatusNotFound
		}

		if transport.PatientId != "" {
			department.TransportsList[transportIndx].PatientId = transport.PatientId
		}

		if transport.PatientName != "" {
			department.TransportsList[transportIndx].PatientName = transport.PatientName
		}

		if transport.FromDepartmentId != "" {
			department.TransportsList[transportIndx].FromDepartmentId = transport.FromDepartmentId
		}

		if transport.ToDepartmentId != "" {
			department.TransportsList[transportIndx].ToDepartmentId = transport.ToDepartmentId
		}

		if !transport.ScheduledDateTime.IsZero() {
			department.TransportsList[transportIndx].ScheduledDateTime = transport.ScheduledDateTime
		}

		if transport.EstimatedDurationMinutes > 0 {
			department.TransportsList[transportIndx].EstimatedDurationMinutes = transport.EstimatedDurationMinutes
		}

		department.TransportsList[transportIndx].MobilityStatus = transport.MobilityStatus

		return department, department.TransportsList[transportIndx], http.StatusOK
	})
}
