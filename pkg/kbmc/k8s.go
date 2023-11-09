package kbmc

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	kubevirtv1 "kubevirt.io/api/core/v1"

	kubevirtv1type "zespre.com/kubebmc/pkg/generated/clientset/versioned/typed/core/v1"
)

func NewK8sClient(options Options) *kubevirtv1type.KubevirtV1Client {
	var (
		config *rest.Config
		err    error
	)

	// if options.KubeconfigPath == "" {
	// creates the in-cluster config
	config, err = rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// } else {
	// 	// uses the current context in kubeconfig
	// 	// path-to-kubeconfig -- for example, /root/.kube/config
	// 	config, err = clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

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
