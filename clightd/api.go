package clightd

/** Main API object **/
const (
	clightdInterface         = "org.clightd.clightd"
	clightdObjectPath        = "/org/clightd/clightd"

	clightdPropVersion        = clightdInterface + ".Version"
)

type Api interface {
	Version() (string, error)
}

func NewApi() (Api, error) {
	return initialize(clightdObjectPath)
}

func (api api) Version() (string, error) {
	prop, err := api.obj.GetProperty(clightdPropVersion)
	if err != nil {
		return "", err
	}
	var version string
	err = prop.Store(version)
	return version, err
}