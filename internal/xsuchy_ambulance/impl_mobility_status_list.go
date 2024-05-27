package xsuchy_ambulance

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// GetMobilityStatusList - Provides the list of mobility statuses
func (this *implMobilityStatusListAPI) GetMobilityStatusList(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}
