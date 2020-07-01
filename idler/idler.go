package idler

import (
	"fmt"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/state"
	"github.com/godbus/dbus/v5"
)

func Subscribe(api clightd.IdleClientApi, timeout uint) (c chan *dbus.Signal) {
	c = make(chan *dbus.Signal, 10)
	api.Subscribe(c)

	if err := api.SetTimeout(timeout); err != nil {
		panic(err)
	}

	if err := api.Start(); err != nil {
		panic(err)
	}
	return
}

func Update(v *dbus.Signal, bit state.DisplayState) bool {
	if v.Body[0].(bool) {
		state.DisplaySet(bit)
		return true
	}
	state.DisplayClear(bit)
	return false
}

func Close(api clightd.IdleClientApi) {
	if err := api.Destroy(); err != nil {
		fmt.Println(err)
	}
}