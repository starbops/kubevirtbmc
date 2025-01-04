package resourcemanager

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	kubevirtv1 "kubevirt.io/api/core/v1"

	"kubevirt.io/kubevirtbmc/pkg/builder"
	"kubevirt.io/kubevirtbmc/pkg/fake"
	"kubevirt.io/kubevirtbmc/pkg/util"
)

func TestGetPowerStatus(t *testing.T) {
	testCases := []struct {
		name        string
		vm          *kubevirtv1.VirtualMachine
		expect      bool
		shouldError bool
	}{
		{
			name:        "Virtual machine is ready",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Ready(true).Build(),
			expect:      true,
			shouldError: false,
		},
		{
			name:        "Virtual machine is not ready",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Ready(false).Build(),
			expect:      false,
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := new(fake.MockKubevirtClient)
			mockVMInterface := new(fake.MockVirtualMachineInterface)
			mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

			mockVMInterface.On("Get", mock.Anything, "test-vm", mock.Anything).Return(tc.vm, nil)

			vmrm := &VirtualMachineResourceManager{
				ctx:       context.TODO(),
				kvClient:  mockClient,
				namespace: "default",
				name:      "test-vm",
			}

			// Test GetPowerStatus
			status, err := vmrm.GetPowerStatus()
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expect, status)
			}

			// Assertion
			mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
		})
	}
}

func TestPowerOn(t *testing.T) {
	testCases := []struct {
		name        string
		vm          *kubevirtv1.VirtualMachine
		expectedVM  *kubevirtv1.VirtualMachine
		shouldError bool
	}{
		{
			name:        "Power on a virtual machine that should be on should have no effect",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Running(true).Build(),
			expectedVM:  builder.NewVirtualMachineBuilder("default", "test-vm").Running(true).Build(),
			shouldError: false,
		},
		{
			name:        "Power on a virtual machine that should be off should succeed",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Running(false).Build(),
			expectedVM:  builder.NewVirtualMachineBuilder("default", "test-vm").Running(true).Build(),
			shouldError: false,
		},
		{
			name: "Power on a virtual machine whose RunStrategy is set to RerunOnFailure should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyRerunOnFailure).Build(),
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyRerunOnFailure).Build(),
			shouldError: false,
		},
		{
			name: "Power on a virtual machine whose RunStrategy is set to Halted should make it become RerunOnFailure",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyHalted).Build(),
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyRerunOnFailure).Build(),
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := new(fake.MockKubevirtClient)
			mockVMInterface := new(fake.MockVirtualMachineInterface)
			mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

			mockVMInterface.
				On("Get", mock.Anything, "test-vm", mock.Anything).Return(tc.vm, nil).
				On("Update", mock.Anything, tc.vm, mock.Anything).Return(tc.vm, nil)

			vmrm := &VirtualMachineResourceManager{
				ctx:       context.TODO(),
				kvClient:  mockClient,
				namespace: "default",
				name:      "test-vm",
			}

			// Test PowerOn
			err := vmrm.PowerOn()
			require.NoError(t, err)
			require.Equal(t, tc.expectedVM, tc.vm)

			// Assertion
			mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
			mockVMInterface.AssertCalled(t, "Update", mock.Anything, tc.vm, mock.Anything)

		})
	}
}

