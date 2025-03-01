package resourcemanager

import (
	"fmt"
	"time"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/util"
)

type ManagerInterface interface {
	GetID() string
}

type ManagerAdapter struct {
	manager *server.ManagerV1190Manager
}

func (a *ManagerAdapter) GetID() string {
	return a.manager.Id
}

func (a *ManagerAdapter) GetManager() *server.ManagerV1190Manager {
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
		Model:        util.Ptr("KubeVirtBMC"),
		Status:       server.ResourceStatus{},
		ManagerType:  "BMC",
		Links:        server.ManagerV1190Links{},
		Actions:      server.ManagerV1190Actions{},
		DateTime:     util.Ptr(time.Now()),
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

	return &ManagerAdapter{manager: generatedManager}
}
