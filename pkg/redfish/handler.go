package redfish

import (
	"fmt"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
)

type ResourceHandler struct {
	Manager resourcemanager.ResourceManager
}

func NewResourceHandler(resourceManager resourcemanager.ResourceManager) *ResourceHandler {
	return &ResourceHandler{
		Manager: resourceManager,
	}
}

func (h *ResourceHandler) GetServiceRoot() *server.ServiceRootV1170ServiceRoot {
	return &server.ServiceRootV1170ServiceRoot{
		OdataContext:   "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
		OdataId:        "/redfish/v1",
		OdataType:      "#ServiceRoot.v1_17_0.ServiceRoot",
		Description:    "ServiceRoot",
		Name:           "ServiceRoot",
		RedfishVersion: "1.17.0",
		UUID:           Ptr("00000000-0000-0000-0000-000000000000"),
		Chassis: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis",
		},
		Managers: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Managers",
		},
		Registries: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Registries",
		},
		SessionService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/SessionService",
		},
		Systems: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems",
		},
		Tasks: server.OdataV4IdRef{
			OdataId: "/redfish/v1/Tasks",
		},
		AccountService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/AccountService",
		},
		EventService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/EventService",
		},
		TelemetryService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/TelemetryService",
		},
		UpdateService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/UpdateService",
		},
		CompositionService: server.OdataV4IdRef{
			OdataId: "/redfish/v1/CompositionService",
		},
		ProtocolFeaturesSupported: server.ServiceRootV1170ProtocolFeaturesSupported{},
		Links: server.ServiceRootV1170Links{
			ManagerProvidingService: server.OdataV4IdRef{
				OdataId: "/redfish/v1/Managers/BMC",
			},
			Oem: map[string]interface{}{},
			Sessions: server.OdataV4IdRef{
				OdataId: "/redfish/v1/SessionService/Sessions",
			},
		},
	}
}

func (h *ResourceHandler) GetMetadata() *server.OdataV4IdRef {
	return &server.OdataV4IdRef{
		OdataId: "/redfish/v1/$metadata",
	}
}

func (h *ResourceHandler) GetManagerCollection() *server.ManagerCollectionManagerCollection {
	return &server.ManagerCollectionManagerCollection{
		OdataContext: "/redfish/v1/$metadata#ManagerCollection.ManagerCollection",
		OdataId:      "/redfish/v1/Managers",
		OdataType:    "#ManagerCollection.ManagerCollection",
		Description:  "Manager Collection",
		Name:         "Manager Collection",
		Members: []server.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Managers/BMC",
			},
		},
	}
}

func (h *ResourceHandler) GetManager() (*server.ManagerV1192Manager, error) {
	manager, err := h.Manager.GetManager()
	if err != nil {
		return nil, err
	}
	adapter, ok := manager.(*resourcemanager.ManagerV1192Adapter)
	if !ok {
		return nil, fmt.Errorf("unexpected manager type: %T", manager)
	}
	return adapter.GetManager(), nil
}

func (h *ResourceHandler) GetVirtualMediaCollection() *server.VirtualMediaCollectionVirtualMediaCollection {
	return &server.VirtualMediaCollectionVirtualMediaCollection{
		OdataContext: "/redfish/v1/$metadata#VirtualMediaCollection.VirtualMediaCollection",
		OdataId:      "/redfish/v1/Managers/BMC/VirtualMedia",
		OdataType:    "#VirtualMediaCollection.VirtualMediaCollection",
		Description:  "Virtual Media Collection",
		Name:         "Virtual Media Collection",
		Members: []server.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Managers/BMC/VirtualMedia/1",
			},
		},
	}
}

func (h *ResourceHandler) GetVirtualMedia() *server.VirtualMediaV164VirtualMedia {
	return &server.VirtualMediaV164VirtualMedia{
		OdataContext: "/redfish/v1/$metadata#VirtualMedia.VirtualMedia",
		OdataId:      "/redfish/v1/Managers/BMC/VirtualMedia/1",
		OdataType:    "#VirtualMedia.v1_6_4.VirtualMedia",
		Description:  "Virtual Media",
		Name:         "Virtual Media",
		Id:           "1",
		Image:        Ptr(""),
		ImageName:    Ptr(""),
		Inserted:     Ptr(false),
		MediaTypes: []server.VirtualMediaV164MediaType{
			"CD",
			"DVD",
			"USBStick",
			"Floppy",
			"ISO",
			"OEM",
		},
		WriteProtected: Ptr(false),
	}
}

func (h *ResourceHandler) GetComputerSystemCollection() *server.ComputerSystemCollectionComputerSystemCollection {
	return &server.ComputerSystemCollectionComputerSystemCollection{
		OdataContext: "/redfish/v1/$metadata#ComputerSystemCollection.ComputerSystemCollection",
		OdataId:      "/redfish/v1/Systems",
		OdataType:    "#ComputerSystemCollection.ComputerSystemCollection",
		Description:  "Computer System Collection",
		Name:         "Computer System Collection",
		Members: []server.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Systems/1",
			},
		},
	}
}

func (h *ResourceHandler) GetComputerSystem() (*server.ComputerSystemV1230ComputerSystem, error) {
	computerSystem, err := h.Manager.GetComputerSystem()
	if err != nil {
		return nil, err
	}
	adapter, ok := computerSystem.(*resourcemanager.ComputerSystemV1230Adapter)
	if !ok {
		return nil, fmt.Errorf("unexpected computer system type: %T", computerSystem)
	}
	return adapter.GetComputerSystem(), nil
}

func (h *ResourceHandler) ComputerSystemReset(resetType server.ResourceResetType) error {
	powerActionMap := map[server.ResourceResetType]func() error{
		server.RESOURCERESETTYPE_ON:                h.Manager.PowerOn,
		server.RESOURCERESETTYPE_GRACEFUL_SHUTDOWN: h.Manager.PowerOff,
		server.RESOURCERESETTYPE_FORCE_OFF:         h.Manager.PowerOff,
		server.RESOURCERESETTYPE_GRACEFUL_RESTART:  h.Manager.PowerCycle,
		server.RESOURCERESETTYPE_FORCE_RESTART:     h.Manager.PowerCycle,
	}

	powerAction, ok := powerActionMap[resetType]
	if !ok {
		return fmt.Errorf("unsupported reset type: %s", resetType)
	}
	return powerAction()
}

// ComputerSystemSetDefaultBootOrder sets the boot order for the computer system back to default.
// TODO: Implement real default boot order setting. Right now we intentionally misuse the handler to set the first boot
// device.
func (h *ResourceHandler) ComputerSystemSetDefaultBootOrder(bootDevices []string) error {
	var bootDevice resourcemanager.BootDevice
	if len(bootDevices) > 0 {
		bootDevice = resourcemanager.BootDevice(bootDevices[0])
	}
	return h.Manager.SetBootDevice(bootDevice)
}

func Ptr[T any](value T) *T {
	return &value
}