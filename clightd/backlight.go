package clightd

import (
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/conf"
)

/** Backlight API object **/
const (
	backlightInterface         = "org.clightd.clightd.Backlight"
	backlightObjectPath        = "/org/clightd/clightd/Backlight"

	backlightMethodSetAll      = backlightInterface + ".SetAll"
)

type BacklightApi interface {
	ClightdApi
	GoSetAll(val float64, smooth *conf.BacklightSmooth, ch chan *dbus.Call) error
	SetAll(val float64, smooth *conf.BacklightSmooth) (bool, error)
}

func NewBacklightApi() (BacklightApi, error) {
	return initialize(backlightObjectPath)
}

func (api api) GoSetAll(val float64, smooth *conf.BacklightSmooth, ch chan *dbus.Call) error {
	call := api.obj.Go(backlightMethodSetAll, 0, ch, val, smooth,"")
	return call.Err
}

func (api api) SetAll(val float64, smooth *conf.BacklightSmooth) (bool, error) {
	call := api.obj.Call(backlightMethodSetAll, 0, val, smooth,"")
	if call.Err != nil {
		return false, call.Err
	}
	return true, nil
}
