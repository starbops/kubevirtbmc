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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const BMCDeploymentNameSuffix = "-bmc-deployment"

// deploymentForVirtBMC returns a Deployment object for VirtualMachineBMC
func (r *VirtualMachineBMCReconciler) NewDeployment(bmc *bmcv1beta1.VirtualMachineBMC) *appsv1.Deployment {
	podSpec := r.NewPodSpec(bmc)
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmc.Name + BMCDeploymentNameSuffix,
			Namespace: bmc.Namespace,
			Labels:    LabelsForBMC(bmc.Name, bmc.Spec.VirtualMachineRef.Name),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.To[int32](1),
			Selector: &metav1.LabelSelector{
				MatchLabels: LabelsForBMC(bmc.Name, bmc.Spec.VirtualMachineRef.Name),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      bmc.Name,
					Namespace: bmc.Namespace,
					Labels:    LabelsForBMC(bmc.Name, bmc.Spec.VirtualMachineRef.Name),
				},
				Spec: podSpec,
			},
		},
	}

	if err := controllerutil.SetControllerReference(bmc, dep, r.Scheme); err != nil {
		return nil
	}

	return dep
}

func (r *VirtualMachineBMCReconciler) ReconcileDeployment(ctx context.Context, virtBMC *bmcv1beta1.VirtualMachineBMC) (ctrl.Result, error) {
	foundDep := &appsv1.Deployment{}

	depName := types.NamespacedName{
		Name:      virtBMC.Name + BMCDeploymentNameSuffix,
		Namespace: virtBMC.Namespace,
	}
	if err := r.Get(ctx, depName, foundDep); err != nil {
		if apierrors.IsNotFound(err) {
			dep := r.NewDeployment(virtBMC)
			r.Log.Info("Creating NEW Deployment", "Deployment", depName)
			if err := r.Create(ctx, dep); err != nil {
				r.Log.Error(err, "Failed to create Deployment", "Deployment", depName)
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true}, nil
		}
		r.Log.Error(err, "Failed to get Deployment", "Deployment", depName)
		return ctrl.Result{}, err
	}

	size := int32(1)
	if *foundDep.Spec.Replicas != size {
		foundDep.Spec.Replicas = &size
		if err := r.Update(ctx, foundDep); err != nil {
			r.Log.Error(err, "Failed to update Deployment replicas", "Deployment", depName)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}
	r.Log.Info("Deployment already exists and is up to date", "Deployment", depName)
	return ctrl.Result{}, nil
}
