/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type IPv4AddressOrigin string

// List of IPv4AddressOrigin
const (
	STATIC IPv4AddressOrigin = "Static"
	DHCP IPv4AddressOrigin = "DHCP"
	BOOTP IPv4AddressOrigin = "BOOTP"
	I_PV4_LINK_LOCAL IPv4AddressOrigin = "IPv4LinkLocal"
)