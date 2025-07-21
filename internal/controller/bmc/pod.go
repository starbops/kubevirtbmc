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
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

const (
	ServiceAccountName   = "kubevirtbmc-virtbmc"
	virtBMCContainerName = "virtbmc"
	VirtBMCImageName     = "starbops/virtbmc"
)

func (r *VirtualMachineBMCReconciler) NewPodSpec(bmc *bmcv1beta1.VirtualMachineBMC) corev1.PodSpec {
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:            virtBMCContainerName,
				Image:           fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag),
				ImagePullPolicy: corev1.PullAlways,
				Args: []string{
					"--address",
					"0.0.0.0",
					"--ipmi-port",
					strconv.Itoa(IpmiContainerPort),
					"--redfish-port",
					strconv.Itoa(RedfishContainerPort),
					bmc.Namespace,
					bmc.Spec.VirtualMachineRef.Name,
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
			},
		},

		ServiceAccountName: ServiceAccountName,
	}
}
