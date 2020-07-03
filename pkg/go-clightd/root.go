package clightd

/** Main API object **/
const (
	clightdInterface  = "org.clightd.clightd"
	clightdObjectPath = "/org/clightd/clightd"

	clightdPropVersion = clightdInterface + ".Version"
)

type Root interface {
	Version() (string, error)
}

func NewRoot() (Root, error) {
	return initialize(clightdObjectPath)
}

func (api api) Version() (version string, err error) {
	prop, err := api.obj.GetProperty(clightdPropVersion)
	if err == nil {
		err = prop.Store(version)
	}
	return
}
