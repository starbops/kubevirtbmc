package bmc

import (
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

const (
	VirtualMachineBMCNameLabel = "kubevirtbmc.io/virtualmachinebmc"
	VMNameLabel                = "kubevirtbmc.io/vmname"
	virtBMCContainerName       = "virtbmc"
	ServiceAccountName         = "kubevirtbmc-virtbmc"
)

func (r *VirtualMachineBMCReconciler) NewPodSpec(bmc *bmcv1beta1.VirtualMachineBMC) corev1.PodSpec {
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  virtBMCContainerName,
				Image: fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag),
				Args: []string{
					"--address",
					"0.0.0.0",
					"--ipmi-port",
					strconv.Itoa(ipmiPortNumber),
					"--redfish-port",
					strconv.Itoa(redfishPortNumber),
					bmc.Spec.VirtualMachineRef.Name,
					bmc.Spec.AuthSecretRef.Name,
				},
				Ports: []corev1.ContainerPort{
					{
						Name:          ipmiPortName,
						ContainerPort: ipmiPortNumber,
						Protocol:      corev1.ProtocolUDP,
					},
					{
						Name:          redfishPortName,
						ContainerPort: redfishPortNumber,
						Protocol:      corev1.ProtocolTCP,
					},
				},
				LivenessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: "/livez",
							Port: intstr.FromInt(redfishPortNumber),
						},
					},
					InitialDelaySeconds: 10,
					PeriodSeconds:       10,
					SuccessThreshold:    1,
					TimeoutSeconds:      1,
					FailureThreshold:    3,
				},
				ReadinessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: "/healthz",
							Port: intstr.FromInt(redfishPortNumber),
						},
					},
					InitialDelaySeconds: 5,
					PeriodSeconds:       10,
					SuccessThreshold:    1,
					TimeoutSeconds:      1,
					FailureThreshold:    3,
				},
			},
		},

		ServiceAccountName: ServiceAccountName,
	}
}
