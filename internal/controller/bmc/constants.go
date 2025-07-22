package bmc

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	BMCFinalizer = "bmc_kubevirt_io_finalizer"

	ConditionReady = "Ready"
	AppLabelKey    = "kubevirtbmc"
	ReasonReady    = "AllComponentsReady"
	MessageReady   = "VirtualMachineBMC_is_ready"

	DefaultReplicas = 1

	// Internal container ports
	IpmiContainerPort    = 10623
	RedfishContainerPort = 10080

	// Service-facing ports
	ipmiServicePort    = 623
	redfishServicePort = 80

	ipmiPortName    = "ipmi"
	redfishPortName = "redfish"

	// Service account name for the BMC pod
	ServiceAccountName      = "kubevirtbmc-virtbmc"
	BMCProxyLabelSuffix     = "-bmc-proxy"
	BMCServiceNameSuffix    = "-bmc-service"
	BMCDeploymentNameSuffix = "-bmc-deployment"
)

// LabelsForBMC returns the labels for a VirtualMachineBMC resource.
// Example: kubevirtbmc: test-bmc-proxy
func LabelsForBMC(bmcName string) map[string]string {
	return map[string]string{
		AppLabelKey: bmcName + BMCProxyLabelSuffix,
	}
}

func MetaForBMC(bmcName, namespace, suffix string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      bmcName + suffix,
		Namespace: namespace,
	}
}
