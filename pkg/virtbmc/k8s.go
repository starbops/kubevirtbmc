package virtbmc

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	cdiclient "kubevirt.io/client-go/containerizeddataimporter"
	kvclient "kubevirt.io/client-go/kubevirt"
)

func NewVirtClient(options Options) kvclient.Interface {
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
	virtClient, err := kvclient.NewForConfig(config)
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