func TestPowerOff(t *testing.T) {
	testCases := []struct {
		name        string
		vm          *kubevirtv1.VirtualMachine
		expectedVM  *kubevirtv1.VirtualMachine
		shouldError bool
	}{
		{
			name:        "Power off a virtual machine that should be on should succeed",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Running(true).Build(),
			expectedVM:  builder.NewVirtualMachineBuilder("default", "test-vm").Running(false).Build(),
			shouldError: false,
		},
		{
			name:        "Power off a virtual machine that should be off should have no effect",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Running(false).Build(),
			expectedVM:  builder.NewVirtualMachineBuilder("default", "test-vm").Running(false).Build(),
			shouldError: false,
		},
		{
			name: "Power on a virtual machine whose RunStrategy is set to RerunOnFailure should make it become Halted",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyRerunOnFailure).Build(),
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyHalted).Build(),
			shouldError: false,
		},
		{
			name: "Power on a virtual machine whose RunStrategy is set to Halted should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyHalted).Build(),
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				RunStrategy(kubevirtv1.RunStrategyHalted).Build(),
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := new(fake.MockKubevirtClient)
			mockVMInterface := new(fake.MockVirtualMachineInterface)
			mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

			mockVMInterface.
				On("Get", mock.Anything, "test-vm", mock.Anything).Return(tc.vm, nil).
				On("Update", mock.Anything, tc.vm, mock.Anything).Return(tc.vm, nil)

			vmrm := &VirtualMachineResourceManager{
				ctx:       context.TODO(),
				kvClient:  mockClient,
				namespace: "default",
				name:      "test-vm",
			}

			// Test PowerOff
			err := vmrm.PowerOff()
			require.NoError(t, err)
			require.Equal(t, tc.expectedVM, tc.vm)

			// Assertion
			mockVMInterface.AssertCalled(t, "Get", mock.Anything, "test-vm", mock.Anything)
			mockVMInterface.AssertCalled(t, "Update", mock.Anything, tc.vm, mock.Anything)
		})
	}
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
	testCases := []struct {
		name        string
		vm          *kubevirtv1.VirtualMachine
		bootDevice  BootDevice
		expectedVM  *kubevirtv1.VirtualMachine
		shouldError bool
	}{
		{
			name: "Set boot device to HDD for a virtual machine with single disk should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with single disk whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with single disk whose boot order is already set to 2 should bring it to 1",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](2)).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with multiple disks should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", nil).
				AddDisk("test-disk-2", nil).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", util.Ptr[uint](1)).
				AddDisk("test-disk-2", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with multiple disks whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", util.Ptr[uint](1)).
				AddDisk("test-disk-2", nil).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", util.Ptr[uint](1)).
				AddDisk("test-disk-2", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with multiple disks whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", nil).
				AddDisk("test-disk-2", util.Ptr[uint](1)).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk-1", util.Ptr[uint](1)).
				AddDisk("test-disk-2", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with single interface should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", nil).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with single interface whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with single interface whose boot order is already set to 2 should bring it to 1",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", util.Ptr[uint](2)).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with multiple interfaces should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", nil).
				AddInterface("test-interface-2", nil).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", util.Ptr[uint](1)).
				AddInterface("test-interface-2", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with multiple interfaces whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", util.Ptr[uint](1)).
				AddInterface("test-interface-2", nil).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", util.Ptr[uint](1)).
				AddInterface("test-interface-2", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with multiple interfaces whose boot order is already set to 1 should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", nil).
				AddInterface("test-interface-2", util.Ptr[uint](1)).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface-1", util.Ptr[uint](1)).
				AddInterface("test-interface-2", nil).Build(),
			shouldError: false,
		},
		{
			name:        "Set boot device to HDD for a virtual machine with no disks and interfaces should fail",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Build(),
			bootDevice:  BootDeviceHdd,
			shouldError: true,
		},
		{
			name: "Set boot device to HDD for a virtual machine with no disks but interfaces should fail",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddInterface("test-interface", nil).Build(),
			bootDevice:  BootDeviceHdd,
			shouldError: true,
		},
		{
			name:        "Set boot device to PXE for a virtual machine with no disks and interfaces should fail",
			vm:          builder.NewVirtualMachineBuilder("default", "test-vm").Build(),
			bootDevice:  BootDevicePxe,
			shouldError: true,
		},
		{
			name: "Set boot device to PXE for a virtual machine with no interfaces but disks should fail",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).Build(),
			bootDevice:  BootDevicePxe,
			shouldError: true,
		},
		{
			name: "Set boot device to HDD for a virtual machine with disks and interfaces but no boot order specified should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", nil).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).
				AddInterface("test-interface", nil).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to HDD for a virtual machine with disks and interfaces and with first boot order set to disk should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).
				AddInterface("test-interface", nil).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).
				AddInterface("test-interface", nil).Build(),
			shouldError: false,
		},
		// {
		// 	name: "Set boot device to HDD for a virtual machine with disks and interfaces and with first boot order set to disk should have no effect",
		// 	vm: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](1)).
		// 		AddInterface("test-interface", util.Ptr[uint](2)).Build(),
		// 	bootDevice: BootDeviceHdd,
		// 	expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](1)).
		// 		AddInterface("test-interface", util.Ptr[uint](2)).Build(),
		// 	shouldError: false,
		// },
		{
			name: "Set boot device to HDD for a virtual machine with disks and interfaces and with first boot order set to interface should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			bootDevice: BootDeviceHdd,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).
				AddInterface("test-interface", nil).Build(),
			shouldError: false,
		},
		// {
		// 	name: "Set boot device to HDD for a virtual machine with disks and interfaces and with first boot order set to interface should succeed",
		// 	vm: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](2)).
		// 		AddInterface("test-interface", util.Ptr[uint](1)).Build(),
		// 	bootDevice: BootDeviceHdd,
		// 	expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](1)).
		// 		AddInterface("test-interface", util.Ptr[uint](2)).Build(),
		// 	shouldError: false,
		// },
		{
			name: "Set boot device to PXE for a virtual machine with disks and interfaces but no boot order specified should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", nil).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		{
			name: "Set boot device to PXE for a virtual machine with disks and interfaces and with first boot order set to disk should succeed",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", util.Ptr[uint](1)).
				AddInterface("test-interface", nil).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		// {
		// 	name: "Set boot device to PXE for a virtual machine with disks and interfaces and with first boot order set to disk should succeed",
		// 	vm: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](1)).
		// 		AddInterface("test-interface", util.Ptr[uint](2)).Build(),
		// 	bootDevice: BootDevicePxe,
		// 	expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](2)).
		// 		AddInterface("test-interface", util.Ptr[uint](1)).Build(),
		// 	shouldError: false,
		// },
		{
			name: "Set boot device to PXE for a virtual machine with disks and interfaces and with first boot order set to interface should have no effect",
			vm: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			bootDevice: BootDevicePxe,
			expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
				AddDisk("test-disk", nil).
				AddInterface("test-interface", util.Ptr[uint](1)).Build(),
			shouldError: false,
		},
		// {
		// 	name: "Set boot device to PXE for a virtual machine with disks and interfaces and with first boot order set to interface should have no effect",
		// 	vm: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](2)).
		// 		AddInterface("test-interface", util.Ptr[uint](1)).Build(),
		// 	bootDevice: BootDevicePxe,
		// 	expectedVM: builder.NewVirtualMachineBuilder("default", "test-vm").
		// 		AddDisk("test-disk", util.Ptr[uint](2)).
		// 		AddInterface("test-interface", util.Ptr[uint](1)).Build(),
		// 	shouldError: false,
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := new(fake.MockKubevirtClient)
			mockVMInterface := new(fake.MockVirtualMachineInterface)
			mockClient.On("VirtualMachines", "default").Return(mockVMInterface)

			mockVMInterface.
				On("Get", mock.Anything, "test-vm", mock.Anything).Return(tc.vm, nil).
				On("Update", mock.Anything, tc.vm, mock.Anything).Return(tc.vm, nil)

			vmrm := &VirtualMachineResourceManager{
				ctx:       context.TODO(),
				kvClient:  mockClient,
				namespace: "default",
				name:      "test-vm",
			}

			// Test SetBootDevice
			err := vmrm.SetBootDevice(tc.bootDevice)
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedVM, tc.vm)

				// Assertion
				mockVMInterface.AssertCalled(t, "Update", mock.Anything, tc.vm, mock.Anything)
			}
		})
	}
}
