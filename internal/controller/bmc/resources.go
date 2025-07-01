package bmc

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

// serviceForVirtualMachineBMC returns the desired Service object.
func (r *VirtualMachineBMCReconciler) serviceForVirtualMachineBMC(bmc *bmcv1beta1.VirtualMachineBMC) *corev1.Service {
	labels := map[string]string{
		AppLabelKey: bmc.Name + BMCProxyLabelSuffix,
	}

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmc.Name + BMCServiceNameSuffix,
			Namespace: bmc.Namespace,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bmc, bmcv1beta1.GroupVersion.WithKind("VirtualMachineBMC")),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Name:     BMCPortName,
					Port:     BMCPortNumber,
					Protocol: corev1.ProtocolTCP,
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
}

// deploymentForVirtualMachineBMC returns the desired Deployment object.
func (r *VirtualMachineBMCReconciler) deploymentForVirtualMachineBMC(bmc *bmcv1beta1.VirtualMachineBMC, vm *kubevirtv1.VirtualMachine, secret *corev1.Secret) *appsv1.Deployment {
	labels := map[string]string{
		AppLabelKey: bmc.Name + BMCProxyLabelSuffix,
	}
	replicas := int32(1)

	image := r.AgentImageName
	if r.AgentImageTag != "" {
		image = fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag)
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmc.Name + BMCProxyLabelSuffix,
			Namespace: bmc.Namespace,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bmc, bmcv1beta1.GroupVersion.WithKind("VirtualMachineBMC")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  BMCContainerName,
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									Name:          BMCPortName,
									ContainerPort: BMCPortNumber,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "VIRTUALMACHINE_NAME",
									Value: vm.Name,
								},
								{
									Name:  "VIRTUALMACHINE_NAMESPACE",
									Value: vm.Namespace,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      AuthSecretVolumeName,
									MountPath: AuthSecretMountPath,
									ReadOnly:  true,
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: AuthSecretVolumeName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: secret.Name,
								},
							},
						},
					},
				},
			},
		},
	}
}
