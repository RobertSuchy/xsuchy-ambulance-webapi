package xsuchy_ambulance

import (
  "net/http"

  "github.com/gin-gonic/gin"
//   "github.com/RobertSuchy/xsuchy-ambulance-webapi/internal/db_service"
)

// GetMobilityStatusList - Provides the list of mobility statuses
func (this *implMobilityStatusListAPI) GetMobilityStatusList(ctx *gin.Context) {
    mobilityStatuses := []MobilityStatus{
        {
            Value:       "na vozíku",
            Code:        "wheelchair",
            Description: "Patient requires a wheelchair for transport",
        },
        {
            Value:       "mobilný",
            Code:        "walking",
            Description: "Patient is able to walk",
        },
        {
            Value:       "ležiaci",
            Code:        "bedridden",
            Description: "Patient is bedridden and requires stretcher for transport",
        },
    }

    // Return the hardcoded mobility statuses as JSON response
    ctx.JSON(http.StatusOK, mobilityStatuses)

    //   value, exists := ctx.Get("db_service_mobility_status")
//   if !exists {
//       ctx.JSON(
//           http.StatusInternalServerError,
//           gin.H{
//               "status":  "Internal Server Error",
//               "message": "db not found",
//               "error":   "db not found",
//           })
//       return
//   }

//   db, ok := value.(db_service.DbService[MobilityStatus])
//   if !ok {
//       ctx.JSON(
//           http.StatusInternalServerError,
//           gin.H{
//               "status":  "Internal Server Error",
//               "message": "db context is not of required type",
//               "error":   "cannot cast db context to db_service.DbService",
//           })
//       return
//   }

//   mobilityStatuses, err := db.GetAllDocuments(ctx)
//   if err != nil {
//       ctx.JSON(
//           http.StatusBadGateway,
//           gin.H{
//               "status":  "Bad Gateway",
//               "message": "Failed to fetch mobility statuses from database",
//               "error":   err.Error(),
//           },
//       )
//       return
//   }

//   ctx.JSON(http.StatusOK, mobilityStatuses)  
}
