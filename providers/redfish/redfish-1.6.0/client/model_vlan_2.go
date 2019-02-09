/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Vlan2 struct {
	// This indicates if this VLAN is enabled.
	VLANEnable bool `json:"VLANEnable,omitempty"`
	VLANId float32 `json:"VLANId,omitempty"`
}
