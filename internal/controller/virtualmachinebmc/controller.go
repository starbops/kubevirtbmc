/*
Copyright 2024.

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
	"strconv"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	virtualmachinev1 "kubevirt.io/kubevirtbmc/api/v1alpha1"
)

// VirtualMachineBMCReconciler reconciles a VirtualMachineBMC object
type VirtualMachineBMCReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	AgentImageName string
	AgentImageTag  string
}

var (
	ownerKey = ".metadata.controller"
	apiGVStr = virtualmachinev1.GroupVersion.String()
)

func (r *VirtualMachineBMCReconciler) constructPodFromVirtualMachineBMC(virtualMachineBMC *virtualmachinev1.VirtualMachineBMC) *corev1.Pod {
	name := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Name)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				VirtualMachineBMCNameLabel: virtualMachineBMC.Name,
				VMNameLabel:                virtualMachineBMC.Spec.VirtualMachineName,
			},
			Name:      name,
			Namespace: VirtualMachineBMCNamespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  virtBMCContainerName,
					Image: fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag),
					Args: []string{
						"--address",
						"0.0.0.0",
						"--ipmi-port",
						strconv.Itoa(ipmiPort),
						"--redfish-port",
						strconv.Itoa(redfishPort),
						virtualMachineBMC.Spec.VirtualMachineNamespace,
						virtualMachineBMC.Spec.VirtualMachineName,
					},
					Ports: []corev1.ContainerPort{
						{
							Name:          ipmiPortName,
							ContainerPort: ipmiPort,
							Protocol:      corev1.ProtocolUDP,
						},
						{
							Name:          redfishPortName,
							ContainerPort: redfishPort,
							Protocol:      corev1.ProtocolTCP,
						},
					},
				},
			},
			ServiceAccountName: "kubevirtbmc-virtbmc",
		},
	}

	return pod
}

func (r *VirtualMachineBMCReconciler) constructServiceFromVirtualMachineBMC(virtualMachineBMC *virtualmachinev1.VirtualMachineBMC) *corev1.Service {
	name := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Name)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				VirtualMachineBMCNameLabel: virtualMachineBMC.Name,
				VMNameLabel:                virtualMachineBMC.Spec.VirtualMachineName,
			},
			Name:      name,
			Namespace: VirtualMachineBMCNamespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				VirtualMachineBMCNameLabel: virtualMachineBMC.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       ipmiPortName,
					Protocol:   corev1.ProtocolUDP,
					TargetPort: intstr.FromString(ipmiPortName),
					Port:       IPMISvcPort,
				},
				{
					Name:       redfishPortName,
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromString(redfishPortName),
					Port:       RedfishSvcPort,
				},
			},
		},
	}

	return svc
}

//+kubebuilder:rbac:groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs/finalizers,verbs=update
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

	var virtualMachineBMC virtualmachinev1.VirtualMachineBMC
	if err := r.Get(ctx, req.NamespacedName, &virtualMachineBMC); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	desiredPod := r.constructPodFromVirtualMachineBMC(&virtualMachineBMC)
	if err := ctrl.SetControllerReference(&virtualMachineBMC, desiredPod, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	existingPod := &corev1.Pod{}
	podKey := client.ObjectKey{Name: desiredPod.Name, Namespace: desiredPod.Namespace}

	err := r.Get(ctx, podKey, existingPod)
	podExists := err == nil

	if err != nil && !apierrors.IsNotFound(err) {
		log.Error(err, "unable to fetch existing Pod", "pod", podKey)
		return ctrl.Result{}, err
	}

	if podExists {
		desiredImage := fmt.Sprintf("%s:%s", r.AgentImageName, r.AgentImageTag)
		existingImage := existingPod.Spec.Containers[0].Image

		if existingImage != desiredImage {
			log.Info("agent pod image mismatch detected, deleting pod for rollout",
				"pod", podKey,
				"existingImage", existingImage,
				"desiredImage", desiredImage)

			if err := r.Delete(ctx, existingPod); err != nil {
				if apierrors.IsNotFound(err) {
					log.Info("Pod already deleted by another reconcile, continuing", "pod", podKey)
				} else {
					log.Error(err, "unable to delete Pod", "pod", podKey)
					return ctrl.Result{}, err
				}
			}
			return ctrl.Result{Requeue: true}, nil
		}

		log.V(1).Info("Pod already exists with correct image", "pod", podKey)

	} else {
		if err := r.Create(ctx, desiredPod); err != nil {
			log.Error(err, "unable to create Pod", "pod", desiredPod)
			return ctrl.Result{}, err
		}
		log.V(1).Info("created Pod", "pod", desiredPod)
	}
	// Prepare the virtBMC Service
	svc := r.constructServiceFromVirtualMachineBMC(&virtualMachineBMC)
	if err := ctrl.SetControllerReference(&virtualMachineBMC, svc, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Create the virtBMC Service on the cluster
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
