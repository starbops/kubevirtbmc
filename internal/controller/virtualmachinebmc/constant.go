package virtualmachinebmc

const (
	DefaultUsername            = "admin"
	DefaultPassword            = "password"
	virtBMCContainerName       = "virtbmc"
	VirtBMCImageName           = "starbops/virtbmc"
	ipmiPort                   = 10623
	IPMISvcPort                = 623
	ipmiPortName               = "ipmi"
	VirtualMachineBMCNameLabel = "kubevirt.io/virtualmachinebmc-name"
	VMNameLabel                = "kubevirt.io/vm-name"
	VirtualMachineBMCNamespace = "kubevirtbmc-system"
)
