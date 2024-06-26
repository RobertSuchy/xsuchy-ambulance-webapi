/*
 * Department Transports Api
 *
 * Department Transports management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xsuchy@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package xsuchy_ambulance

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type DepartmentTransportsListAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateTransport - Saves new transport into the list
   CreateTransport(ctx *gin.Context)

    // DeleteTransport - Deletes specific transport
   DeleteTransport(ctx *gin.Context)

    // GetTransport - Provides details about a specific transport
   GetTransport(ctx *gin.Context)

    // GetTransportsList - Provides the department transports list
   GetTransportsList(ctx *gin.Context)

    // UpdateTransport - Updates specific transport
   UpdateTransport(ctx *gin.Context)

 }

// partial implementation of DepartmentTransportsListAPI - all functions must be implemented in add on files
type implDepartmentTransportsListAPI struct {

}

func newDepartmentTransportsListAPI() DepartmentTransportsListAPI {
  return &implDepartmentTransportsListAPI{}
}

func (this *implDepartmentTransportsListAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/transports", this.CreateTransport)
  routerGroup.Handle( http.MethodDelete, "/transports/:transportId", this.DeleteTransport)
  routerGroup.Handle( http.MethodGet, "/transports/:transportId", this.GetTransport)
  routerGroup.Handle( http.MethodGet, "/departments/:departmentId/transports", this.GetTransportsList)
  routerGroup.Handle( http.MethodPut, "/transports/:transportId", this.UpdateTransport)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateTransport - Saves new transport into the list
// func (this *implDepartmentTransportsListAPI) CreateTransport(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteTransport - Deletes specific transport
// func (this *implDepartmentTransportsListAPI) DeleteTransport(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetTransport - Provides details about a specific transport
// func (this *implDepartmentTransportsListAPI) GetTransport(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetTransportsList - Provides the department transports list
// func (this *implDepartmentTransportsListAPI) GetTransportsList(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateTransport - Updates specific transport
// func (this *implDepartmentTransportsListAPI) UpdateTransport(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

