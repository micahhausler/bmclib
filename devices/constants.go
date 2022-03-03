package devices

const (
	// Unknown is the constant that defines unknown things
	Unknown = "Unknown"

	// Vendor constants

	// HP is the constant that defines the vendor HP
	HP = "HP"
	// Dell is the constant that defines the vendor Dell
	Dell = "Dell"
	// Supermicro is the constant that defines the vendor Supermicro
	Supermicro = "Supermicro"
	// Cloudline is the constant that defines the cloudlines
	Cloudline = "Cloudline"
	// Quanta is the contant to identify Quanta hardware
	Quanta = "Quanta"
	// Common is the constant of thinks we could use across multiple vendors
	Common = "Common"

	// Power Constants

	// Grid describes the power redundancy mode when using grid redundancy
	Grid = "Grid"

	// PowerSupply describes the power redundancy mode when using power supply redundancy
	PowerSupply = "PowerSupply"

	// NoRedundancy describes the power redundancy mode we don't have redundancy
	NoRedundancy = "NoRedundancy"

	// Hardware constants

	// BladeHwType is the constant defining the blade hw type
	BladeHwType = "blade"
	// DiscreteHwType is the constant defining the Discrete hw type
	DiscreteHwType = "discrete"
	// ChassisHwType is the constant defining the chassis hw type
	ChassisHwType = "chassis"

	// Generic component slugs
	// Slugs are set on Device types to identify the type of component
	SlugBackplaneExpander     = "Backplane Expander"
	SlugChassis               = "Chassis"
	SlugTPM                   = "TPM"
	SlugGPU                   = "GPU"
	SlugCPU                   = "CPU"
	SlugPhysicalMem           = "PhysicalMemory"
	SlugStorageController     = "StorageController"
	SlugStorageControllers    = "StorageControllers"
	SlugBMC                   = "BMC"
	SlugBIOS                  = "BIOS"
	SlugDrive                 = "Drive"
	SlugDrives                = "Drives"
	SlugDriveTypePCIeNVMEeSSD = "NVMe PCIe SSD"
	SlugDriveTypeSATASSD      = "Sata SSD"
	SlugDriveTypeSATAHDD      = "Sata HDD"
	SlugNIC                   = "NIC"
	SlugNICs                  = "NICs"
	SlugPSU                   = "Power Supply"
	SlugPSUs                  = "Power Supplies"
	SlugSASHBA330Controller   = "SAS HBA330 Controller"
	SlugCPLD                  = "CPLD"
	SlugEnclosure             = "ENCLOSURE"
	SlugUnknown               = "unknown"

	// Dell specific component slugs
	SlugDellSystemCPLD                  = "Dell System CPLD"
	SlugDellBossAdapter                 = "Boss Adapter"
	SlugDellIdracServiceModule          = "IDrac Service Module"
	SlugDellBossAdapterDisk0            = "Boss Adapter - Disk 0"
	SlugDellBossAdapterDisk1            = "Boss Adapter - Disk 1"
	SlugDellLifeCycleController         = "Lifecycle Controller"
	SlugDellOSCollector                 = "OS Collector"
	SlugDell64bitUefiDiagnostics        = "Dell 64 bit uEFI diagnostics"
	SlugDellBackplaneExpander           = "Backplane Expander"
	SlugDellNonExpanderStorageBackplane = "Non-Expander Storage Backplane (SEP)"
)

// ListSupportedVendors  returns a list of supported vendors
func ListSupportedVendors() []string {
	return []string{HP, Dell, Supermicro}
}
