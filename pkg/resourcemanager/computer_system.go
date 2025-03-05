package resourcemanager

import (
	"fmt"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/util"
)

type ComputerSystemInterface interface {
	GetID() string

	GetComputerSystem() *server.ComputerSystemV1220ComputerSystem
	GetPowerState() server.ResourcePowerState
	SetPowerState(powerState server.ResourcePowerState)
	SetBootOverride(server.ComputerSystemBootSource)
}

type ComputerSystemAdapter struct {
	computerSystem *server.ComputerSystemV1220ComputerSystem
}

func (a *ComputerSystemAdapter) GetODataID() string {
	return a.computerSystem.OdataId
}

func (a *ComputerSystemAdapter) Manage(resource ODataInterface) error {
	panic("implement me")
}

func (a *ComputerSystemAdapter) ManagedBy(resource ODataInterface) error {
	a.computerSystem.Links.ManagedBy = append(a.computerSystem.Links.ManagedBy, server.OdataV4IdRef{
		OdataId: resource.GetODataID(),
	})

	return nil
}

func (a *ComputerSystemAdapter) GetComputerSystem() *server.ComputerSystemV1220ComputerSystem {
	return a.computerSystem
}

func (a *ComputerSystemAdapter) GetID() string {
	return a.computerSystem.Id
}

func (a *ComputerSystemAdapter) GetPowerState() server.ResourcePowerState {
	return a.computerSystem.PowerState
}

func (a *ComputerSystemAdapter) SetPowerState(powerState server.ResourcePowerState) {
	a.computerSystem.PowerState = powerState
}

func (a *ComputerSystemAdapter) SetBootOverride(target server.ComputerSystemBootSource) {
	a.computerSystem.Boot.BootSourceOverrideEnabled = server.COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEENABLED_CONTINUOUS
	a.computerSystem.Boot.BootSourceOverrideTarget = target
}

func NewComputerSystem(id, name string, powerState server.ResourcePowerState) *ComputerSystemAdapter {
	generatedComputerSystem := &server.ComputerSystemV1220ComputerSystem{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      fmt.Sprintf("/redfish/v1/Systems/%s", id),
		OdataType:    "#ComputerSystem.v1_22_0.ComputerSystem",
		Description:  "Computer System",
		Name:         name,
		Id:           id,
		UUID:         "00000000-0000-0000-0000-000000000000",
		AssetTag:     util.Ptr(""),
		IndicatorLED: server.COMPUTERSYSTEMV1220INDICATORLED_UNKNOWN,
		Manufacturer: util.Ptr("KubeVirt"),
		Model:        util.Ptr("KubeVirt"),
		PartNumber:   util.Ptr(""),
		SerialNumber: util.Ptr("000000000000"),
		SKU:          util.Ptr(""),
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
			BootSourceOverrideTarget:  server.COMPUTERSYSTEMBOOTSOURCE_HDD,
		},
		OperatingSystem: "/redfish/v1/Systems/1/OperatingSystem",
		VirtualMedia: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/VirtualMedia",
		},
		HostWatchdogTimer: server.ComputerSystemV1220WatchdogTimer{
			FunctionEnabled: util.Ptr(false),
		},
		MemorySummary: server.ComputerSystemV1220MemorySummary{
			Status:               server.ResourceStatus{},
			TotalSystemMemoryGiB: util.Ptr(float32(0)),
		},
		NetworkInterfaces: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/NetworkInterfaces",
		},
		ProcessorSummary: server.ComputerSystemV1220ProcessorSummary{
			Status: server.ResourceStatus{},
			Count:  util.Ptr(int64(0)),
		},
		SimpleStorage: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/SimpleStorage",
		},
		Storage: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/1/Storage",
		},
	}

	return &ComputerSystemAdapter{computerSystem: generatedComputerSystem}
}
