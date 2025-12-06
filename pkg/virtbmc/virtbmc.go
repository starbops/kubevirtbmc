package virtbmc

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	cdiclient "kubevirt.io/client-go/containerizeddataimporter"
	kvclient "kubevirt.io/client-go/kubevirt"

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
	BMCUser        string
	BMCPassword    string
}

type VirtBMC struct {
	context     context.Context
	address     string
	ipmiPort    int
	redfishPort int
	vmNamespace string
	vmName      string

	virtClient kvclient.Interface
	cdiClient  cdiclient.Interface

	resourceManager *resourcemanager.VirtualMachineResourceManager

	ipmiSimulator   *ipmi.Simulator
	redfishEmulator *redfish.Emulator
}

func NewVirtBMC(ctx context.Context, options Options, inCluster bool) (*VirtBMC, error) {
	virtClient := NewVirtClient(options)
	cdiClient := NewCdiClient(options)
	resourceManager := resourcemanager.NewVirtualMachineResourceManager(ctx, virtClient, cdiClient)
	return &VirtBMC{
		context:         ctx,
		address:         options.Address,
		ipmiPort:        options.IPMIPort,
		redfishPort:     options.RedfishPort,
		vmNamespace:     ctx.Value(VMNamespaceKey{}).(string),
		vmName:          ctx.Value(VMNameKey{}).(string),
		virtClient:      virtClient,
		cdiClient:       cdiClient,
		resourceManager: resourceManager,
		ipmiSimulator:   ipmi.NewSimulator(options.Address, options.IPMIPort, resourceManager),
		redfishEmulator: redfish.NewEmulator(ctx, options.RedfishPort, options.BMCUser, options.BMCPassword, resourceManager),
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
