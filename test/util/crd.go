package util

import (
	"context"

	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateOrUpdateCRD(k8sClient apiextcs.Interface, relativePath string) error {
	crd, err := parseCRDYaml(relativePath)
	if err != nil {
		return err
	}

	if _, err := k8sClient.ApiextensionsV1().CustomResourceDefinitions().Get(
		context.TODO(),
		crd.Name,
		metav1.GetOptions{},
	); err == nil {
		// CRD already exists, update it
		_, err = k8sClient.ApiextensionsV1().CustomResourceDefinitions().Update(context.TODO(), crd, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	} else if apierrors.IsNotFound(err) {
		// CRD does not exist, create it
		_, err = k8sClient.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func DeleteCRD(k8sClient apiextcs.Interface, relativePath string) error {
	crd, err := parseCRDYaml(relativePath)
	if err != nil {
		return err
	}

	return k8sClient.ApiextensionsV1().CustomResourceDefinitions().Delete(context.TODO(), crd.Name, metav1.DeleteOptions{})
}
