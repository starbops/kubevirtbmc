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

type ComputerSystemV1220Adapter struct {
	computerSystem *server.ComputerSystemV1220ComputerSystem
}

func (a *ComputerSystemV1220Adapter) GetComputerSystem() *server.ComputerSystemV1220ComputerSystem {
	return a.computerSystem
}

func (a *ComputerSystemV1220Adapter) GetID() string {
	return a.computerSystem.Id
}

func (a *ComputerSystemV1220Adapter) GetPowerState() server.ResourcePowerState {
	return a.computerSystem.PowerState
}

func (a *ComputerSystemV1220Adapter) SetPowerState(powerState server.ResourcePowerState) {
	a.computerSystem.PowerState = powerState
}

func NewComputerSystem(id, name string, powerState server.ResourcePowerState) ComputerSystemInterface {
	generatedComputerSystem := &server.ComputerSystemV1220ComputerSystem{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      fmt.Sprintf("/redfish/v1/Systems/%s", id),
		OdataType:    "#ComputerSystem.v1_22_0.ComputerSystem",
		Description:  "Computer System",
		Name:         name,
		Id:           id,
		UUID:         "00000000-0000-0000-0000-000000000000",
		AssetTag:     Ptr(""),
		IndicatorLED: server.COMPUTERSYSTEMV1220INDICATORLED_UNKNOWN,
		Manufacturer: Ptr("KubeVirt"),
		Model:        Ptr("KubeVirt"),
		PartNumber:   Ptr(""),
		SerialNumber: Ptr("000000000000"),
		SKU:          Ptr(""),
		Status:       server.ResourceStatus{},
		SystemType:   server.COMPUTERSYSTEMV1220SYSTEMTYPE_VIRTUAL,
		Links:        server.ComputerSystemV1220Links{},
		PowerState:   powerState,
		Actions: server.ComputerSystemV1220Actions{
			ComputerSystemReset: server.ComputerSystemV1220Reset{
				Target: "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset",
				Title:  "Reset",
			},
		},
		Boot: server.ComputerSystemV1220Boot{
			BootSourceOverrideEnabled: server.COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEENABLED_DISABLED,
			BootSourceOverrideMode:    server.COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEMODE_LEGACY,
			BootSourceOverrideTarget:  server.COMPUTERSYSTEMBOOTSOURCE_PXE,
		},
		OperatingSystem: "/redfish/v1/Systems/1/OperatingSystem",
		VirtualMedia: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/VirtualMedia",
		},
		HostWatchdogTimer: server.ComputerSystemV1220WatchdogTimer{
			FunctionEnabled: Ptr(false),
		},
		MemorySummary: server.ComputerSystemV1220MemorySummary{
			Status:               server.ResourceStatus{},
			TotalSystemMemoryGiB: Ptr(float32(0)),
		},
		NetworkInterfaces: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/NetworkInterfaces",
		},
		ProcessorSummary: server.ComputerSystemV1220ProcessorSummary{
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

	return &ComputerSystemV1220Adapter{computerSystem: generatedComputerSystem}
}

type ManagerInterface interface {
	GetID() string
}

type ManagerV1190Adapter struct {
	manager *server.ManagerV1190Manager
}

func (a *ManagerV1190Adapter) GetID() string {
	return a.manager.Id
}

func (a *ManagerV1190Adapter) GetManager() *server.ManagerV1190Manager {
	return a.manager
}

func NewManager(id, name string) ManagerInterface {
	generatedManager := &server.ManagerV1190Manager{
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
		Links:        server.ManagerV1190Links{},
		Actions:      server.ManagerV1190Actions{},
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

	return &ManagerV1190Adapter{manager: generatedManager}
}
