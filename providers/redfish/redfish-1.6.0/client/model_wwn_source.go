/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type WwnSource string

// List of WWNSource
const (
	CONFIGURED_LOCALLY WwnSource = "ConfiguredLocally"
	PROVIDED_BY_FABRIC WwnSource = "ProvidedByFabric"
)