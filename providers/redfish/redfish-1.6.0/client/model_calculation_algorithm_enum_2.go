/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CalculationAlgorithmEnum2 : Specifies the function to apply to the list of metric properties.
type CalculationAlgorithmEnum2 string

// List of CalculationAlgorithmEnum_2
const (
	AVERAGE CalculationAlgorithmEnum2 = "Average"
	MAXIMUM CalculationAlgorithmEnum2 = "Maximum"
	MINIMUM CalculationAlgorithmEnum2 = "Minimum"
	SUMMATION CalculationAlgorithmEnum2 = "Summation"
)