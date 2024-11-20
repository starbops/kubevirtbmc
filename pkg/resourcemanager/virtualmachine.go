package resourcemanager

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"

	kubevirttypev1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
)

type KubeVirtClientInterface interface {
	VirtualMachines(namespace string) kubevirttypev1.VirtualMachineInterface
	VirtualMachineInstances(namespace string) kubevirttypev1.VirtualMachineInstanceInterface
}

type VirtualMachineResourceManager struct {
	ctx            context.Context
	kvClient       KubeVirtClientInterface
	virtualMachine *kubevirtv1.VirtualMachine

	computerSystem *server.ComputerSystemV1230ComputerSystem
	manager        *server.ManagerV1192Manager
}

func NewVirtualMachineResourceManager(
	ctx context.Context,
	kvClient KubeVirtClientInterface,
) *VirtualMachineResourceManager {
	return &VirtualMachineResourceManager{
		ctx:      ctx,
		kvClient: kvClient,
	}
}

func initVirtualMachineComputerSystem() *server.ComputerSystemV1230ComputerSystem {
	return &server.ComputerSystemV1230ComputerSystem{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      "/redfish/v1/Systems/1",
		OdataType:    "#ComputerSystem.v1_23_0.ComputerSystem",
		Description:  "Computer System",
		Name:         "Computer System",
		Id:           "1",
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
		PowerState:   Ptr(server.RESOURCEPOWERSTATE_OFF),
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
}

func initVirtualMachineManager() *server.ManagerV1192Manager {
	return &server.ManagerV1192Manager{
		OdataContext: "/redfish/v1/$metadata#Manager.Manager",
		OdataId:      "/redfish/v1/Managers/BMC",
		OdataType:    "#Manager.v1_19_2.Manager",
		Description:  "Manager",
		Name:         "BMC",
		Id:           "BMC",
		UUID:         "00000000-0000-0000-0000-000000000000",
		Model:        Ptr("BMC"),
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
}

func (m *VirtualMachineResourceManager) Initialize(namespace, name string) error {
	vm, err := m.kvClient.VirtualMachines(namespace).Get(m.ctx, name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	m.virtualMachine = vm

	m.computerSystem = initVirtualMachineComputerSystem()
	m.manager = initVirtualMachineManager()

	return nil
}

func (m *VirtualMachineResourceManager) GetComputerSystem() (interface{}, error) {
	vm, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Get(m.ctx, m.virtualMachine.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if vm.Status.Ready {
		m.computerSystem.PowerState = Ptr(server.RESOURCEPOWERSTATE_ON)
	} else {
		m.computerSystem.PowerState = Ptr(server.RESOURCEPOWERSTATE_OFF)
	}

	return m.computerSystem, nil
}

func (m *VirtualMachineResourceManager) GetManager() (interface{}, error) {
	return m.manager, nil
}

func (m *VirtualMachineResourceManager) GetPowerStatus() (bool, error) {
	return true, nil
}

func (m *VirtualMachineResourceManager) PowerOn() error {
	vm, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Get(m.ctx, m.virtualMachine.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if vm.Spec.RunStrategy == nil {
		running := func(b bool) *bool { return &b }(true)
		vm.Spec.Running = running
	} else {
		runStrategy := func(
			rs kubevirtv1.VirtualMachineRunStrategy,
		) *kubevirtv1.VirtualMachineRunStrategy {
			return &rs
		}(kubevirtv1.RunStrategyRerunOnFailure)
		vm.Spec.RunStrategy = runStrategy
	}
	if _, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineResourceManager) PowerOff() error {
	vm, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Get(m.ctx, m.virtualMachine.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if vm.Spec.RunStrategy == nil {
		running := func(b bool) *bool { return &b }(false)
		vm.Spec.Running = running
	} else {
		runStrategy := func(rs kubevirtv1.VirtualMachineRunStrategy) *kubevirtv1.VirtualMachineRunStrategy {
			return &rs
		}(kubevirtv1.RunStrategyHalted)
		vm.Spec.RunStrategy = runStrategy
	}
	if _, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineResourceManager) PowerCycle() error {
	return m.kvClient.VirtualMachineInstances(m.virtualMachine.Namespace).
		Delete(m.ctx, m.virtualMachine.Name, metav1.DeleteOptions{})
}

func (m *VirtualMachineResourceManager) SetBootDevice(bootDevice BootDevice) error {
	logrus.Info("setVirtualMachineBootDevice")
	vm, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Get(m.ctx, m.virtualMachine.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for i, dev := range vm.Spec.Template.Spec.Domain.Devices.Disks {
		logrus.Infof("disk: %+v", dev)
		if dev.BootOrder == nil {
			continue
		}
		newOrder := *dev.BootOrder + 1
		vm.Spec.Template.Spec.Domain.Devices.Disks[i].BootOrder = &newOrder
	}
	for i, intf := range vm.Spec.Template.Spec.Domain.Devices.Interfaces {
		logrus.Infof("interface: %+v", intf)
		if intf.BootOrder == nil {
			continue
		}
		newOrder := *intf.BootOrder + 1
		vm.Spec.Template.Spec.Domain.Devices.Interfaces[i].BootOrder = &newOrder
	}

	var firstOrder uint = 1
	switch bootDevice {
	case BootDevicePXE:
		vm.Spec.Template.Spec.Domain.Devices.Interfaces[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Interfaces[0])
	case BootDeviceDisk, BootDeviceHdd:
		vm.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Disks[0])
	}

	if _, err := m.kvClient.VirtualMachines(m.virtualMachine.Namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		logrus.Errorf("update vm error: %v", err)
		return err
	}

	return nil
}

func Ptr[T any](value T) *T {
	return &value
}
