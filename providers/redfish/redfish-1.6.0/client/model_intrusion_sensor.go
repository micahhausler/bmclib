/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type IntrusionSensor string

// List of IntrusionSensor
const (
	NORMAL IntrusionSensor = "Normal"
	HARDWARE_INTRUSION IntrusionSensor = "HardwareIntrusion"
	TAMPERING_DETECTED IntrusionSensor = "TamperingDetected"
)