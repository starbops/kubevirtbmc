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
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubevirtv1 "kubevirt.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	bmcv1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

// VirtualMachineBMCReconciler reconciles a VirtualMachineBMC object
type VirtualMachineBMCReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	AgentImageName string
	AgentImageTag  string
}

const (
	clusterRoleName = "kubevirtbmc-virtbmc-role"
)

var (
	ownerKey = ".metadata.controller"
	apiGVStr = bmcv1.GroupVersion.String()
)

func (r *VirtualMachineBMCReconciler) constructServiceAccount(virtualMachineBMC *bmcv1.VirtualMachineBMC) *corev1.ServiceAccount {
	serviceAccountName := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Spec.VirtualMachineRef.Name)
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: virtualMachineBMC.Namespace,
		},
	}
	return sa
}

func (r *VirtualMachineBMCReconciler) constructRoleBinding(virtualMachineBMC *bmcv1.VirtualMachineBMC) *rbacv1.RoleBinding {
	serviceAccountName := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Spec.VirtualMachineRef.Name)
	roleBindingName := fmt.Sprintf("%s-virtbmc-rolebinding", virtualMachineBMC.Spec.VirtualMachineRef.Name)
	rb := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleBindingName,
			Namespace: virtualMachineBMC.Namespace,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     clusterRoleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      serviceAccountName,
				Namespace: virtualMachineBMC.Namespace,
			},
		},
	}
	return rb
}

func (r *VirtualMachineBMCReconciler) ensureRBACResources(ctx context.Context, virtualMachineBMC *bmcv1.VirtualMachineBMC) error {
	log := log.FromContext(ctx)

	sa := r.constructServiceAccount(virtualMachineBMC)

	if err := ctrl.SetControllerReference(virtualMachineBMC, sa, r.Scheme); err != nil {
		return err
	}

	if err := r.Create(ctx, sa); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			log.Error(err, "unable to create ServiceAccount", "serviceAccount", sa.Name)
			return err
		}
		log.V(1).Info("ServiceAccount already exists", "serviceAccount", sa.Name)
	} else {
		log.V(1).Info("created ServiceAccount", "serviceAccount", sa.Name)
	}

	rb := r.constructRoleBinding(virtualMachineBMC)

	if err := ctrl.SetControllerReference(virtualMachineBMC, rb, r.Scheme); err != nil {
		return err
	}

	if err := r.Create(ctx, rb); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			log.Error(err, "unable to create RoleBinding", "roleBinding", rb.Name)
			return err
		}
		log.V(1).Info("RoleBinding already exists", "roleBinding", rb.Name)
	} else {
		log.V(1).Info("created RoleBinding", "roleBinding", rb.Name)
	}

	return nil
}

func (r *VirtualMachineBMCReconciler) constructPodFromVirtualMachineBMC(virtualMachineBMC *bmcv1.VirtualMachineBMC) *corev1.Pod {
	name := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Spec.VirtualMachineRef.Name)
	serviceAccountName := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Spec.VirtualMachineRef.Name)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				VirtualMachineBMCNameLabel: virtualMachineBMC.Name,
				VMNameLabel:                virtualMachineBMC.Spec.VirtualMachineRef.Name,
			},
			Name:      name,
			Namespace: virtualMachineBMC.Namespace,
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
						virtualMachineBMC.Namespace,
						virtualMachineBMC.Spec.VirtualMachineRef.Name,
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
			ServiceAccountName: serviceAccountName,
		},
	}

	return pod
}

func (r *VirtualMachineBMCReconciler) constructServiceFromVirtualMachineBMC(virtualMachineBMC *bmcv1.VirtualMachineBMC) *corev1.Service {
	name := fmt.Sprintf("%s-virtbmc", virtualMachineBMC.Spec.VirtualMachineRef.Name)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				VirtualMachineBMCNameLabel: virtualMachineBMC.Name,
				VMNameLabel:                virtualMachineBMC.Spec.VirtualMachineRef.Name,
			},
			Name:      name,
			Namespace: virtualMachineBMC.Namespace,
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

