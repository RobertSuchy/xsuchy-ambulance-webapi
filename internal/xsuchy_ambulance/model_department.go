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

type Department struct {

	// Unique ID of the department
	Id string `json:"id"`

	// Name of the department
	Name string `json:"name"`

	// City where the department is located
	City string `json:"city"`
}
