package clightd

import (
	"github.com/godbus/dbus/v5"
)

type GammaSmooth struct {
	Smooth  bool
	Step    uint32
	Timeout uint32
}

/** Gamma API object **/
const (
	gammaInterface  = "org.clightd.clightd.Gamma"
	gammaObjectPath = "/org/clightd/clightd/Gamma"

	gammaMethodSet = gammaInterface + ".Set"
	gammaMethodGet = gammaInterface + ".Get"
)

type GammaApi interface {
	ClightdApi
	GoSetTemp(temp int32, smooth *GammaSmooth, ch chan *dbus.Call) error
	SetTemp(temp int32, smooth *GammaSmooth) (bool, error)

	GoGetTemp(ch chan *dbus.Call) error
	GetTemp() (int32, error)
}

func NewGammaApi() (GammaApi, error) {
	err := ensureXorg()
	if err == nil {
		return initialize(gammaObjectPath)
	}
	return nil, err
}

func (api api) GoSetTemp(temp int32, smooth *GammaSmooth, ch chan *dbus.Call) error {
	call := api.obj.Go(gammaMethodSet, 0, ch, xdisplay, xauth, temp, smooth)
	return call.Err
}

func (api api) SetTemp(temp int32, smooth *GammaSmooth) (ok bool, err error) {
	call := api.obj.Call(gammaMethodSet, 0, xdisplay, xauth, temp, smooth)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&ok)
	}
	return
}

func (api api) GoGetTemp(ch chan *dbus.Call) error {
	call := api.obj.Go(gammaMethodGet, 0, ch, xdisplay, xauth)
	return call.Err
}

func (api api) GetTemp() (gammaLvl int32, err error) {
	call := api.obj.Call(gammaMethodGet, 0, xdisplay, xauth)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&gammaLvl)
	}
	return
}
