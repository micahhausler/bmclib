/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type PrivilegeType string

// List of PrivilegeType
const (
	LOGIN PrivilegeType = "Login"
	CONFIGURE_MANAGER PrivilegeType = "ConfigureManager"
	CONFIGURE_USERS PrivilegeType = "ConfigureUsers"
	CONFIGURE_SELF PrivilegeType = "ConfigureSelf"
	CONFIGURE_COMPONENTS PrivilegeType = "ConfigureComponents"
)