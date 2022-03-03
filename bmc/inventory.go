package bmc

import (
	"context"
	"fmt"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// InventoryGetter defines methods to retrieve device hardware and firmware inventory
type InventoryGetter interface {
	GetInventory(ctx context.Context) (device *devices.Device, err error)
}

// GetInventoryFromInterfaces is a pass through to library function
func GetInventoryFromInterfaces(ctx context.Context, generic []interface{}) (device *devices.Device, err error) {
	getters := make([]InventoryGetter, 0)
	for _, elem := range generic {
		switch p := elem.(type) {
		case InventoryGetter:
			getters = append(getters, p)
		default:
			e := fmt.Sprintf("not a InventoryGetter implementation: %T", p)
			err = multierror.Append(err, errors.New(e))
		}
	}

	if len(getters) == 0 {
		return nil, multierror.Append(err, errors.New("no InventoryGetter implementations found"))
	}

	return GetInventory(ctx, getters)
}

// GetInventory returns the complete hardware and firmware inventory
func GetInventory(ctx context.Context, p []InventoryGetter) (inventory *devices.Device, err error) {
Loop:
	for _, elem := range p {
		if elem == nil {
			continue
		}
		select {
		case <-ctx.Done():
			err = multierror.Append(err, ctx.Err())
			break Loop
		default:
			version, vErr := elem.GetInventory(ctx)
			if vErr != nil {
				err = multierror.Append(err, vErr)
				continue
			}
			return version, nil
		}
	}

	return inventory, multierror.Append(err, errors.New("failed to get device inventory"))
}
