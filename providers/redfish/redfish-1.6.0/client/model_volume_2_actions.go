/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The available actions for this resource.
type Volume2Actions struct {
	VolumeInitialize Initialize `json:"#Volume.Initialize,omitempty"`
	// The available OEM specific actions for this resource.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
}
