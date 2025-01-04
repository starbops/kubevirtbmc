package virtbmc

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	kubevirtv1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
	"kubevirt.io/kubevirtbmc/pkg/ipmi"
	"kubevirt.io/kubevirtbmc/pkg/redfish"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
)

type VMNameKey struct{}

type VMNamespaceKey struct{}

type Options struct {
	KubeconfigPath string
	Address        string
	IPMIPort       int
	RedfishPort    int
}

type KubeVirtClientInterface interface {
	VirtualMachines(namespace string) kubevirtv1.VirtualMachineInterface
	VirtualMachineInstances(namespace string) kubevirtv1.VirtualMachineInstanceInterface
}

type VirtBMC struct {
	context     context.Context
	address     string
	ipmiPort    int
	redfishPort int
	vmNamespace string
	vmName      string

	kvClient KubeVirtClientInterface

	resourceManager *resourcemanager.VirtualMachineResourceManager

	ipmiSimulator   *ipmi.Simulator
	redfishEmulator *redfish.Emulator
}

func NewVirtBMC(ctx context.Context, options Options, inCluster bool) (*VirtBMC, error) {
	kvClient := NewK8sClient(options)
	resourceManager := resourcemanager.NewVirtualMachineResourceManager(ctx, kvClient)
	return &VirtBMC{
		context:         ctx,
		address:         options.Address,
		ipmiPort:        options.IPMIPort,
		redfishPort:     options.RedfishPort,
		vmNamespace:     ctx.Value(VMNamespaceKey{}).(string),
		vmName:          ctx.Value(VMNameKey{}).(string),
		kvClient:        kvClient,
		resourceManager: resourceManager,
		ipmiSimulator:   ipmi.NewSimulator(options.Address, options.IPMIPort, resourceManager),
		redfishEmulator: redfish.NewEmulator(ctx, options.RedfishPort, resourceManager),
	}, nil
}

func (b *VirtBMC) Run() error {
	logrus.Info("Initializing the the VirtBMC agent...")

	if err := b.resourceManager.Initialize(b.vmNamespace, b.vmName); err != nil {
		return fmt.Errorf("unable to initialize the resource manager: %v", err)
	}

	// Start the IPMI simulator
	if err := b.ipmiSimulator.Run(); err != nil {
		return fmt.Errorf("unable to run the ipmi simulator: %v", err)
	}
	logrus.Infof("IPMI service listens on %s:%d", b.address, b.ipmiPort)

	// Start the Redfish emulator
	if err := b.redfishEmulator.Run(); err != nil {
		return fmt.Errorf("unable to run the redfish emulator: %v", err)
	}
	logrus.Infof("Redfish service listens on %s:%d", b.address, b.redfishPort)

	<-b.context.Done()
	logrus.Info("Gracefully shutting down the VirtBMC agent...")
	b.ipmiSimulator.Stop()
	b.redfishEmulator.Stop()

	return nil
}
