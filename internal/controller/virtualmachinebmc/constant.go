package virtualmachinebmc

const (
	DefaultUsername   = "admin"
	DefaultPassword   = "password"
	kbmcContainerName = "kbmc"
	kbmcImageName     = "starbops/kbmc"
	kbmcImageTag      = "dev"
	ipmiPort          = 623
	IPMISvcPort       = 623
	ipmiPortName      = "ipmi"
	ManagedByLabel    = "app.kubernetes.io/managed-by"
	KBMCNameLabel     = "kubevirt.org/virtualmachinebmc-name"
	VMNameLabel       = "kubevirt.org/vm-name"
	KBMCNamespace     = "virtualmachinebmc-system"
)
