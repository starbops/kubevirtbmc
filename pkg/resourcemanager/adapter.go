package resourcemanager

import (
	"fmt"
	"time"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
)

type ComputerSystemInterface interface {
	GetID() string

	GetPowerState() server.ResourcePowerState
	SetPowerState(powerState server.ResourcePowerState)
}

type ComputerSystemV1230Adapter struct {
	computerSystem *server.ComputerSystemV1230ComputerSystem
}

func (a *ComputerSystemV1230Adapter) GetComputerSystem() *server.ComputerSystemV1230ComputerSystem {
	return a.computerSystem
}

func (a *ComputerSystemV1230Adapter) GetID() string {
	return a.computerSystem.Id
}

func (a *ComputerSystemV1230Adapter) GetPowerState() server.ResourcePowerState {
	return *a.computerSystem.PowerState
}

func (a *ComputerSystemV1230Adapter) SetPowerState(ps server.ResourcePowerState) {
	a.computerSystem.PowerState = Ptr(ps)
}

func NewComputerSystem(id, name string, powerState server.ResourcePowerState) ComputerSystemInterface {
	generatedComputerSystem := &server.ComputerSystemV1230ComputerSystem{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      fmt.Sprintf("/redfish/v1/Systems/%s", id),
		OdataType:    "#ComputerSystem.v1_23_0.ComputerSystem",
		Description:  "Computer System",
		Name:         name,
		Id:           id,
		UUID:         "00000000-0000-0000-0000-000000000000",
		AssetTag:     Ptr(""),
		IndicatorLED: Ptr(server.COMPUTERSYSTEMV1230INDICATORLED_UNKNOWN),
		Manufacturer: Ptr("KubeVirt"),
		Model:        Ptr("KubeVirt"),
		PartNumber:   Ptr(""),
		SerialNumber: Ptr("000000000000"),
		SKU:          Ptr(""),
		Status:       server.ResourceStatus{},
		SystemType:   server.COMPUTERSYSTEMV1230SYSTEMTYPE_VIRTUAL,
		Links:        server.ComputerSystemV1230Links{},
		PowerState:   Ptr(powerState),
		Actions: server.ComputerSystemV1230Actions{
			ComputerSystemReset: server.ComputerSystemV1230Reset{
				Target: "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset",
				Title:  "Reset",
			},
		},
		Boot: server.ComputerSystemV1230Boot{
			BootSourceOverrideEnabled: Ptr(server.COMPUTERSYSTEMV1230BOOTSOURCEOVERRIDEENABLED_DISABLED),
			BootSourceOverrideMode:    Ptr(server.COMPUTERSYSTEMV1230BOOTSOURCEOVERRIDEMODE_LEGACY),
			BootSourceOverrideTarget:  Ptr(server.COMPUTERSYSTEMBOOTSOURCE_PXE),
		},
		OperatingSystem: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/OperatingSystem",
		},
		VirtualMedia: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/VirtualMedia",
		},
		HostWatchdogTimer: server.ComputerSystemV1230WatchdogTimer{
			FunctionEnabled: Ptr(false),
		},
		MemorySummary: server.ComputerSystemV1230MemorySummary{
			Status:               server.ResourceStatus{},
			TotalSystemMemoryGiB: Ptr(float32(0)),
		},
		NetworkInterfaces: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/NetworkInterfaces",
		},
		ProcessorSummary: server.ComputerSystemV1230ProcessorSummary{
			Status: server.ResourceStatus{},
			Count:  Ptr(int64(0)),
		},
		SimpleStorage: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/SimpleStorage",
		},
		Storage: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/Storage",
		},
	}

	return &ComputerSystemV1230Adapter{computerSystem: generatedComputerSystem}
}

type ManagerInterface interface {
	GetID() string
}

type ManagerV1192Adapter struct {
	manager *server.ManagerV1192Manager
}

func (a *ManagerV1192Adapter) GetID() string {
	return a.manager.Id
}

func (a *ManagerV1192Adapter) GetManager() *server.ManagerV1192Manager {
	return a.manager
}

func NewManager(id, name string) ManagerInterface {
	generatedManager := &server.ManagerV1192Manager{
		OdataContext: "/redfish/v1/$metadata#Manager.Manager",
		OdataId:      fmt.Sprintf("/redfish/v1/Managers/%s", id),
		OdataType:    "#Manager.v1_19_2.Manager",
		Description:  "Manager",
		Name:         name,
		Id:           id,
		UUID:         "00000000-0000-0000-0000-000000000000",
		Model:        Ptr("KubeVirtBMC"),
		Status:       server.ResourceStatus{},
		ManagerType:  "BMC",
		Links:        server.ManagerV1192Links{},
		Actions:      server.ManagerV1192Actions{},
		DateTime:     Ptr(time.Now()),
		EthernetInterfaces: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Managers/BMC/EthernetInterfaces",
		},
		LogServices: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Managers/BMC/LogServices",
		},
		SerialInterfaces: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Managers/BMC/SerialInterfaces",
		},
		VirtualMedia: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Managers/BMC/VirtualMedia",
		},
	}

	return &ManagerV1192Adapter{manager: generatedManager}
}
