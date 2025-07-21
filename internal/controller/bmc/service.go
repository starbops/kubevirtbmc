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
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	IpmiContainerPort    = 10623
	RedfishContainerPort = 10080
	ipmiServicePort      = 623
	redfishServicePort   = 80
	ipmiPortName         = "ipmi"
	redfishPortName      = "redfish"
	BMCServiceNameSuffix = "-bmc-service"
)

func (r *VirtualMachineBMCReconciler) NewService(bmc *bmcv1beta1.VirtualMachineBMC) *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmc.Name + BMCServiceNameSuffix,
			Namespace: bmc.Namespace,
			Labels:    LabelsForBMC(bmc.Name, bmc.Spec.VirtualMachineRef.Name),
		},
		Spec: corev1.ServiceSpec{
			Selector: LabelsForBMC(bmc.Name, bmc.Spec.VirtualMachineRef.Name),
			Ports: []corev1.ServicePort{
				{
					Name:       ipmiPortName,
					Port:       ipmiServicePort,
					Protocol:   corev1.ProtocolUDP,
					TargetPort: intstr.FromInt(IpmiContainerPort),
				},
				{
					Name:       redfishPortName,
					Port:       redfishServicePort,
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromInt(RedfishContainerPort),
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	if err := controllerutil.SetControllerReference(bmc, svc, r.Scheme); err != nil {
		return nil
	}

	return svc
}

func (r *VirtualMachineBMCReconciler) ReconcileService(ctx context.Context, virtBMC *bmcv1beta1.VirtualMachineBMC) (*corev1.Service, ctrl.Result, error) {
	foundSvc := &corev1.Service{}
	svcName := types.NamespacedName{
		Name:      virtBMC.Name + BMCServiceNameSuffix,
		Namespace: virtBMC.Namespace,
	}
	if err := r.Get(ctx, svcName, foundSvc); err != nil {
		if apierrors.IsNotFound(err) {
			svc := r.NewService(virtBMC)
			r.Log.Info("Creating Service", "Service", svcName)
			if err := r.Create(ctx, svc); err != nil {
				r.Log.Error(err, "Failed to create Service", "Service", svcName)
				return nil, ctrl.Result{}, err
			}
			return svc, ctrl.Result{Requeue: true}, nil
		}
		r.Log.Error(err, "Failed to get Service", "Service", svcName)
		return nil, ctrl.Result{}, err
	}
	return foundSvc, ctrl.Result{}, nil
}
