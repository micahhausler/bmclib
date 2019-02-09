/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Stateless Address Automatic Configuration (SLAAC) parameters for this interface.
type StatelessAddressAutoConfiguration struct {
	// Indicates whether IPv4 SLAAC is enabled for this interface.
	IPv4AutoConfigEnabled bool `json:"IPv4AutoConfigEnabled,omitempty"`
	// Indicates whether IPv6 SLAAC is enabled for this interface.
	IPv6AutoConfigEnabled bool `json:"IPv6AutoConfigEnabled,omitempty"`
}
