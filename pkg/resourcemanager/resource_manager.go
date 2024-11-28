package resourcemanager

type BootDevice string

const (
	BootDevicePxe BootDevice = "Pxe"
	BootDeviceHdd BootDevice = "Hdd"
)

type ResourceManager interface {
	GetComputerSystem() (ComputerSystemInterface, error)
	GetManager() (ManagerInterface, error)

	GetPowerStatus() (bool, error)
	PowerOn() error
	PowerOff() error
	PowerCycle() error
	SetBootDevice(BootDevice) error
}
