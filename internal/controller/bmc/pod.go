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
	"os"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

type agentImageConfig struct {
	ImageName     string
	ImageTag      string
	ContainerName string
	FullImage     string
}

func NewAgentImageConfig() agentImageConfig {
	registryHost := os.Getenv("REGISTRY_HOST")
	if registryHost == "" {
		registryHost = "registry"
	}

	registryPort := os.Getenv("REGISTRY_PORT")
	if registryPort == "" {
		registryPort = "5000"
	}

	imageRepo := os.Getenv("VIRTBMC_IMAGE_NAME")
	if imageRepo == "" {
		imageRepo = "virtbmc"
	}

	imageTag := os.Getenv("VIRTBMC_IMAGE_TAG")
	if imageTag == "" {
		imageTag = "latest"
	}
	registry := fmt.Sprintf("%s:%s", registryHost, registryPort) // e.g. "registry:5000"
	imageName := fmt.Sprintf("%s/%s", registry, imageRepo)       // e.g. "registry:5000/virtbmc"

	return agentImageConfig{
		ImageName:     imageName,
		ImageTag:      imageTag,
		ContainerName: imageRepo,
		FullImage:     fmt.Sprintf("%s:%s", imageName, imageTag), // e.g. "registry:5000/virtbmc:latest"
	}
}

func (r *VirtualMachineBMCReconciler) NewPodSpec(bmc *bmcv1beta1.VirtualMachineBMC) corev1.PodSpec {
	fmt.Printf("DEBUG: Using ContainerName=%s, Image=%s\n", r.AgentImageName.ContainerName, r.AgentImageName.FullImage)
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:            r.AgentImageName.ContainerName,
				Image:           r.AgentImageName.FullImage,
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
				// LivenessProbe: &corev1.Probe{
				// 	ProbeHandler: corev1.ProbeHandler{
				// 		HTTPGet: &corev1.HTTPGetAction{
				// 			Path: "/livez",
				// 			Port: intstr.FromInt(RedfishContainerPort),
				// 		},
				// 	},
				// 	InitialDelaySeconds: 10,
				// 	PeriodSeconds:       10,
				// 	SuccessThreshold:    1,
				// 	TimeoutSeconds:      1,
				// 	FailureThreshold:    3,
				// },
				// ReadinessProbe: &corev1.Probe{
				// 	ProbeHandler: corev1.ProbeHandler{
				// 		HTTPGet: &corev1.HTTPGetAction{
				// 			Path: "/healthz",
				// 			Port: intstr.FromInt(RedfishContainerPort),
				// 		},
				// 	},
				// 	InitialDelaySeconds: 5,
				// 	PeriodSeconds:       10,
				// 	SuccessThreshold:    1,
				// 	TimeoutSeconds:      1,
				// 	FailureThreshold:    3,
				// },
			},
		},

		ServiceAccountName: ServiceAccountName,
	}
}
