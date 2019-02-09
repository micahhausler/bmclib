/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The HTTP and JSON payload details for this Task.
type Payload struct {
	// This represents the HTTP headers used in the operation of this Task.
	HttpHeaders []string `json:"HttpHeaders,omitempty"`
	// The HTTP operation to perform to execute this Task.
	HttpOperation string `json:"HttpOperation,omitempty"`
	// This property contains the JSON payload to use in the execution of this Task.
	JsonBody string `json:"JsonBody,omitempty"`
	// The URI of the target for this task.
	TargetUri string `json:"TargetUri,omitempty"`
}
