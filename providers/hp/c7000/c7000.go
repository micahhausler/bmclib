package c7000

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclib/errors"
	"github.com/bmc-toolbox/bmclib/internal/httpclient"
	"github.com/bmc-toolbox/bmclib/internal/sshclient"
	"github.com/bmc-toolbox/bmclib/providers/hp"

	// this make possible to setup logging and properties at any stage
	_ "github.com/bmc-toolbox/bmclib/logging"
)

const (
	// BMCType defines the bmc model that is supported by this package
	BMCType = "c7000"
)

// C7000 holds the status and properties of a connection to a BladeSystem device
type C7000 struct {
	ip         string
	username   string
	password   string
	XMLToken   string //required to send SOAP XML payloads
	httpClient *http.Client
	sshClient  *sshclient.SSHClient
	Rimp       *hp.Rimp
}

// New returns a connection to C7000
func New(ip string, username string, password string) (chassis *C7000, err error) {
	client, err := httpclient.Build()
	if err != nil {
		return chassis, err
	}

	url := fmt.Sprintf("https://%s/xmldata?item=all", ip)
	resp, err := client.Get(url)
	if err != nil {
		return chassis, err
	}
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return chassis, err
	}
	defer resp.Body.Close()

	Rimp := &hp.Rimp{}
	err = xml.Unmarshal(payload, Rimp)
	if err != nil {
		return chassis, err
	}

	if Rimp.Infra2 == nil {
		return chassis, errors.ErrUnableToReadData
	}

	return &C7000{ip: ip, username: username, password: password, Rimp: Rimp}, err
}

// CheckCredentials verify whether the credentials are valid or not
func (c *C7000) CheckCredentials() (err error) {
	err = c.httpLogin()
	if err != nil {
		return err
	}
	return err
}

// Name returns the hostname of the machine
func (c *C7000) Name() (name string, err error) {
	return c.Rimp.Infra2.Encl, err
}

// BmcType returns the model id string - c7000
func (c *C7000) BmcType() (model string) {
	return BMCType
}

// Model returns the full device model string
func (c *C7000) Model() (model string, err error) {
	return c.Rimp.MP.Pn, err
}

// Serial returns the device serial
func (c *C7000) Serial() (serial string, err error) {
	return strings.ToLower(strings.TrimSpace(c.Rimp.Infra2.EnclSn)), err
}

// PowerKw returns the current power usage in Kw
func (c *C7000) PowerKw() (power float64, err error) {
	return c.Rimp.Infra2.ChassisPower.PowerConsumed / 1000.00, err
}

// TempC returns the current temperature of the machine
func (c *C7000) TempC() (temp int, err error) {
	return c.Rimp.Infra2.Temp.C, err
}

// Psus returns a list of psus installed on the device
func (c *C7000) Psus() (psus []*devices.Psu, err error) {
	for _, psu := range c.Rimp.Infra2.ChassisPower.Powersupply {
		if psus == nil {
			psus = make([]*devices.Psu, 0)
		}

		p := &devices.Psu{
			Serial:     strings.ToLower(strings.TrimSpace(psu.Sn)),
			Status:     psu.Status,
			PowerKw:    psu.ActualOutput / 1000.00,
			CapacityKw: psu.Capacity / 1000.00,
			PartNumber: psu.Pn,
			Position:   psu.Bay.Connection,
		}
		psus = append(psus, p)
	}

	return psus, err
}

// Nics returns all found Nics in the device
func (c *C7000) Nics() (nics []*devices.Nic, err error) {
	for _, manager := range c.Rimp.Infra2.Managers {
		if nics == nil {
			nics = make([]*devices.Nic, 0)
		}

		n := &devices.Nic{
			Name:       manager.Name,
			MacAddress: strings.ToLower(manager.MacAddr),
		}
		nics = append(nics, n)
	}

	return nics, err
}

// Fans returns all found fans in the device
func (c *C7000) Fans() (fans []*devices.Fan, err error) {
	serial, err := c.Serial()
	if err != nil {
		return fans, err
	}

	for _, fan := range c.Rimp.Infra2.Fans {
		if fans == nil {
			fans = make([]*devices.Fan, 0)
		}

		f := &devices.Fan{
			Serial:     fmt.Sprintf("%d_%s", fan.Bay.Connection, serial),
			Status:     fan.Status,
			Position:   fan.Bay.Connection,
			Model:      fan.ProducName,
			CurrentRPM: fan.RpmCUR,
			PowerKw:    float64(fan.PowerUsed) / 1000,
		}
		fans = append(fans, f)
	}

	return fans, err
}

// Status returns health string status from the bmc
func (c *C7000) Status() (status string, err error) {
	return c.Rimp.Infra2.Status, err
}

// IsActive returns health string status from the bmc
func (c *C7000) IsActive() bool {
	for _, manager := range c.Rimp.Infra2.Managers {
		if manager.MgmtIPAddr == strings.Split(c.ip, ":")[0] && manager.Role == "ACTIVE" {
			return true
		}
	}
	return false
}

// FwVersion returns the current firmware version of the bmc
func (c *C7000) FwVersion() (version string, err error) {
	return c.Rimp.MP.Fwri, err
}

