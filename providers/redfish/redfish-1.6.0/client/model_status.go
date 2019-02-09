/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes the status and health of a resource and its children.
type Status struct {
	Health Health `json:"Health,omitempty"`
	HealthRollup Health `json:"HealthRollup,omitempty"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	State State `json:"State,omitempty"`
}
