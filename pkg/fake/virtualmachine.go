package fake

import (
	"context"

	"github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

type MockVirtualMachineInterface struct {
	mock.Mock
}

func (m *MockVirtualMachineInterface) Get(
	ctx context.Context, name string, options metav1.GetOptions,
) (*kubevirtv1.VirtualMachine, error) {
	args := m.Called(ctx, name, options)
	return args.Get(0).(*kubevirtv1.VirtualMachine), args.Error(1)
}

func (m *MockVirtualMachineInterface) Create(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options metav1.CreateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Update(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options metav1.UpdateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	args := m.Called(ctx, vm, options)
	return args.Get(0).(*kubevirtv1.VirtualMachine), args.Error(1)
}

func (m *MockVirtualMachineInterface) Delete(ctx context.Context, name string, options metav1.DeleteOptions) error {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) DeleteCollection(
	ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions,
) error {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) List(
	ctx context.Context, options metav1.ListOptions,
) (*kubevirtv1.VirtualMachineList, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	options metav1.PatchOptions,
	subresources ...string,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) UpdateStatus(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options metav1.UpdateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Watch(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}
