package resourcemanager

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	kubevirtv1 "kubevirt.io/api/core/v1"

	"kubevirt.io/kubevirtbmc/pkg/builder"
	"kubevirt.io/kubevirtbmc/pkg/fake"
)

func TestGetPowerStatus(t *testing.T) {
	mockClient := new(fake.MockKubevirtClient)
	mockVMInterface := new(fake.MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	mockVM := builder.NewVirtualMachineBuilder("default", "test-vm").Ready(true).Build()

	mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil)

	vmrm := &VirtualMachineResourceManager{
		ctx:       context.TODO(),
		kvClient:  mockClient,
		namespace: "default",
		name:      "test-vm",
	}

	// Test GetPowerStatus
	status, err := vmrm.GetPowerStatus()
	require.NoError(t, err)
	require.True(t, status)

	// Add another test case where the VM is not ready
	mockVM.Status.Ready = false

	status, err = vmrm.GetPowerStatus()
	require.NoError(t, err)
	require.False(t, status)
}

func TestPowerOn(t *testing.T) {
	mockClient := new(fake.MockKubevirtClient)
	mockVMInterface := new(fake.MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	mockVM := builder.NewVirtualMachineBuilder("default", "test-vm").Running(false).Build()

	mockVMInterface.
		On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil).
		On("Update", mock.Anything, mockVM, mock.Anything).Return(mockVM, nil)

	vmrm := &VirtualMachineResourceManager{
		ctx:       context.TODO(),
		kvClient:  mockClient,
		namespace: "default",
		name:      "test-vm",
	}

	// Test PowerOn
	err := vmrm.PowerOn()
	require.NoError(t, err)

	// Assertion
	mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
	mockVMInterface.AssertCalled(t, "Update", mock.Anything, mockVM, mock.Anything)

	require.NotNil(t, mockVM.Spec.Running)
	require.True(t, *mockVM.Spec.Running)
}

func TestPowerOff(t *testing.T) {
	mockClient := new(fake.MockKubevirtClient)
	mockVMInterface := new(fake.MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	mockVM := builder.NewVirtualMachineBuilder("default", "test-vm").Running(true).Build()

	mockVMInterface.
		On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil).
		On("Update", mock.Anything, mockVM, mock.Anything).Return(mockVM, nil)

	vmrm := &VirtualMachineResourceManager{
		ctx:       context.TODO(),
		kvClient:  mockClient,
		namespace: "default",
		name:      "test-vm",
	}

	// Test PowerOff
	err := vmrm.PowerOff()
	require.NoError(t, err)

	// Assertion
	mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
	mockVMInterface.AssertCalled(t, "Update", mock.Anything, mockVM, mock.Anything)

	require.NotNil(t, mockVM.Spec.Running)
	require.False(t, *mockVM.Spec.Running)
}

func TestPowerCycle(t *testing.T) {
	mockClient := new(fake.MockKubevirtClient)
	mockVMIInterface := new(fake.MockVirtualMachineInstanceInterface)
	mockClient.On("VirtualMachineInstances", "default").Return(mockVMIInterface)

	mockVMIInterface.On("Delete", mock.Anything, "test-vm", mock.Anything).Return(nil)

	vmrm := &VirtualMachineResourceManager{
		ctx:       context.TODO(),
		kvClient:  mockClient,
		namespace: "default",
		name:      "test-vm",
	}

	// Test PowerCycle
	err := vmrm.PowerCycle()
	require.NoError(t, err)

	// Assertion
	mockVMIInterface.AssertCalled(t, "Delete", mock.Anything, "test-vm", mock.Anything)
}

func TestSetBootDevice(t *testing.T) {
	mockClient := new(fake.MockKubevirtClient)
	mockVMInterface := new(fake.MockVirtualMachineInterface)
	mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

	mockVM := &kubevirtv1.VirtualMachine{
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

	mockVMInterface.
		On("Get", mock.Anything, "test-vm", mock.Anything).Return(mockVM, nil).
		On("Update", mock.Anything, mockVM, mock.Anything).Return(mockVM, nil)

	vmrm := &VirtualMachineResourceManager{
		ctx:       context.TODO(),
		kvClient:  mockClient,
		namespace: "default",
		name:      "test-vm",
	}

	// Test PXE boot device
	err := vmrm.SetBootDevice(BootDevicePxe)
	require.NoError(t, err)

	require.Equal(t, uint(1), *mockVM.Spec.Template.Spec.Domain.Devices.Interfaces[0].BootOrder)

	// Test Disk boot device
	err = vmrm.SetBootDevice(BootDeviceHdd)
	require.NoError(t, err)

	require.Equal(t, uint(1), *mockVM.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder)

	mockVMInterface.AssertCalled(t, "Update", mock.Anything, mockVM, mock.Anything)
}
