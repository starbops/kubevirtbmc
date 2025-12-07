package resourcemanager

type OdataInterface interface {
	OdataId() string

	Manage(OdataInterface) error
	ManagedBy(OdataInterface) error
}
