/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object contains the boot information for the current resource.
type Boot struct {
	// This property is the BootOptionReference of the Boot Option to perform a one time boot from when BootSourceOverrideTarget is UefiBootNext.
	BootNext string `json:"BootNext,omitempty"`
	BootOptions IdRef `json:"BootOptions,omitempty"`
	// Ordered array of BootOptionReference strings representing the persistent Boot Order associated with this computer system.
	BootOrder []string `json:"BootOrder,omitempty"`
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:"BootSourceOverrideEnabled,omitempty"`
	BootSourceOverrideMode BootSourceOverrideMode `json:"BootSourceOverrideMode,omitempty"`
	BootSourceOverrideTarget BootSource `json:"BootSourceOverrideTarget,omitempty"`
	// This property is the UEFI Device Path of the device to boot from when BootSourceOverrideTarget is UefiTarget.
	UefiTargetBootSourceOverride string `json:"UefiTargetBootSourceOverride,omitempty"`
}
