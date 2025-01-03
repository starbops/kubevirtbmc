package fake

import (
	"github.com/stretchr/testify/mock"
	kubevirttypev1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type MockKubevirtClient struct {
	mock.Mock
}

func (m *MockKubevirtClient) VirtualMachines(namespace string) kubevirttypev1.VirtualMachineInterface {
	args := m.Called(namespace)
	return args.Get(0).(kubevirttypev1.VirtualMachineInterface)
}

func (m *MockKubevirtClient) VirtualMachineInstances(namespace string) kubevirttypev1.VirtualMachineInstanceInterface {
	args := m.Called(namespace)
	return args.Get(0).(kubevirttypev1.VirtualMachineInstanceInterface)
}
