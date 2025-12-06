package resourcemanager

import (
	"fmt"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/util"
)

type VirtualMediaInterface interface {
	GetID() string

	GetVirtualMedia() *server.VirtualMediaV163VirtualMedia
}

type VirtualMediaAdapter struct {
	virtualMedia *server.VirtualMediaV163VirtualMedia
}

func NewVirtualMedia(id, name string) *VirtualMediaAdapter {
	generatedVirtualMedia := &server.VirtualMediaV163VirtualMedia{
		OdataContext: "/redfish/v1/$metadata#VirtualMedia.VirtualMedia",
		OdataId:      fmt.Sprintf("/redfish/v1/Managers/BMC/VirtualMedia/%s", id),
		OdataType:    "#VirtualMedia.v1_6_3.VirtualMedia",
		ConnectedVia: server.VIRTUALMEDIAV163CONNECTEDVIA_NOT_CONNECTED,
		Description:  "Virtual Media",
		Name:         name,
		Id:           id,
		Image:        util.Ptr(""),
		ImageName:    util.Ptr(""),
		Inserted:     util.Ptr(false),
		MediaTypes: []server.VirtualMediaV163MediaType{
			"CD",
			"DVD",
		},
		WriteProtected: util.Ptr(false),
	}

	return &VirtualMediaAdapter{virtualMedia: generatedVirtualMedia}
}

func (a *VirtualMediaAdapter) GetID() string {
	return a.virtualMedia.Id
}

func (a *VirtualMediaAdapter) GetODataID() string {
	return a.virtualMedia.OdataId
}

func (a *VirtualMediaAdapter) Manage(resource ODataInterface) error {
	panic("implement me")
}

func (a *VirtualMediaAdapter) ManagedBy(resource ODataInterface) error {
	panic("implement me")
}

func (a *VirtualMediaAdapter) GetVirtualMedia() *server.VirtualMediaV163VirtualMedia {
	return a.virtualMedia
}

func (a *VirtualMediaAdapter) SetVirtualMedia(imageURL string, inserted bool) {
	a.virtualMedia.Image = util.Ptr(imageURL)
	a.virtualMedia.Inserted = util.Ptr(inserted)
	if inserted {
		a.virtualMedia.ConnectedVia = server.VIRTUALMEDIAV163CONNECTEDVIA_URI
	} else {
		a.virtualMedia.ConnectedVia = server.VIRTUALMEDIAV163CONNECTEDVIA_NOT_CONNECTED
	}
}