func (r *VirtualMachineBMCReconciler) validateVirtualMachineExists(ctx context.Context, virtualMachineBMC *bmcv1.VirtualMachineBMC) (bool, error) {
	log := log.FromContext(ctx)

	vmKey := client.ObjectKey{
		Namespace: virtualMachineBMC.Namespace,
		Name:      virtualMachineBMC.Spec.VirtualMachineRef.Name,
	}

	var vm kubevirtv1.VirtualMachine
	if err := r.Get(ctx, vmKey, &vm); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Referenced VirtualMachine not found",
				"vm", virtualMachineBMC.Spec.VirtualMachineRef.Name,
				"namespace", virtualMachineBMC.Namespace)

			if changed := meta.SetStatusCondition(&virtualMachineBMC.Status.Conditions, metav1.Condition{
				Type:   bmcv1.ConditionReady,
				Status: metav1.ConditionFalse,
				Reason: "VirtualMachineNotFound",
				Message: fmt.Sprintf("VirtualMachine %q not found in namespace %q",
					virtualMachineBMC.Spec.VirtualMachineRef.Name,
					virtualMachineBMC.Namespace),
			}); changed {
				return false, r.Status().Update(ctx, virtualMachineBMC)
			}
			return false, nil
		}

		log.Error(err, "error checking VirtualMachine existence")
		return false, err
	}

	return true, nil
}

//+kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachines,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachineinstances,verbs=get;list;watch;delete
//+kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods;services;serviceaccounts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

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

	var virtualMachineBMC bmcv1.VirtualMachineBMC
	if err := r.Get(ctx, req.NamespacedName, &virtualMachineBMC); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	exists, err := r.validateVirtualMachineExists(ctx, &virtualMachineBMC)
	if err != nil {
		return ctrl.Result{}, err
	}
	if !exists {
		return ctrl.Result{}, nil
	}

	if err := r.ensureRBACResources(ctx, &virtualMachineBMC); err != nil {
		return ctrl.Result{}, err
	}

	// Prepare the virtBMC Pod
	pod := r.constructPodFromVirtualMachineBMC(&virtualMachineBMC)
	if err := ctrl.SetControllerReference(&virtualMachineBMC, pod, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Create the virtBMC Pod on the cluster
	if err := r.Create(ctx, pod); err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create Pod for VirtualMachineBMC", "pod", pod)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created Pod for VirtualMachineBMC", "pod", pod)

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
		For(&bmcv1.VirtualMachineBMC{}).
		Owns(&corev1.Pod{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&rbacv1.RoleBinding{}).
		Watches(
			&kubevirtv1.VirtualMachine{},
			handler.EnqueueRequestsFromMapFunc(r.findVirtualMachineBMCsForVM),
		).
		Complete(r)
}

func (r *VirtualMachineBMCReconciler) findVirtualMachineBMCsForVM(ctx context.Context, vm client.Object) []reconcile.Request {
	log := log.FromContext(ctx)

	// List all VirtualMachineBMCs in the same namespace as the VM
	var vmBMCList bmcv1.VirtualMachineBMCList
	if err := r.List(ctx, &vmBMCList, client.InNamespace(vm.GetNamespace())); err != nil {
		log.Error(err, "unable to list VirtualMachineBMCs")
		return []reconcile.Request{}
	}

	var requests []reconcile.Request
	for _, vmBMC := range vmBMCList.Items {
		if vmBMC.Spec.VirtualMachineRef.Name == vm.GetName() {
			log.Info("Enqueueing VirtualMachineBMC due to VirtualMachine change",
				"vm", vm.GetName(),
				"vmBMC", vmBMC.Name)
			requests = append(requests, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      vmBMC.Name,
					Namespace: vmBMC.Namespace,
				},
			})
		}
	}

	return requests
}
