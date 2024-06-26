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

// MobilityStatus - Describes the mobility status of the patient for transport
type MobilityStatus struct {

	Value string `json:"value"`

	Code string `json:"code,omitempty"`

	Description string `json:"description,omitempty"`
}
