package redfish

import (
	"context"
	"errors"
	"net/http"

	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
)

type BootOrderRequest struct {
	BootOrder []string `json:"BootOrder"`
}

type APIService struct {
	handler *ResourceHandler
}

func NewAPIService(resourceManager resourcemanager.ResourceManager) *APIService {
	return &APIService{
		handler: NewResourceHandler(resourceManager),
	}
}

func (s *APIService) RedfishV1Get(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(200, s.handler.GetServiceRoot()), nil
}

func (s *APIService) RedfishV1MetadataGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1MetadataGet method not implemented")
}

func (s *APIService) RedfishV1ManagersGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(200, s.handler.GetManagerCollection()), nil
}

func (s *APIService) RedfishV1ManagersManagerIdGet(ctx context.Context, managerId string) (server.ImplResponse, error) {
	manager, err := s.handler.GetManager()
	if err != nil {
		return server.Response(http.StatusInternalServerError, nil), err
	}

	return server.Response(200, manager), nil
}

func (s *APIService) RedfishV1ManagersManagerIdPut(ctx context.Context, managerId string, managerV1192Manager server.ManagerV1192Manager) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdPut method not implemented")
}

func (s *APIService) RedfishV1ManagersManagerIdPatch(ctx context.Context, managerId string, managerV1192Manager server.ManagerV1192Manager) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdPatch method not implemented")
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaGet(ctx context.Context, managerId string) (server.ImplResponse, error) {
	return server.Response(200, s.handler.GetVirtualMediaCollection()), nil
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdGet(ctx context.Context, managerId string, virtualMediaId string) (server.ImplResponse, error) {
	return server.Response(200, s.handler.GetVirtualMedia()), nil
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdPut(ctx context.Context, managerId string, virtualMediaId string, virtualMediaV164VirtualMedia server.VirtualMediaV164VirtualMedia) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdPut method not implemented")
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdPatch(ctx context.Context, managerId string, virtualMediaId string, virtualMediaV164VirtualMedia server.VirtualMediaV164VirtualMedia) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdPatch method not implemented")
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdActionsVirtualMediaEjectMediaPost(ctx context.Context, managerId string, virtualMediaId string, body map[string]interface{}) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdActionsVirtualMediaEjectMediaPost method not implemented")
}

func (s *APIService) RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdActionsVirtualMediaInsertMediaPost(ctx context.Context, managerId string, virtualMediaId string, virtualMediaV164InsertMediaRequestBody server.VirtualMediaV164InsertMediaRequestBody) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1ManagersManagerIdVirtualMediaVirtualMediaIdActionsVirtualMediaInsertMediaPost method not implemented")
}

func (s *APIService) RedfishV1SessionServiceGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServiceGet method not implemented")
}

func (s *APIService) RedfishV1SessionServicePut(ctx context.Context, sessionServiceV119SessionService server.SessionServiceV119SessionService) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServicePut method not implemented")
}

func (s *APIService) RedfishV1SessionServicePatch(ctx context.Context, sessionServiceV119SessionService server.SessionServiceV119SessionService) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServicePatch method not implemented")
}

func (s *APIService) RedfishV1SessionServiceSessionsGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServiceSessionsGet method not implemented")
}

func (s *APIService) RedfishV1SessionServiceSessionsPost(ctx context.Context, sessionV172Session server.SessionV172Session) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServiceSessionsPost method not implemented")
}

func (s *APIService) RedfishV1SessionServiceSessionsSessionIdGet(ctx context.Context, sessionId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServiceSessionsSessionIdGet method not implemented")
}

func (s *APIService) RedfishV1SessionServiceSessionsSessionIdDelete(ctx context.Context, sessionId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServiceSessionsSessionIdDelete method not implemented")
}

func (s *APIService) RedfishV1SystemsGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(200, s.handler.GetComputerSystemCollection()), nil
}

func (s *APIService) RedfishV1SystemsPost(ctx context.Context, computerSystemV1230ComputerSystem server.ComputerSystemV1230ComputerSystem) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	computerSystem, err := s.handler.GetComputerSystem()
	if err != nil {
		return server.Response(http.StatusInternalServerError, nil), err
	}

	return server.Response(200, computerSystem), nil
}

