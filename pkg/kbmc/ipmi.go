package kbmc

import (
	"github.com/sirupsen/logrus"
	ipmi "github.com/vmware/goipmi"
)

func (k *KBMC) chassisControlHandler(m *ipmi.Message) ipmi.Response {
	r := &ipmi.ChassisControlRequest{}
	if err := m.Request(r); err != nil {
		return err
	}

	var err error

	switch r.ChassisControl {
	case ipmi.ControlPowerDown, ipmi.ControlPowerAcpiSoft:
		logrus.Info("power off")
		err = k.stopVirtualMachine()
	case ipmi.ControlPowerUp:
		logrus.Info("power on")
		err = k.startVirtualMachine()
	case ipmi.ControlPowerCycle, ipmi.ControlPowerHardReset:
		logrus.Info("power cycle")
		err = k.rebootVirtualMachine()
	}

	if err != nil {
		return &ipmi.ChassisControlResponse{
			CompletionCode: ipmi.ErrInvalidState,
		}
	}

	return &ipmi.ChassisControlResponse{
		CompletionCode: ipmi.CommandCompleted,
	}
}

func (k *KBMC) chassisStatusHandler(*ipmi.Message) ipmi.Response {
	logrus.Info("power status")

	isUp, err := k.getVirtualMachinePowerStatus()
	if err != nil {
		return &ipmi.ChassisStatusResponse{
			CompletionCode: ipmi.ErrInvalidState,
		}
	}
	if !isUp {
		return &ipmi.ChassisStatusResponse{
			CompletionCode: ipmi.CommandCompleted,
			PowerState:     ipmi.PowerOverload,
		}
	}
	return &ipmi.ChassisStatusResponse{
		CompletionCode: ipmi.CommandCompleted,
		PowerState:     ipmi.SystemPower,
	}
}
