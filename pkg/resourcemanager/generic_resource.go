package resourcemanager

type ODataInterface interface {
	GetODataID() string

	Manage(ODataInterface) error
	ManagedBy(ODataInterface) error
}
