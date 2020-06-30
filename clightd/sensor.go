package clightd

import (
	"errors"
	"github.com/godbus/dbus/v5"
)

/** Sensor API object **/
const (
	sensInterface         = "org.clightd.clightd.Sensor"
	sensObjectPath        = "/org/clightd/clightd/Sensor"

	sensMethodCapture        = sensInterface + ".Capture"
)

type SensType string
const(
	SensAny SensType 	= ""
	SensCamera 			= "/Camera"
	SensAls 			= "/Als"
	SensCustom 			= "/Custom"
)

type SensorApi interface {
	ClightdApi
	GoCapture(blC chan *dbus.Call, ncaptures int32) error
	Capture(ncaptures int32) ([]float64, error)
}

func NewSensorApi(sensType SensType) (SensorApi, error) {
	switch sensType {
	case SensAny, SensCamera, SensAls, SensCustom:
		path := sensObjectPath + sensType
		return initialize(string(path))
	}
	return nil, errors.New("Wrong Sensor type.")
}

func (api api) GoCapture(blC chan *dbus.Call, ncaptures int32) error {
	call := api.obj.Go(sensMethodCapture,0, blC, "", ncaptures, "")
	return call.Err
}

func (api api) Capture(ncaptures int32) ([]float64, error) {
	call := api.obj.Call(sensMethodCapture,0, "", ncaptures, "")
	if call.Err != nil {
		return nil, call.Err
	}
	var Sensor string
	var Val []float64
	err := call.Store(&Sensor, &Val)
	if err != nil {
		return nil, err
	}
	return Val, err
}
