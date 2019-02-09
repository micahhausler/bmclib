/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A maximum bandwidth allocation percentage for a Network Device Functions associated a port.
type NetDevFuncMaxBwAlloc struct {
	// The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	MaxBWAllocPercent int32 `json:"MaxBWAllocPercent,omitempty"`
	NetworkDeviceFunction IdRef `json:"NetworkDeviceFunction,omitempty"`
}
