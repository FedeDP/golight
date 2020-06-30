package clightd

import (
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

var(
	xauth						string = os.Getenv("XAUTHORITY")
	xdisplay					string = os.Getenv("DISPLAY")
)

type ClightdApi interface {
	Destroy() error
}

type ApiDtor func(api api) error

type api struct {
	conn 	*dbus.Conn
	obj  	dbus.BusObject
	dtor	ApiDtor
}

func (api api) Destroy() error {
	if api.dtor != nil {
		_ = api.dtor(api)
	}
	return api.conn.Close()
}

func init() {
	fmt.Printf("Environment:\n\tXDisplay -> %s\n\tXauth -> %s\n", xdisplay, xauth)
}

func initialize(path string) (api, error) {
	var api api
	var err error
	api.dtor = nil
	api.conn, err = dbus.ConnectSystemBus()
	if err != nil {
		return api, err
	}
	api.obj = api.conn.Object(clightdInterface, dbus.ObjectPath(path))
	return api, nil
}

func ensureXorg() error {
	if len(xauth) == 0 || len(xdisplay) == 0 {
		return errors.New("Only supported on X.")
	}
	return nil
}
