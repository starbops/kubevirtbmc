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

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// deploymentForVirtBMC returns a Deployment object for VirtualMachineBMC
func (r *VirtualMachineBMCReconciler) NewDeployment(bmc *bmcv1beta1.VirtualMachineBMC) *appsv1.Deployment {
	replicas := int32(DefaultReplicas)
	podSpec := r.NewPodSpec(bmc)
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmc.Name + BMCProxyLabelSuffix,
			Namespace: bmc.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{AppLabelKey: bmc.Name + BMCProxyLabelSuffix},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{AppLabelKey: bmc.Name + BMCProxyLabelSuffix},
				},
				Spec: podSpec,
			},
		},
	}
	controllerutil.SetControllerReference(bmc, dep, r.Scheme)

	return dep
}

func (r *VirtualMachineBMCReconciler) deleteDeployment(ctx context.Context, bmc *bmcv1beta1.VirtualMachineBMC, log logr.Logger) error {
	log.Info("Deleting deployment for VirtualMachineBMC", "name", bmc.Name)
	deploy := &appsv1.Deployment{}
	if err := r.Get(ctx, client.ObjectKey{
		Name:      bmc.Name + BMCProxyLabelSuffix,
		Namespace: bmc.Namespace,
	}, deploy); err == nil {
		if err := r.Delete(ctx, deploy); err != nil {
			return err
		}
	}

	log.Info("Deleted Deployment", "name", deploy.Name)

	return nil
}
