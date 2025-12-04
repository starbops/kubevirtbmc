package virtbmc

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"kubevirt.io/client-go/kubecli"
	cdiclient "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned"
)

// func NewK8sClient(options Options) *kubevirtv1type.KubevirtV1Client {
// 	var (
// 		config *rest.Config
// 		err    error
// 	)

// 	// creates the in-cluster config
// 	config, err = rest.InClusterConfig()
// 	if err != nil {
// 		if err == rest.ErrNotInCluster {
// 			goto localConfig
// 		}
// 		panic(err.Error())
// 	}
// 	goto genClientset

// localConfig:
// 	// uses the current context in kubeconfig
// 	// path-to-kubeconfig -- for example, /root/.kube/config
// 	config, err = clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// genClientset:
// 	clientset, err := kubevirtv1type.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return clientset
// }

func NewVirtClient(options Options) kubecli.KubevirtClient {
	// Build config
	config, err := clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
	if err != nil {
		// Fallback to in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	}

	// KubeVirt client
	virtClient, err := kubecli.GetKubevirtClientFromRESTConfig(config)
	if err != nil {
		panic(err)
	}

	return virtClient
}

func NewCdiClient(options Options) *cdiclient.Clientset {
	// Build config
	config, err := clientcmd.BuildConfigFromFlags("", options.KubeconfigPath)
	if err != nil {
		// Fallback to in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	}

	// CDI client
	cdiClient, err := cdiclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return cdiClient
}
