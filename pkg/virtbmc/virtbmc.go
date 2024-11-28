package virtbmc

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	ipmi "github.com/vmware/goipmi"

	kubevirtv1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
	"kubevirt.io/kubevirtbmc/pkg/redfish"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
)

type VMNameKey struct{}

type VMNamespaceKey struct{}

type Options struct {
	KubeconfigPath string
	Address        string
	Port           int
	RedfishPort    int
}

type KubeVirtClientInterface interface {
	VirtualMachines(namespace string) kubevirtv1.VirtualMachineInterface
	VirtualMachineInstances(namespace string) kubevirtv1.VirtualMachineInstanceInterface
}

type VirtBMC struct {
	context         context.Context
	address         string
	port            int
	redfishPort     int
	vmNamespace     string
	vmName          string
	kvClient        KubeVirtClientInterface
	resourceManager *resourcemanager.VirtualMachineResourceManager
	sim             *ipmi.Simulator
	redfishEmulator *redfish.Emulator
}

func NewVirtBMC(ctx context.Context, options Options, inCluster bool) (*VirtBMC, error) {
	kvClient := NewK8sClient(options)
	resourceManager := resourcemanager.NewVirtualMachineResourceManager(ctx, kvClient)
	return &VirtBMC{
		context:         ctx,
		address:         options.Address,
		port:            options.Port,
		redfishPort:     options.RedfishPort,
		vmNamespace:     ctx.Value(VMNamespaceKey{}).(string),
		vmName:          ctx.Value(VMNameKey{}).(string),
		kvClient:        kvClient,
		resourceManager: resourceManager,
		sim: ipmi.NewSimulator(net.UDPAddr{
			IP:   net.ParseIP(options.Address).To4(),
			Port: options.Port,
		}),
		redfishEmulator: redfish.NewEmulator(ctx, options.RedfishPort, resourceManager),
	}, nil
}

func (b *VirtBMC) register() {
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisControl, b.chassisControlHandler)
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisStatus, b.chassisStatusHandler)
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandSetSystemBootOptions, b.setSystemBootOptionsHandler)
}

func (b *VirtBMC) Run() error {
	logrus.Info("Initializing the the VirtBMC agent...")
	b.register()
	if err := b.resourceManager.Initialize(b.vmNamespace, b.vmName); err != nil {
		return fmt.Errorf("unable to initialize the resource manager: %v", err)
	}

	// Start the IPMI simulator
	if err := b.sim.Run(); err != nil {
		return fmt.Errorf("unable to run the ipmi simulator: %v", err)
	}
	logrus.Infof("Listen on %s:%d", b.address, b.port)

	// Start the Redfish emulator
	if err := b.redfishEmulator.Run(); err != nil {
		return fmt.Errorf("unable to run the redfish emulator: %v", err)
	}
	logrus.Infof("Listen on %s:%d", b.address, b.redfishPort)

	<-b.context.Done()
	logrus.Info("Gracefully shutting down the VirtBMC agent...")
	b.sim.Stop()
	b.redfishEmulator.Stop()

	return nil
}
