/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RedfishErrorError struct {
	// An array of message objects describing one or more error message(s).
	MessageExtendedInfo []Message `json:"@Message.ExtendedInfo,omitempty"`
	// A string indicating a specific MessageId from the message registry.
	Code string `json:"code,omitempty"`
	// A human-readable error message corresponding to the message in the message registry.
	Message string `json:"message,omitempty"`
}
