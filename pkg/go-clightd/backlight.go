package clightd

import (
	"github.com/godbus/dbus/v5"
)

type BacklightSmooth struct {
	Smooth  bool
	Step    float64
	Timeout uint32
}

/** Backlight API object **/
const (
	backlightInterface  = "org.clightd.clightd.Backlight"
	backlightObjectPath = "/org/clightd/clightd/Backlight"

	backlightMethodSetAll   = backlightInterface + ".SetAll"
	backlightMethodRaiseAll = backlightInterface + ".RaiseAll"
	backlightMethodLowerAll = backlightInterface + ".LowerAll"
	backlightMethodGetAll   = backlightInterface + ".GetAll"

	backlightMethodSet   = backlightInterface + ".Set"
	backlightMethodRaise = backlightInterface + ".Raise"
	backlightMethodLower = backlightInterface + ".Lower"
	backlightMethodGet   = backlightInterface + ".Get"
)

type BacklightStruct struct {
	serialNumber string
	level        float64
}

type BacklightApi interface {
	ClightdApi

	GoSetAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error
	SetAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (bool, error)
	GoRaiseAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error
	RaiseAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (bool, error)
	GoLowerAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error
	LowerAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (bool, error)
	GoGetAll(internalBlInterface string, ch chan *dbus.Call) error
	GetAll(internalBlInterface string) ([]BacklightStruct, error)

	GoSet(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error
	Set(val float64, smooth *BacklightSmooth, serial string) (bool, error)
	GoRaise(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error
	Raise(val float64, smooth *BacklightSmooth, serial string) (bool, error)
	GoLower(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error
	Lower(val float64, smooth *BacklightSmooth, serial string) (bool, error)
}

func NewBacklightApi() (BacklightApi, error) {
	return initialize(backlightObjectPath)
}

func goSet(api api, method string, val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error {
	call := api.obj.Go(method, 0, ch, val, smooth, internalBlInterface)
	return call.Err
}

func set(api api, method string, val float64, smooth *BacklightSmooth, internalBlInterface string) (ok bool, err error) {
	call := api.obj.Call(backlightMethodSetAll, 0, val, smooth, internalBlInterface)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&ok)
	}
	return
}

func (api api) GoSetAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodSetAll, val, smooth, internalBlInterface, ch)
}

func (api api) SetAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (ok bool, err error) {
	return set(api, backlightMethodSetAll, val, smooth, internalBlInterface)
}

func (api api) GoRaiseAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodRaiseAll, val, smooth, internalBlInterface, ch)
}

func (api api) RaiseAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (ok bool, err error) {
	return set(api, backlightMethodRaiseAll, val, smooth, internalBlInterface)
}

func (api api) GoLowerAll(val float64, smooth *BacklightSmooth, internalBlInterface string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodRaiseAll, val, smooth, internalBlInterface, ch)
}

func (api api) LowerAll(val float64, smooth *BacklightSmooth, internalBlInterface string) (ok bool, err error) {
	return set(api, backlightMethodRaiseAll, val, smooth, internalBlInterface)
}

func (api api) GoGetAll(internalBlInterface string, ch chan *dbus.Call) error {
	call := api.obj.Go(backlightMethodGetAll, 0, ch, internalBlInterface)
	return call.Err
}

func (api api) GetAll(internalBlInterface string) (values []BacklightStruct, err error) {
	call := api.obj.Call(backlightMethodGetAll, 0, internalBlInterface)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&values)
	}
	return
}

func (api api) GoSet(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodSet, val, smooth, serial, ch)
}

func (api api) Set(val float64, smooth *BacklightSmooth, serial string) (bool, error) {
	return set(api, backlightMethodSet, val, smooth, serial)
}

func (api api) GoRaise(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodRaise, val, smooth, serial, ch)
}

func (api api) Raise(val float64, smooth *BacklightSmooth, serial string) (bool, error) {
	return set(api, backlightMethodRaise, val, smooth, serial)
}

func (api api) GoLower(val float64, smooth *BacklightSmooth, serial string, ch chan *dbus.Call) error {
	return goSet(api, backlightMethodLower, val, smooth, serial, ch)
}

func (api api) Lower(val float64, smooth *BacklightSmooth, serial string) (bool, error) {
	return set(api, backlightMethodLower, val, smooth, serial)
}

func (api api) GoGet(serial string, ch chan *dbus.Call) error {
	call := api.obj.Go(backlightMethodGet, 0, ch, serial)
	return call.Err
}

func (api api) Get(serial string) (value *BacklightStruct, err error) {
	call := api.obj.Call(backlightMethodGet, 0, serial)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(value)
	}
	return
}
