package bmc

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
	ServiceAccountName   = "kubevirtbmc-virtbmc"
	BMCProxyLabelSuffix  = "-bmc-proxy"
	BMCServiceNameSuffix = "-bmc-service"
)