func (s *APIService) RedfishV1SystemsComputerSystemIdPut(ctx context.Context, computerSystemId string, computerSystemV1230ComputerSystem server.ComputerSystemV1230ComputerSystem) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdPut method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdDelete(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdDelete method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdPatch(ctx context.Context, computerSystemId string, computerSystemV1230ComputerSystem server.ComputerSystemV1230ComputerSystem) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdPatch method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdActionsComputerSystemResetPost(ctx context.Context, computerSystemId string, computerSystemV1230ResetRequestBody server.ComputerSystemV1230ResetRequestBody) (server.ImplResponse, error) {
	if !computerSystemV1230ResetRequestBody.ResetType.IsValid() {
		return server.Response(http.StatusBadRequest, server.RedfishError{
			Error: server.RedfishErrorError{
				MessageExtendedInfo: []server.MessageV121Message{
					{
						MessageId: "Base.1.2.PropertyValueNotInList",
						Message:   "The value " + string(computerSystemV1230ResetRequestBody.ResetType) + " for the property ResetType is not in the list of acceptable values.",
					},
				},
			},
		}), nil
	}

	if err := s.handler.ComputerSystemReset(computerSystemV1230ResetRequestBody.ResetType); err != nil {
		return server.Response(http.StatusInternalServerError, server.RedfishError{
			Error: server.RedfishErrorError{
				MessageExtendedInfo: []server.MessageV121Message{
					{
						MessageId: "Base.1.2.GeneralError",
						Message:   err.Error(),
					},
				},
			},
		}), nil
	}

	return server.Response(204, nil), nil
}

func (s *APIService) RedfishV1SystemsComputerSystemIdActionsComputerSystemSetDefaultBootOrderPost(ctx context.Context, computerSystemId string, body map[string]interface{}) (server.ImplResponse, error) {
	rawBootOrder, ok := body["BootOrder"].([]interface{})
	if !ok {
		return server.Response(http.StatusBadRequest, server.RedfishError{
			Error: server.RedfishErrorError{
				MessageExtendedInfo: []server.MessageV121Message{
					{
						MessageId: "Base.1.2.PropertyValueTypeError",
						Message:   "The property BootOrder must be an array of strings.",
					},
				},
			},
		}), nil
	}

	bootOrder := make([]string, len(rawBootOrder))
	for i, v := range rawBootOrder {
		str, ok := v.(string)
		if !ok {
			return server.Response(http.StatusBadRequest, server.RedfishError{
				Error: server.RedfishErrorError{
					MessageExtendedInfo: []server.MessageV121Message{
						{
							MessageId: "Base.1.2.PropertyValueTypeError",
							Message:   "The property BootOrder must only contain strings.",
						},
					},
				},
			}), nil
		}
		bootOrder[i] = str
	}

	if err := s.handler.ComputerSystemSetDefaultBootOrder(bootOrder); err != nil {
		return server.Response(http.StatusInternalServerError, server.RedfishError{
			Error: server.RedfishErrorError{
				MessageExtendedInfo: []server.MessageV121Message{
					{
						MessageId: "Base.1.2.GeneralError",
						Message:   err.Error(),
					},
				},
			},
		}), nil
	}

	return server.Response(204, nil), nil
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsPost(ctx context.Context, computerSystemId string, bootOptionV106BootOption server.BootOptionV106BootOption) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdGet(ctx context.Context, computerSystemId string, bootOptionId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdPut(ctx context.Context, computerSystemId string, bootOptionId string, bootOptionV106BootOption server.BootOptionV106BootOption) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdPut method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdDelete(ctx context.Context, computerSystemId string, bootOptionId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdDelete method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdPatch(ctx context.Context, computerSystemId string, bootOptionId string, bootOptionV106BootOption server.BootOptionV106BootOption) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdBootOptionsBootOptionIdPatch method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	operatingSystem := server.OperatingSystemV102OperatingSystem{
		OdataContext: "/redfish/v1/$metadata#OperatingSystem.OperatingSystem",
		OdataId:      "/redfish/v1/Systems/1/OperatingSystem",
		OdataType:    "#OperatingSystem.v1_0_2.OperatingSystem",
		Name:         "Operating System",
		Id:           "1",
	}

	return server.Response(200, operatingSystem), nil
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsApplicationIdGet(ctx context.Context, computerSystemId string, applicationId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsApplicationIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsApplicationIdActionsApplicationResetPost(ctx context.Context, computerSystemId string, applicationId string, applicationV101ResetRequestBody server.ApplicationV101ResetRequestBody) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemApplicationsApplicationIdActionsApplicationResetPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainerImagesGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainerImagesGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainerImagesContainerImageIdGet(ctx context.Context, computerSystemId string, containerImageId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainerImagesContainerImageIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesPost(ctx context.Context, computerSystemId string, ethernetInterfaceV1122EthernetInterface server.EthernetInterfaceV1122EthernetInterface) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdGet(ctx context.Context, computerSystemId string, ethernetInterfaceId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdPut(ctx context.Context, computerSystemId string, ethernetInterfaceId string, ethernetInterfaceV1122EthernetInterface server.EthernetInterfaceV1122EthernetInterface) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdPut method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdDelete(ctx context.Context, computerSystemId string, ethernetInterfaceId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdDelete method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdPatch(ctx context.Context, computerSystemId string, ethernetInterfaceId string, ethernetInterfaceV1122EthernetInterface server.EthernetInterfaceV1122EthernetInterface) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersEthernetInterfacesEthernetInterfaceIdPatch method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersContainerIdGet(ctx context.Context, computerSystemId string, containerId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersContainerIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdOperatingSystemContainersContainerIdActionsContainerResetPost(ctx context.Context, computerSystemId string, containerId string, containerV101ResetRequestBody server.ContainerV101ResetRequestBody) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdOperatingSystemContainersContainerIdActionsContainerResetPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaGet(ctx context.Context, computerSystemId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdGet(ctx context.Context, computerSystemId string, virtualMediaId string) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdGet method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdPut(ctx context.Context, computerSystemId string, virtualMediaId string, virtualMediaV164VirtualMedia server.VirtualMediaV164VirtualMedia) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdPut method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdPatch(ctx context.Context, computerSystemId string, virtualMediaId string, virtualMediaV164VirtualMedia server.VirtualMediaV164VirtualMedia) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdPatch method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdActionsVirtualMediaEjectMediaPost(ctx context.Context, computerSystemId string, virtualMediaId string, body map[string]interface{}) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdActionsVirtualMediaEjectMediaPost method not implemented")
}

func (s *APIService) RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdActionsVirtualMediaInsertMediaPost(ctx context.Context, computerSystemId string, virtualMediaId string, virtualMediaV164InsertMediaRequestBody server.VirtualMediaV164InsertMediaRequestBody) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdVirtualMediaVirtualMediaIdActionsVirtualMediaInsertMediaPost method not implemented")
}

func (s *APIService) RedfishV1OdataGet(ctx context.Context) (server.ImplResponse, error) {
	return server.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1OdataGet method not implemented")
}
