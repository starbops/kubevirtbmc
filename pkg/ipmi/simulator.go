package ipmi

import (
	"net"

	"github.com/sirupsen/logrus"
	goipmi "github.com/vmware/goipmi"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
)

type Simulator struct {
	ip   string
	port int

	handler *handler
	sim     *goipmi.Simulator
}

func NewSimulator(ip string, port int, resourceManager resourcemanager.ResourceManager) *Simulator {
	return &Simulator{
		ip:   ip,
		port: port,

		handler: NewHandler(resourceManager),
		sim: goipmi.NewSimulator(net.UDPAddr{
			IP:   net.ParseIP(ip).To4(),
			Port: port,
		}),
	}
}

func (s *Simulator) initialize() {
	s.sim.SetHandler(
		goipmi.NetworkFunctionChassis,
		goipmi.CommandChassisStatus,
		s.handler.chassisStatusHandler,
	)
	s.sim.SetHandler(
		goipmi.NetworkFunctionChassis,
		goipmi.CommandChassisControl,
		s.handler.chassisControlHandler,
	)
	s.sim.SetHandler(
		goipmi.NetworkFunctionChassis,
		goipmi.CommandSetSystemBootOptions,
		s.handler.setSystemBootOptionsHandler,
	)
}

func (s *Simulator) Run() error {
	s.initialize()
	return s.sim.Run()
}

func (s *Simulator) Stop() {
	s.sim.Stop()
	logrus.Info("Redfish emulator gracefully stopped")
}
