package bmc

const (
	// Finalizer constants
	FinalizerName = "bmc.kubevirt.io/finalizer"

	// Label and name suffix constants
	AppLabelKey          = "app"
	BMCProxyLabelSuffix  = "-bmc-proxy"
	BMCServiceNameSuffix = "-bmc-service"

	// Container and port constants
	BMCContainerName = "bmc-proxy"
	BMCPortName      = "bmc"
	BMCPortNumber    = 623

	// Auth secret mount constants
	AuthSecretVolumeName = "auth-secret"
	AuthSecretMountPath  = "/etc/bmc-auth"

	// Condition and message constants
	ConditionReady    = "Ready"
	ReasonBMCReady    = "BMCReady"
	ReasonBMCNotReady = "BMCNotReady"

	MsgBMCReady    = "BMC is ready and operational."
	MsgBMCNotReady = "BMC is not ready yet, waiting for deployment to stabilize. " +
		"Ensure the BMC proxy is running and has the correct configuration. " +
		"Check logs for more details."
)
