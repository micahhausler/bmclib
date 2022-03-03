package redfish

import "github.com/pkg/errors"

var (
	ErrRedfishChassisOdataID    = errors.New("no compatible Chassis Odata IDs identified")
	ErrRedfishSystemOdataID     = errors.New("no compatible System Odata IDs identified")
	ErrRedfishManagerOdataID    = errors.New("no compatible Manager Odata IDs identified")
	ErrRedfishServiceNil        = errors.New("redfish connection returned a nil redfish Service object")
	ErrRedfishSoftwareInventory = errors.New("error collecting redfish software inventory")
)
