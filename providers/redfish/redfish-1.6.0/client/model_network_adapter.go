/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A NetworkAdapter represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	Assembly IdRef `json:"Assembly,omitempty"`
	// The set of network controllers ASICs that make up this NetworkAdapter.
	Controllers []Controllers `json:"Controllers,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The manufacturer or OEM of this network adapter.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// The model string for this network adapter.
	Model string `json:"Model,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	NetworkDeviceFunctions IdRef `json:"NetworkDeviceFunctions,omitempty"`
	NetworkPorts IdRef `json:"NetworkPorts,omitempty"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	// Part number for this network adapter.
	PartNumber string `json:"PartNumber,omitempty"`
	// The manufacturer SKU for this network adapter.
	SKU string `json:"SKU,omitempty"`
	// The serial number for this network adapter.
	SerialNumber string `json:"SerialNumber,omitempty"`
	Status Status `json:"Status,omitempty"`
}
