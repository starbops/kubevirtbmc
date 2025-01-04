package fake

import (
	"context"

	"github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

type MockVirtualMachineInstanceInterface struct {
	mock.Mock
}

func (m *MockVirtualMachineInstanceInterface) Get(
	ctx context.Context, name string, options metav1.GetOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Create(
	ctx context.Context, vmi *kubevirtv1.VirtualMachineInstance, options metav1.CreateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Update(
	ctx context.Context, vmi *kubevirtv1.VirtualMachineInstance, options metav1.UpdateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	args := m.Called(ctx, vmi, options)
	return args.Get(0).(*kubevirtv1.VirtualMachineInstance), args.Error(1)
}

func (m *MockVirtualMachineInstanceInterface) Delete(
	ctx context.Context, name string, options metav1.DeleteOptions,
) error {
	args := m.Called(ctx, name, options)
	return args.Error(0)
}

func (m *MockVirtualMachineInstanceInterface) DeleteCollection(
	ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions,
) error {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) List(
	ctx context.Context, options metav1.ListOptions,
) (*kubevirtv1.VirtualMachineInstanceList, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	options metav1.PatchOptions,
	subresources ...string,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) UpdateStatus(
	ctx context.Context, vm *kubevirtv1.VirtualMachineInstance, options metav1.UpdateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Watch(
	ctx context.Context, options metav1.ListOptions,
) (watch.Interface, error) {
	panic("implement me")
}
