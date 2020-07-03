package clightd

import (
	"github.com/godbus/dbus/v5"
)

/** DPMS API object **/
const (
	dpmsInterface  = "org.clightd.clightd.Dpms"
	dpmsObjectPath = "/org/clightd/clightd/Dpms"

	dpmsMethodSet = dpmsInterface + ".Set"
	dpmsMethodGet = dpmsInterface + ".Get"
)

type DpmsApi interface {
	ClightdApi
	GoSetDpms(dpmsLvl int32, ch chan *dbus.Call) error
	SetDpms(dpmsLvl int32) (bool, error)
	GoGetDpms(ch chan *dbus.Call) error
	GetDpms() (int32, error)
}

func NewDpmsApi() (DpmsApi, error) {
	err := ensureXorg()
	if err == nil {
		return initialize(dpmsObjectPath)
	}
	return nil, err
}

func (api api) GoSetDpms(dpmsLvl int32, ch chan *dbus.Call) error {
	call := api.obj.Go(dpmsMethodSet, 0, ch, xdisplay, xauth, dpmsLvl)
	return call.Err
}

func (api api) SetDpms(dpmsLvl int32) (bool, error) {
	call := api.obj.Call(dpmsMethodSet, 0, xdisplay, xauth, dpmsLvl)
	if call.Err != nil {
		return false, call.Err
	}
	return true, nil
}

func (api api) GoGetDpms(ch chan *dbus.Call) error {
	call := api.obj.Go(dpmsMethodGet, 0, ch, xdisplay, xauth)
	return call.Err
}

func (api api) GetDpms() (dpmsLvl int32, err error) {
	call := api.obj.Call(dpmsMethodGet, 0, xdisplay, xauth)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&dpmsLvl)
	}
	return
}
