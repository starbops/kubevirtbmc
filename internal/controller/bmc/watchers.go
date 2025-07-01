package bmc

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func mapVMToBMCs(c client.Client) handler.MapFunc {
	return func(ctx context.Context, obj client.Object) []reconcile.Request {
		vm, ok := obj.(*kubevirtv1.VirtualMachine)
		if !ok {
			return nil
		}

		var bmcList bmcv1beta1.VirtualMachineBMCList
		if err := c.List(ctx, &bmcList, client.InNamespace(vm.Namespace)); err != nil {
			return nil
		}

		var reqs []reconcile.Request
		for _, bmc := range bmcList.Items {
			if bmc.Spec.VirtualMachineRef != nil && bmc.Spec.VirtualMachineRef.Name == vm.Name {
				reqs = append(reqs, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: bmc.Namespace,
						Name:      bmc.Name,
					},
				})
			}
		}
		return reqs
	}
}

func mapSecretToBMCs(c client.Client) handler.MapFunc {
	return func(ctx context.Context, obj client.Object) []reconcile.Request {
		secret, ok := obj.(*corev1.Secret)
		if !ok {
			return nil
		}

		var bmcList bmcv1beta1.VirtualMachineBMCList
		if err := c.List(ctx, &bmcList, client.InNamespace(secret.Namespace)); err != nil {
			return nil
		}

		var reqs []reconcile.Request
		for _, bmc := range bmcList.Items {
			if bmc.Spec.AuthSecretRef != nil && bmc.Spec.AuthSecretRef.Name == secret.Name {
				reqs = append(reqs, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: bmc.Namespace,
						Name:      bmc.Name,
					},
				})
			}
		}
		return reqs
	}
}
