/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WatchdogTimeoutActions : The enumerations of WatchdogTimeoutActions specify the choice of action to take when the Host Watchdog Timer reaches its timeout value.
type WatchdogTimeoutActions string

// List of WatchdogTimeoutActions
const (
	NONE WatchdogTimeoutActions = "None"
	RESET_SYSTEM WatchdogTimeoutActions = "ResetSystem"
	POWER_CYCLE WatchdogTimeoutActions = "PowerCycle"
	POWER_DOWN WatchdogTimeoutActions = "PowerDown"
	OEM WatchdogTimeoutActions = "OEM"
)