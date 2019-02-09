/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type PowerState string

// List of PowerState
const (
	ON PowerState = "On"
	OFF PowerState = "Off"
	POWERING_ON PowerState = "PoweringOn"
	POWERING_OFF PowerState = "PoweringOff"
)