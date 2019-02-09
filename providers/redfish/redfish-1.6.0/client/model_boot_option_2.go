/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the schema definition for the BootOption resource. It represents the properties of a bootable device available in the System.
type BootOption2 struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	Alias BootSource `json:"Alias,omitempty"`
	// A flag that shows if the Boot Option is enabled.
	BootOptionEnabled bool `json:"BootOptionEnabled,omitempty"`
	// The unique boot option string that is referenced in the BootOrder.
	BootOptionReference string `json:"BootOptionReference"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The user-readable display string of the Boot Option.
	DisplayName string `json:"DisplayName,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	// The ID(s) of the resources associated with this Boot Option.
	RelatedItem []IdRef `json:"RelatedItem,omitempty"`
	// The number of items in a collection.
	RelatedItemodataCount int32 `json:"RelatedItem@odata.count,omitempty"`
	// The UEFI device path used to access this UEFI Boot Option.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`
}
