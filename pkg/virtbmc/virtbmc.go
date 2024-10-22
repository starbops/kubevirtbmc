package virtbmc

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	ipmi "github.com/vmware/goipmi"
	kubevirtv1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type VMNameKey struct{}

type VMNamespaceKey struct{}

type Options struct {
	KubeconfigPath string
	Address        string
	Port           int
}

type KubeVirtClientInterface interface {
	VirtualMachines(namespace string) kubevirtv1.VirtualMachineInterface
	VirtualMachineInstances(namespace string) kubevirtv1.VirtualMachineInstanceInterface
}

type VirtBMC struct {
	context     context.Context
	address     string
	port        int
	vmNamespace string
	vmName      string
	kvClient    KubeVirtClientInterface
	sim         *ipmi.Simulator
}

func NewVirtBMC(ctx context.Context, options Options, inCluster bool) (*VirtBMC, error) {
	return &VirtBMC{
		context:     ctx,
		address:     options.Address,
		port:        options.Port,
		vmNamespace: ctx.Value(VMNamespaceKey{}).(string),
		vmName:      ctx.Value(VMNameKey{}).(string),
		kvClient:    NewK8sClient(options),
		sim: ipmi.NewSimulator(net.UDPAddr{
			IP:   net.ParseIP(options.Address).To4(),
			Port: options.Port,
		}),
	}, nil
}

func (b *VirtBMC) register() {
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisControl, b.chassisControlHandler)
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisStatus, b.chassisStatusHandler)
	b.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandSetSystemBootOptions, b.setSystemBootOptionsHandler)
}

func (b *VirtBMC) Run() error {
	logrus.Info("Initializing the simulator...")
	b.register()

	if err := b.sim.Run(); err != nil {
		return fmt.Errorf("unable to run the ipmi simulator")
	}
	logrus.Infof("Listen on %s:%d", b.address, b.port)

	<-b.context.Done()
	logrus.Info("Gracefully shutting down IPMIService")
	b.sim.Stop()

	return nil
}
