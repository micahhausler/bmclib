/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Settings for a network protocol associated with a manager.
type NtpProtocol struct {
	// Indicates to which NTP servers this manager is subscribed.
	NTPServers []string `json:"NTPServers,omitempty"`
	// Indicates the protocol port.
	Port int32 `json:"Port,omitempty"`
	// Indicates if the protocol is enabled or disabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`
}
