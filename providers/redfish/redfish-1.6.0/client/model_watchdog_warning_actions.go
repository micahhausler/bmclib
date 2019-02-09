/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WatchdogWarningActions : The enumerations of WatchdogWarningActions specify the choice of action to take when the Host Watchdog Timer is close (typically 3-10 seconds) to reaching its timeout value.
type WatchdogWarningActions string

// List of WatchdogWarningActions
const (
	NONE WatchdogWarningActions = "None"
	DIAGNOSTIC_INTERRUPT WatchdogWarningActions = "DiagnosticInterrupt"
	SMI WatchdogWarningActions = "SMI"
	MESSAGING_INTERRUPT WatchdogWarningActions = "MessagingInterrupt"
	SCI WatchdogWarningActions = "SCI"
	OEM WatchdogWarningActions = "OEM"
)