package clightd

import (
	"github.com/godbus/dbus/v5"
)

/** DPMS API object **/
const (
	dpmsInterface         = "org.clightd.clightd.Dpms"
	dpmsObjectPath        = "/org/clightd/clightd/Dpms"

	dpmsMethodSet        = dpmsInterface + ".Set"
)

type DpmsApi interface {
	ClightdApi
	GoSetDpms(dpmsLvl int32, ch chan *dbus.Call) error
	SetDpms(dpmsLvl int32) (bool, error)
}

func NewDpmsApi() (DpmsApi, error) {
	err := ensureXorg()
	if err == nil {
		return initialize(dpmsObjectPath)
	}
	return nil, err
}

func (api api) GoSetDpms(dpmsLvl int32, ch chan *dbus.Call) error {
	call := api.obj.Go(dpmsMethodSet,0, ch, xdisplay, xauth, dpmsLvl)
	return call.Err
}

func (api api) SetDpms(dpmsLvl int32) (bool, error) {
	call := api.obj.Call(dpmsMethodSet,0, xdisplay, xauth, dpmsLvl)
	if call.Err != nil {
		return false, call.Err
	}
	return true, nil
}
