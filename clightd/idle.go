package clightd

import (
	"github.com/godbus/dbus/v5"
)

/** Idle API object **/
const (
	idleInterface         	= "org.clightd.clightd.Idle"
	idleObjectPath        	= "/org/clightd/clightd/Idle"
	idleClientInterface	  	= idleInterface + ".Client"

	idleMethodGetClient     = idleInterface + ".GetClient"
	idleMethodDestroyClient = idleInterface + ".DestroyClient"

	idleMethodStartClient	= idleClientInterface + ".Start"
	idleMethodStopClient	= idleClientInterface + ".Stop"
	idlePropClientTimeout	= idleClientInterface + ".Timeout"
)

type IdleApi interface {
	ClightdApi
	Subscribe() chan *dbus.Signal
	Update(v *dbus.Signal) bool
	SetTimeout(timeout uint) error
	Start() error
	Stop() error
	Reset() error
}

func NewIdleApi() (IdleApi, error) {
	var cl api
	cl.dtor = dtor
	cl.conn, _ = dbus.ConnectSystemBus()
	sysobj := cl.conn.Object(clightdInterface, idleObjectPath)
	call := sysobj.Call(idleMethodGetClient, 0)
	if call.Err != nil {
		panic(call.Err)
	}
	var clientPath dbus.ObjectPath
	err := call.Store(&clientPath)
	if err != nil {
		panic(err)
	}
	cl.obj = cl.conn.Object(clightdInterface, clientPath)
	return cl, nil
}

func (api api) Subscribe() chan *dbus.Signal {
	api.obj.AddMatchSignal(idleInterface, "Idle", dbus.WithMatchObjectPath(api.obj.Path()))
	c := make(chan *dbus.Signal, 10)
	api.conn.Signal(c)
	return c
}

func (api api) Update(v *dbus.Signal) bool {
	return v.Body[0].(bool)
}

func (api api) SetTimeout(timeout uint) error {
	if timeout > 0 {
		return api.obj.SetProperty(idlePropClientTimeout, dbus.MakeVariant(timeout))
	}
	return api.Stop()
}

func (api api) Start() error {
	call := api.obj.Call(idleMethodStartClient, 0)
	return call.Err
}

func (api api) Stop() error {
	call := api.obj.Call(idleMethodStopClient, 0)
	return call.Err
}

func (api api) Reset() error {
	err := api.Stop()
	if err != nil {
		return err
	}
	return api.Start()
}

func dtor(api api) error {
	_ = api.Stop()
	sysobj := api.conn.Object(clightdInterface, idleObjectPath)
	call := sysobj.Call(idleMethodDestroyClient, 0, api.obj.Path())
	return call.Err
}
