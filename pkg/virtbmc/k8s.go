package virtbmc

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	kubevirtv1type "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
)

func NewK8sClient(options Options) *kubevirtv1type.KubevirtV1Client {
	var (
		config *rest.Config
		err    error
	)

	// creates the in-cluster config
	config, err = rest.InClusterConfig()
	if err != nil {
		if err == rest.ErrNotInCluster {
			goto localConfig
		}
		panic(err.Error())
	}
	goto genClientset

localConfig:
	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	config, err = clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
	if err != nil {
		panic(err.Error())
	}

genClientset:
	clientset, err := kubevirtv1type.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
