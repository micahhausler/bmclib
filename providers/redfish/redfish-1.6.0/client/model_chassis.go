/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The Chassis schema represents the physical components of a system.  This resource represents the sheet-metal confined spaces and logical zones such as racks, enclosures, chassis and all other containers. Subsystems (like sensors) that operate outside of a system's data plane (meaning the resources are not accessible to software running on the system) are linked either directly or indirectly through this resource.
type Chassis struct {
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
	// The user assigned asset tag of this chassis.
	AssetTag string `json:"AssetTag,omitempty"`
	ChassisType ChassisType `json:"ChassisType"`
	// The depth of the chassis.
	DepthMm float32 `json:"DepthMm,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The height of the chassis.
	HeightMm float32 `json:"HeightMm,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	IndicatorLED IndicatorLed `json:"IndicatorLED,omitempty"`
	Links Links2 `json:"Links,omitempty"`
	Location Location2 `json:"Location,omitempty"`
	LogServices IdRef `json:"LogServices,omitempty"`
	// The manufacturer of this chassis.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// The model number of the chassis.
	Model string `json:"Model,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	NetworkAdapters IdRef `json:"NetworkAdapters,omitempty"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	PCIeSlots IdRef `json:"PCIeSlots,omitempty"`
	// The part number of the chassis.
	PartNumber string `json:"PartNumber,omitempty"`
	PhysicalSecurity PhysicalSecurity `json:"PhysicalSecurity,omitempty"`
	Power IdRef `json:"Power,omitempty"`
	PowerState PowerState `json:"PowerState,omitempty"`
	// The SKU of the chassis.
	SKU string `json:"SKU,omitempty"`
	// The serial number of the chassis.
	SerialNumber string `json:"SerialNumber,omitempty"`
	Status Status `json:"Status,omitempty"`
	Thermal IdRef `json:"Thermal,omitempty"`
	UUID string `json:"UUID,omitempty"`
	// The weight of the chassis.
	WeightKg float32 `json:"WeightKg,omitempty"`
	// The width of the chassis.
	WidthMm float32 `json:"WidthMm,omitempty"`
}
