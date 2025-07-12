/*
 * This file is part of the KubeVirt/KubeVirtBMC project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 *
 */
package bmc

import (
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

func (r *VirtualMachineBMCReconciler) NewPodSpec(bmc *bmcv1beta1.VirtualMachineBMC) corev1.PodSpec {
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  r.AgentImageName,
				Image: fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag),
				Args: []string{
					"--address",
					"0.0.0.0",
					"--ipmi-port",
					strconv.Itoa(IpmiContainerPort),
					"--redfish-port",
					strconv.Itoa(RedfishContainerPort),
					bmc.Spec.VirtualMachineRef.Name,
					bmc.Spec.AuthSecretRef.Name,
				},
				Ports: []corev1.ContainerPort{
					{
						Name:          ipmiPortName,
						ContainerPort: IpmiContainerPort,
						Protocol:      corev1.ProtocolUDP,
					},
					{
						Name:          redfishPortName,
						ContainerPort: RedfishContainerPort,
						Protocol:      corev1.ProtocolTCP,
					},
				},
				LivenessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: "/livez",
							Port: intstr.FromInt(RedfishContainerPort),
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
							Port: intstr.FromInt(RedfishContainerPort),
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
