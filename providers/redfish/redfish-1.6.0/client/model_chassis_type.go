/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type ChassisType string

// List of ChassisType
const (
	RACK ChassisType = "Rack"
	BLADE ChassisType = "Blade"
	ENCLOSURE ChassisType = "Enclosure"
	STAND_ALONE ChassisType = "StandAlone"
	RACK_MOUNT ChassisType = "RackMount"
	CARD ChassisType = "Card"
	CARTRIDGE ChassisType = "Cartridge"
	ROW ChassisType = "Row"
	POD ChassisType = "Pod"
	EXPANSION ChassisType = "Expansion"
	SIDECAR ChassisType = "Sidecar"
	ZONE ChassisType = "Zone"
	SLED ChassisType = "Sled"
	SHELF ChassisType = "Shelf"
	DRAWER ChassisType = "Drawer"
	MODULE ChassisType = "Module"
	COMPONENT ChassisType = "Component"
	IP_BASED_DRIVE ChassisType = "IPBasedDrive"
	RACK_GROUP ChassisType = "RackGroup"
	STORAGE_ENCLOSURE ChassisType = "StorageEnclosure"
	OTHER ChassisType = "Other"
)