package kbmc

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	kubevirtv1 "kubevirt.io/api/core/v1"

	kubevirtv1type "zespre.com/kubebmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type BootDevice string

const (
	Pxe  BootDevice = "pxe"
	Disk BootDevice = "disk"
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

func (k *KBMC) getVirtualMachine(namespace, name string) (*kubevirtv1.VirtualMachine, error) {
	vm, err := k.kvClient.VirtualMachines(namespace).Get(k.context, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func (k *KBMC) getVirtualMachineInstance(namespace, name string) (*kubevirtv1.VirtualMachineInstance, error) {
	vmi, err := k.kvClient.VirtualMachineInstances(namespace).Get(k.context, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return vmi, nil
}

func (k *KBMC) getVirtualMachinePowerStatus() (bool, error) {
	vm, err := k.kvClient.VirtualMachines(k.vmNamespace).Get(k.context, k.vmName, v1.GetOptions{})
	if err != nil {
		return false, err
	}
	if !vm.Status.Ready {
		return false, nil
	}
	return true, nil
}

func (k *KBMC) stopVirtualMachine() error {
	vm, err := k.kvClient.VirtualMachines(k.vmNamespace).Get(k.context, k.vmName, v1.GetOptions{})
	if err != nil {
		return err
	}
	runStrategy := kubevirtv1.RunStrategyHalted
	vm.Spec.RunStrategy = &runStrategy
	if _, err := k.kvClient.VirtualMachines(k.vmNamespace).Update(k.context, vm, v1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (k *KBMC) startVirtualMachine() error {
	vm, err := k.kvClient.VirtualMachines(k.vmNamespace).Get(k.context, k.vmName, v1.GetOptions{})
	if err != nil {
		return err
	}
	runStrategy := kubevirtv1.RunStrategyRerunOnFailure
	vm.Spec.RunStrategy = &runStrategy
	if _, err := k.kvClient.VirtualMachines(k.vmNamespace).Update(k.context, vm, v1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (k *KBMC) rebootVirtualMachine() error {
	if err := k.kvClient.VirtualMachineInstances(k.vmNamespace).Delete(k.context, k.vmName, v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (k *KBMC) setVirtualMachineBootDevice(bd BootDevice) error {
	logrus.Info("setVirtualMachineBootDevice")
	vm, err := k.kvClient.VirtualMachines(k.vmNamespace).Get(k.context, k.vmName, v1.GetOptions{})
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
	case Disk:
		vm.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Disks[0])
	}

	if _, err := k.kvClient.VirtualMachines(k.vmNamespace).Update(k.context, vm, v1.UpdateOptions{}); err != nil {
		logrus.Errorf("update vm error: %v", err)
		return err
	}

	return nil
}
