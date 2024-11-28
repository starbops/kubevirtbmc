package virtbmc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	kubevirtv1 "kubevirt.io/api/core/v1"
	kubevirttypev1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type MockKubevirtClient struct {
	mock.Mock
}

func (m *MockKubevirtClient) VirtualMachines(namespace string) kubevirttypev1.VirtualMachineInterface {
	args := m.Called(namespace)
	return args.Get(0).(kubevirttypev1.VirtualMachineInterface)
}

type MockVirtualMachineInterface struct {
	mock.Mock
}

func (m *MockVirtualMachineInterface) Get(
	ctx context.Context, name string, options v1.GetOptions,
) (*kubevirtv1.VirtualMachine, error) {
	args := m.Called(ctx, name, options)
	return args.Get(0).(*kubevirtv1.VirtualMachine), args.Error(1)
}

func (m *MockVirtualMachineInterface) Create(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options v1.CreateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Update(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options v1.UpdateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	args := m.Called(ctx, vm, options)
	return args.Get(0).(*kubevirtv1.VirtualMachine), args.Error(1)
}

func (m *MockVirtualMachineInterface) Delete(ctx context.Context, name string, options v1.DeleteOptions) error {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) DeleteCollection(
	ctx context.Context, options v1.DeleteOptions, listOptions v1.ListOptions,
) error {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) List(
	ctx context.Context, options v1.ListOptions,
) (*kubevirtv1.VirtualMachineList, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	options v1.PatchOptions,
	subresources ...string,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) UpdateStatus(
	ctx context.Context, vm *kubevirtv1.VirtualMachine, options v1.UpdateOptions,
) (*kubevirtv1.VirtualMachine, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInterface) Watch(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (m *MockKubevirtClient) VirtualMachineInstances(namespace string) kubevirttypev1.VirtualMachineInstanceInterface {
	args := m.Called(namespace)
	return args.Get(0).(kubevirttypev1.VirtualMachineInstanceInterface)
}

type MockVirtualMachineInstanceInterface struct {
	mock.Mock
}

func (m *MockVirtualMachineInstanceInterface) Get(
	ctx context.Context, name string, options v1.GetOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Create(
	ctx context.Context, vmi *kubevirtv1.VirtualMachineInstance, options v1.CreateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Update(
	ctx context.Context, vmi *kubevirtv1.VirtualMachineInstance, options v1.UpdateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	args := m.Called(ctx, vmi, options)
	return args.Get(0).(*kubevirtv1.VirtualMachineInstance), args.Error(1)
}

func (m *MockVirtualMachineInstanceInterface) Delete(ctx context.Context, name string, options v1.DeleteOptions) error {
	args := m.Called(ctx, name, options)
	return args.Error(0)
}

func (m *MockVirtualMachineInstanceInterface) DeleteCollection(
	ctx context.Context, options v1.DeleteOptions, listOptions v1.ListOptions,
) error {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) List(
	ctx context.Context, options v1.ListOptions,
) (*kubevirtv1.VirtualMachineInstanceList, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	options v1.PatchOptions,
	subresources ...string,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) UpdateStatus(
	ctx context.Context, vm *kubevirtv1.VirtualMachineInstance, options v1.UpdateOptions,
) (*kubevirtv1.VirtualMachineInstance, error) {
	panic("implement me")
}

func (m *MockVirtualMachineInstanceInterface) Watch(
	ctx context.Context, options v1.ListOptions,
) (watch.Interface, error) {
	panic("implement me")
}

func TestGetVirtualMachinePowerStatus(t *testing.T) {
	mockClient := new(MockKubevirtClient)
	mockVMInterface := new(MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	mockVM := &kubevirtv1.VirtualMachine{
		Status: kubevirtv1.VirtualMachineStatus{
			Ready: true,
		},
	}
	mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil)

	bmc := &VirtBMC{
		context:     context.TODO(),
		address:     "127.0.0.1",
		port:        623,
		vmNamespace: "default",
		vmName:      "test-vm",
		kvClient:    mockClient,
	}

	// Test getVirtualMachinePowerStatus
	status, err := bmc.getVirtualMachinePowerStatus()
	require.NoError(t, err)
	require.True(t, status)

	// Add another test case where the VM is not ready
	mockVM.Status.Ready = false
	mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil)

	status, err = bmc.getVirtualMachinePowerStatus()
	require.NoError(t, err)
	require.False(t, status)
}

// Helper function to create a fake VM and set up common mock expectations
func setupVMTest(running bool) (*VirtBMC, *MockVirtualMachineInterface, *kubevirtv1.VirtualMachine) {
	mockClient := new(MockKubevirtClient)
	mockVMInterface := new(MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	fakeVM := &kubevirtv1.VirtualMachine{
		Spec: kubevirtv1.VirtualMachineSpec{
			Running: func(b bool) *bool { return &b }(running),
		},
	}
	mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(fakeVM, nil)
	mockVMInterface.On("Update", mock.Anything, fakeVM, mock.Anything).Return(fakeVM, nil)

	bmc := &VirtBMC{
		context:     context.TODO(),
		address:     "127.0.0.1",
		port:        623,
		vmNamespace: "default",
		vmName:      "test-vm",
		kvClient:    mockClient,
	}

	return bmc, mockVMInterface, fakeVM
}

func TestStopVirtualMachine(t *testing.T) {
	bmc, mockVMInterface, fakeVM := setupVMTest(true)

	// Test stopVirtualMachine
	err := bmc.stopVirtualMachine()
	require.NoError(t, err)

	// Assertion
	mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
	mockVMInterface.AssertCalled(t, "Update", mock.Anything, fakeVM, mock.Anything)

	require.NotNil(t, fakeVM.Spec.Running)
	require.False(t, *fakeVM.Spec.Running)
}

func TestStartVirtualMachine(t *testing.T) {
	bmc, mockVMInterface, fakeVM := setupVMTest(false)

	// Test startVirtualMachine
	err := bmc.startVirtualMachine()
	require.NoError(t, err)

	// Assertion
	mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
	mockVMInterface.AssertCalled(t, "Update", mock.Anything, fakeVM, mock.Anything)

	require.NotNil(t, fakeVM.Spec.Running)
	require.True(t, *fakeVM.Spec.Running)
}

func TestRebootVirtualMachine(t *testing.T) {
	mockClient := new(MockKubevirtClient)
	mockVMIInterface := new(MockVirtualMachineInstanceInterface)
	mockClient.On("VirtualMachineInstances", "default").Return(mockVMIInterface)

	mockVMIInterface.On("Delete", mock.Anything, "test-vm", mock.Anything).Return(nil)

	bmc := &VirtBMC{
		context:     context.TODO(),
		address:     "127.0.0.1",
		port:        623,
		vmNamespace: "default",
		vmName:      "test-vm",
		kvClient:    mockClient,
	}

	// Test rebootVirtualMachine
	err := bmc.rebootVirtualMachine()
	require.NoError(t, err)

	// Assertion
	mockVMIInterface.AssertCalled(t, "Delete", mock.Anything, "test-vm", mock.Anything)
}

func TestSetVirtualMachineBootDevice(t *testing.T) {
	mockClient := new(MockKubevirtClient)
	mockVMInterface := new(MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	fakeVM := &kubevirtv1.VirtualMachine{
		Spec: kubevirtv1.VirtualMachineSpec{
			Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{
				Spec: kubevirtv1.VirtualMachineInstanceSpec{
					Domain: kubevirtv1.DomainSpec{
						Devices: kubevirtv1.Devices{
							Disks: []kubevirtv1.Disk{
								{
									Name:      "disk1",
									BootOrder: new(uint),
								},
							},
							Interfaces: []kubevirtv1.Interface{
								{
									Name:      "net1",
									BootOrder: new(uint),
								},
							},
						},
					},
				},
			},
		},
	}
	mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(fakeVM, nil)
	mockVMInterface.On("Update", mock.Anything, fakeVM, mock.Anything).Return(fakeVM, nil)

	bmc := &VirtBMC{
		context:     context.TODO(),
		address:     "127.0.0.1",
		port:        623,
		vmNamespace: "default",
		vmName:      "test-vm",
		kvClient:    mockClient,
	}

	// Test PXE boot device
	err := bmc.setVirtualMachineBootDevice(Pxe)
	require.NoError(t, err)

	require.Equal(t, uint(1), *fakeVM.Spec.Template.Spec.Domain.Devices.Interfaces[0].BootOrder)

	// Test Disk boot device
	err = bmc.setVirtualMachineBootDevice(Hdd)
	require.NoError(t, err)

	require.Equal(t, uint(1), *fakeVM.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder)

	mockVMInterface.AssertCalled(t, "Update", mock.Anything, fakeVM, mock.Anything)
}