// PassThru returns the type of switch we have for this chassis
func (c *C7000) PassThru() (passthru string, err error) {
	passthru = "1G"
	for _, hpswitch := range c.Rimp.Infra2.Switches {
		if strings.Contains(hpswitch.Spn, "10G") {
			passthru = "10G"
		}
		break
	}
	return passthru, err
}

// StorageBlades returns all StorageBlades found in this chassis
func (c *C7000) StorageBlades() (storageBlades []*devices.StorageBlade, err error) {
	if c.Rimp.Infra2.Blades != nil {
		chassisSerial, _ := c.Serial()
		for _, hpStorageBlade := range c.Rimp.Infra2.Blades {
			if hpStorageBlade.Type == "STORAGE" {
				storageBlade := devices.StorageBlade{}
				storageBlade.Serial = strings.ToLower(strings.TrimSpace(hpStorageBlade.Bsn))
				storageBlade.BladePosition = hpStorageBlade.Bay.Connection
				storageBlade.Status = hpStorageBlade.Status
				storageBlade.PowerKw = hpStorageBlade.Power.PowerConsumed / 1000.00
				storageBlade.TempC = hpStorageBlade.Temp.C
				storageBlade.Vendor = hp.VendorID
				storageBlade.FwVersion = hpStorageBlade.BladeRomVer
				storageBlade.Model = hpStorageBlade.Spn
				storageBlade.ChassisSerial = chassisSerial
				for _, hpBlade := range c.Rimp.Infra2.Blades {
					if hpStorageBlade.AssociatedBlade == hpBlade.Bay.Connection {
						storageBlade.BladeSerial = strings.ToLower(strings.TrimSpace(hpBlade.Bsn))
					}
				}
				storageBlades = append(storageBlades, &storageBlade)
			}
		}
	}
	return storageBlades, err
}

// Blades returns all StorageBlades found in this chassis
func (c *C7000) Blades() (blades []*devices.Blade, err error) {
	if c.Rimp.Infra2.Blades != nil {
		chassisSerial, _ := c.Serial()
		for _, hpBlade := range c.Rimp.Infra2.Blades {
			if hpBlade.Type == "SERVER" {
				blade := devices.Blade{}
				blade.BladePosition = hpBlade.Bay.Connection
				blade.Status = hpBlade.Status
				blade.Serial = strings.ToLower(strings.TrimSpace(hpBlade.Bsn))
				blade.ChassisSerial = chassisSerial
				blade.PowerKw = hpBlade.Power.PowerConsumed / 1000.00
				blade.PowerState = strings.ToLower(hpBlade.Power.PowerState)
				blade.TempC = hpBlade.Temp.C
				blade.Vendor = hp.VendorID
				blade.Model = hpBlade.Spn
				blade.Name = hpBlade.Name
				blade.BmcAddress = hpBlade.MgmtIPAddr
				blade.BmcVersion = hpBlade.MgmtVersion
				blade.BmcType = hpBlade.MgmtType
				blade.BiosVersion = hpBlade.BladeRomVer
				blades = append(blades, &blade)
			}
		}
	}
	return blades, err
}

// Vendor returns bmc's vendor
func (c *C7000) Vendor() (vendor string) {
	return hp.VendorID
}

// ChassisSnapshot do best effort to populate the server data and returns a blade or discrete
func (c *C7000) ChassisSnapshot() (chassis *devices.Chassis, err error) {
	chassis = &devices.Chassis{}
	chassis.Vendor = c.Vendor()
	chassis.BmcAddress = c.ip
	chassis.Name, err = c.Name()
	if err != nil {
		return nil, err
	}
	chassis.Serial, err = c.Serial()
	if err != nil {
		return nil, err
	}
	chassis.Model, err = c.Model()
	if err != nil {
		return nil, err
	}
	chassis.PowerKw, err = c.PowerKw()
	if err != nil {
		return nil, err
	}
	chassis.TempC, err = c.TempC()
	if err != nil {
		return nil, err
	}
	chassis.Status, err = c.Status()
	if err != nil {
		return nil, err
	}
	chassis.FwVersion, err = c.FwVersion()
	if err != nil {
		return nil, err
	}
	chassis.PassThru, err = c.PassThru()
	if err != nil {
		return nil, err
	}
	chassis.Blades, err = c.Blades()
	if err != nil {
		return nil, err
	}
	chassis.StorageBlades, err = c.StorageBlades()
	if err != nil {
		return nil, err
	}
	chassis.Nics, err = c.Nics()
	if err != nil {
		return nil, err
	}
	chassis.Psus, err = c.Psus()
	if err != nil {
		return nil, err
	}
	chassis.Fans, err = c.Fans()
	if err != nil {
		return nil, err
	}

	return chassis, err
}

// UpdateCredentials updates login credentials
func (c *C7000) UpdateCredentials(username string, password string) {
	c.username = username
	c.password = password
}

// GetConfigure returns itself as a configure interface to avoid using reflect
func (c *C7000) GetConfigure() devices.Configure {
	return c
}

// GetSetup returns itself as a configure interface to avoid using reflect
func (c *C7000) GetSetup() devices.CmcSetup {
	return c
}

// GetCollection returns itself as a configure interface to avoid using reflect
func (c *C7000) GetCollection() devices.CmcCollection {
	return c
}
