package clightd

import (
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/conf"
)

/** Gamma API object **/
const (
	gammaInterface         = "org.clightd.clightd.Gamma"
	gammaObjectPath        = "/org/clightd/clightd/Gamma"

	gammaMethodSet        = gammaInterface + ".Set"
)

type GammaApi interface {
	ClightdApi
	GoSetTemp(temp int32, smooth *conf.GammaSmooth, ch chan *dbus.Call) error
	SetTemp(temp int32, smooth *conf.GammaSmooth) (bool, error)
}

func NewGammaApi() (GammaApi, error) {
	err := ensureXorg()
	if err == nil {
		return initialize(gammaObjectPath)
	}
	return nil, err
}

func (api api) GoSetTemp(temp int32, smooth *conf.GammaSmooth, ch chan *dbus.Call) error {
	call := api.obj.Go(gammaMethodSet,0, ch, xdisplay, xauth, temp, smooth)
	return call.Err
}

func (api api) SetTemp(temp int32, smooth *conf.GammaSmooth) (bool, error) {
	call := api.obj.Call(gammaMethodSet,0, xdisplay, xauth, temp, smooth)
	if call.Err != nil {
		return false, call.Err
	}
	return true, nil
}
