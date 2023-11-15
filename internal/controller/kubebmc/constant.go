package kubebmc

const (
	DefaultUsername   = "admin"
	DefaultPassword   = "password"
	kbmcContainerName = "kbmc"
	kbmcImageName     = "starbops/kbmc"
	kbmcImageTag      = "dev"
	ipmiPort          = 623
	ipmiSvcPort       = 623
	ipmiPortName      = "ipmi"
	ManagedByLabel    = "app.kubernetes.io/managed-by"
	KBMCNameLabel     = "zespre.com/kubebmc-name"
	VMNameLabel       = "zespre.com/vm-name"
	KBMCNamespace     = "kubebmc-system"
)
