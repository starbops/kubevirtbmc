package kbmc

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	ipmi "github.com/vmware/goipmi"
	kubevirtv1 "kubevirt.org/virtualmachinebmc/pkg/generated/clientset/versioned/typed/core/v1"
)

type Options struct {
	KubeconfigPath string
	Address        string
	Port           int
}

type KBMC struct {
	context     context.Context
	address     string
	port        int
	vmNamespace string
	vmName      string
	kvClient    *kubevirtv1.KubevirtV1Client
	sim         *ipmi.Simulator
}

func NewKBMC(ctx context.Context, options Options, inCluster bool) (*KBMC, error) {
	return &KBMC{
		context:     ctx,
		address:     options.Address,
		port:        options.Port,
		vmNamespace: ctx.Value("VMNamespace").(string),
		vmName:      ctx.Value("VMName").(string),
		kvClient:    NewK8sClient(options),
		sim: ipmi.NewSimulator(net.UDPAddr{
			IP:   net.ParseIP(options.Address).To4(),
			Port: options.Port,
		}),
	}, nil
}

func (k *KBMC) register() {
	k.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisControl, k.chassisControlHandler)
	k.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandChassisStatus, k.chassisStatusHandler)
	k.sim.SetHandler(ipmi.NetworkFunctionChassis, ipmi.CommandSetSystemBootOptions, k.setSystemBootOptionsHandler)
}

func (k *KBMC) Run() error {
	logrus.Info("Initializing the simulator...")
	k.register()

	if err := k.sim.Run(); err != nil {
		return fmt.Errorf("unable to run the ipmi simulator")
	}
	logrus.Infof("Listen on %s:%d", k.address, k.port)

	<-k.context.Done()
	logrus.Info("Gracefully shutting down IPMIService")
	k.sim.Stop()

	return nil
}
