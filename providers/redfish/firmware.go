package redfish

import (
	"context"

	"github.com/stmcginnis/gofish/redfish"
)

func (c *Conn) GetFirmwareInventory(ctx context.Context) (inventory map[string]string, err error) {
	redfish.GetUpdateService(c.conn,"/redfish/v1/UpdateService/FirmwareInventory/%s' % " )
	return nil, nil
}
