/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Specifies the usage of the parameter in the calculation.
type CalculationParamsType struct {
	// The metric property used to store the results of the calculation.
	ResultMetric string `json:"ResultMetric,omitempty"`
	// The metric property used as the input into the calculation.
	SourceMetric string `json:"SourceMetric,omitempty"`
}
