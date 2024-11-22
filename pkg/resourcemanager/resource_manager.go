package resourcemanager

type BootDevice string

const (
	BootDevicePXE  BootDevice = "pxe"
	BootDeviceDisk BootDevice = "disk"
	BootDeviceHdd  BootDevice = "hdd"
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
