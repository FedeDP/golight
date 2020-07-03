package clightd

import (
	"errors"
	"github.com/godbus/dbus/v5"
)

/** Sensor API object **/
const (
	sensInterface  = "org.clightd.clightd.Sensor"
	sensObjectPath = "/org/clightd/clightd/Sensor"

	sensMethodCapture   = sensInterface + ".Capture"
	sensMethodAvailable = sensInterface + ".IsAvailable"
)

type SensType string

const (
	SensAny    SensType = ""
	SensCamera          = "/Camera"
	SensAls             = "/Als"
	SensCustom          = "/Custom"
)

type SensorApi interface {
	ClightdApi

	GoCapture(blC chan *dbus.Call, sens string, ncaptures int32, sensOpts string) error
	Capture(sens string, ncaptures int32, sensOpts string) (string, []float64, error)
	GoIsAvailable(blC chan *dbus.Call, sens string) error
	IsAvailable(sens string) (string, bool, error)

	SubscribeChanged(c chan *dbus.Signal)
	ParseChanged(v *dbus.Signal) (string, string)
}

func NewSensorApi(sensType SensType) (SensorApi, error) {
	switch sensType {
	case SensAny, SensCamera, SensAls, SensCustom:
		path := sensObjectPath + sensType
		return initialize(string(path))
	}
	return nil, errors.New("Wrong Sensor type.")
}

func (api api) GoCapture(blC chan *dbus.Call, sens string, ncaptures int32, sensOpts string) error {
	call := api.obj.Go(sensMethodCapture, 0, blC, sens, ncaptures, sensOpts)
	return call.Err
}

func (api api) Capture(sens string, ncaptures int32, sensOpts string) (foundSens string, values []float64, err error) {
	call := api.obj.Call(sensMethodCapture, 0, sens, ncaptures, sensOpts)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&foundSens, &values)
	}
	return
}

func (api api) GoIsAvailable(blC chan *dbus.Call, sens string) error {
	call := api.obj.Go(sensMethodAvailable, 0, blC, sens)
	return call.Err
}

func (api api) IsAvailable(sens string) (foundSens string, ok bool, err error) {
	call := api.obj.Call(sensMethodAvailable, 0, sens)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&foundSens, &ok)
	}
	return
}

func (api api) SubscribeChanged(c chan *dbus.Signal) {
	api.obj.AddMatchSignal(sensInterface, "Changed", dbus.WithMatchObjectPath(api.obj.Path()))
	api.conn.Signal(c)
}

func (api api) ParseChanged(v *dbus.Signal) (sens string, action string) {
	sens = v.Body[0].(string)
	action = v.Body[1].(string)
	return
}
