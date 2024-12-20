package virtbmc

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	kubevirtv1 "kubevirt.io/api/core/v1"

	kubevirtv1type "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type BootDevice string

const (
	Pxe BootDevice = "pxe"
	Hdd BootDevice = "disk"
)

func NewK8sClient(options Options) *kubevirtv1type.KubevirtV1Client {
	var (
		config *rest.Config
		err    error
	)

	// creates the in-cluster config
	config, err = rest.InClusterConfig()
	if err != nil {
		if err == rest.ErrNotInCluster {
			goto localConfig
		}
		panic(err.Error())
	}
	goto genClientset

localConfig:
	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	config, err = clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
	if err != nil {
		panic(err.Error())
	}

genClientset:
	clientset, err := kubevirtv1type.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func (b *VirtBMC) getVirtualMachinePowerStatus() (bool, error) {
	vm, err := b.kvClient.VirtualMachines(b.vmNamespace).Get(b.context, b.vmName, v1.GetOptions{})
	if err != nil {
		return false, err
	}
	if !vm.Status.Ready {
		return false, nil
	}
	return true, nil
}

func (b *VirtBMC) stopVirtualMachine() error {
	vm, err := b.kvClient.VirtualMachines(b.vmNamespace).Get(b.context, b.vmName, v1.GetOptions{})
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
	if _, err := b.kvClient.VirtualMachines(b.vmNamespace).Update(b.context, vm, v1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (b *VirtBMC) startVirtualMachine() error {
	vm, err := b.kvClient.VirtualMachines(b.vmNamespace).Get(b.context, b.vmName, v1.GetOptions{})
	if err != nil {
		return err
	}
	if vm.Spec.RunStrategy == nil {
		running := func(b bool) *bool { return &b }(true)
		vm.Spec.Running = running
	} else {
		runStrategy := func(rs kubevirtv1.VirtualMachineRunStrategy) *kubevirtv1.VirtualMachineRunStrategy {
			return &rs
		}(kubevirtv1.RunStrategyRerunOnFailure)
		vm.Spec.RunStrategy = runStrategy
	}
	if _, err := b.kvClient.VirtualMachines(b.vmNamespace).Update(b.context, vm, v1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (b *VirtBMC) rebootVirtualMachine() error {
	return b.kvClient.VirtualMachineInstances(b.vmNamespace).Delete(b.context, b.vmName, v1.DeleteOptions{})
}

func (b *VirtBMC) setVirtualMachineBootDevice(bd BootDevice) error {
	logrus.Info("setVirtualMachineBootDevice")
	vm, err := b.kvClient.VirtualMachines(b.vmNamespace).Get(b.context, b.vmName, v1.GetOptions{})
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
	switch bd {
	case Pxe:
		vm.Spec.Template.Spec.Domain.Devices.Interfaces[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Interfaces[0])
	case Hdd:
		vm.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Disks[0])
	}

	if _, err := b.kvClient.VirtualMachines(b.vmNamespace).Update(b.context, vm, v1.UpdateOptions{}); err != nil {
		logrus.Errorf("update vm error: %v", err)
		return err
	}

	return nil
}
