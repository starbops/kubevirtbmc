/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package virtualmachinebmc

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	virtualmachinev1 "kubevirt.org/kubevirtbmc/api/v1"
)

// VirtualMachineBMCReconciler reconciles a VirtualMachineBMC object
type VirtualMachineBMCReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var (
	ownerKey = ".metadata.controller"
	apiGVStr = virtualmachinev1.GroupVersion.String()
)

func (r *VirtualMachineBMCReconciler) constructPodForVirtualMachineBMC(kubeBMC *virtualmachinev1.VirtualMachineBMC) (*corev1.Pod, error) {
	name := fmt.Sprintf("%s-virt-bmc", kubeBMC.Name)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels: map[string]string{
				ManagedByLabel: "virt-bmc",
				KBMCNameLabel:  kubeBMC.Name,
				VMNameLabel:    kubeBMC.Spec.VirtualMachineName,
			},
			Name:      name,
			Namespace: KBMCNamespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  kbmcContainerName,
					Image: fmt.Sprintf("%s:%s", kbmcImageName, kbmcImageTag),
					Args: []string{
						"--address",
						"0.0.0.0",
						"--port",
						"623",
						kubeBMC.Spec.VirtualMachineNamespace,
						kubeBMC.Spec.VirtualMachineName,
					},
					Ports: []corev1.ContainerPort{
						{
							Name:          ipmiPortName,
							ContainerPort: ipmiPort,
							Protocol:      corev1.ProtocolUDP,
						},
					},
				},
			},
			ServiceAccountName: "kubevirtbmc-virt-bmc",
		},
	}

	return pod, nil
}

func (r *VirtualMachineBMCReconciler) constructServiceForVirtualMachineBMC(kubeBMC *virtualmachinev1.VirtualMachineBMC) (*corev1.Service, error) {
	name := fmt.Sprintf("%s-virt-bmc", kubeBMC.Name)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels: map[string]string{
				ManagedByLabel: "virt-bmc",
				KBMCNameLabel:  kubeBMC.Name,
				VMNameLabel:    kubeBMC.Spec.VirtualMachineName,
			},
			Name:      name,
			Namespace: KBMCNamespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				ManagedByLabel: "virt-bmc",
				KBMCNameLabel:  kubeBMC.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       ipmiPortName,
					Protocol:   corev1.ProtocolUDP,
					TargetPort: intstr.FromString(ipmiPortName),
					Port:       IPMISvcPort,
				},
			},
		},
	}

	return svc, nil
}

//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VirtualMachineBMC object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *VirtualMachineBMCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var kubeBMC virtualmachinev1.VirtualMachineBMC
	if err := r.Get(ctx, req.NamespacedName, &kubeBMC); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Prepare the virt-bmc Pod
	pod, err := r.constructPodForVirtualMachineBMC(&kubeBMC)
	if err != nil {
		log.Error(err, "unable to construct pod from template")
		return ctrl.Result{}, err
	}
	if err := ctrl.SetControllerReference(&kubeBMC, pod, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Create the virt-bmc Pod on the cluster
	if err := r.Create(ctx, pod); err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create Pod for VirtualMachineBMC", "pod", pod)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created Pod for VirtualMachineBMC", "pod", pod)

	// Prepare the virt-bmc Service
	svc, err := r.constructServiceForVirtualMachineBMC(&kubeBMC)
	if err != nil {
		log.Error(err, "unable to construct svc from template")
		return ctrl.Result{}, err
	}
	if err := ctrl.SetControllerReference(&kubeBMC, svc, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Create the virt-bmc Service on the cluster
	if err := r.Create(ctx, svc); err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create Service for VirtualMachineBMC", "svc", svc)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created Service for VirtualMachineBMC", "svc", svc)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VirtualMachineBMCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &corev1.Pod{}, ownerKey, func(rawObj client.Object) []string {
		// grab the pod object, extract the owner...
		pod := rawObj.(*corev1.Pod)
		owner := metav1.GetControllerOf(pod)
		if owner == nil {
			return nil
		}
		// ...make sure it's a VirtualMachineBMC...
		if owner.APIVersion != apiGVStr || owner.Kind != "VirtualMachineBMC" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}

	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &corev1.Service{}, ownerKey, func(rawObj client.Object) []string {
		// grab the svc object, extract the owner...
		svc := rawObj.(*corev1.Service)
		owner := metav1.GetControllerOf(svc)
		if owner == nil {
			return nil
		}
		// ...make sure it's a VirtualMachineBMC...
		if owner.APIVersion != apiGVStr || owner.Kind != "VirtualMachineBMC" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&virtualmachinev1.VirtualMachineBMC{}).
		Owns(&corev1.Pod{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
