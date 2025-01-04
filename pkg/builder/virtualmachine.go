package builder

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

// VirtualMachineBuilder builds a VirtualMachine object.
type VirtualMachineBuilder struct {
	vm *kubevirtv1.VirtualMachine
}

func NewVirtualMachineBuilder(namespace, name string) *VirtualMachineBuilder {
	return &VirtualMachineBuilder{
		vm: &kubevirtv1.VirtualMachine{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      name,
			},
		},
	}
}

func (b *VirtualMachineBuilder) Build() *kubevirtv1.VirtualMachine {
	return b.vm
}

func (b *VirtualMachineBuilder) Running(running bool) *VirtualMachineBuilder {
	b.vm.Spec.Running = &running
	return b
}

func (b *VirtualMachineBuilder) RunStrategy(strategy kubevirtv1.VirtualMachineRunStrategy) *VirtualMachineBuilder {
	b.vm.Spec.RunStrategy = &strategy
	return b
}

func (b *VirtualMachineBuilder) Ready(ready bool) *VirtualMachineBuilder {
	b.vm.Status.Ready = ready
	return b
}
