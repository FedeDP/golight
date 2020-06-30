package dimmer

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/backlight"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"github.com/FedeDP/golight/state"
)

var api, _ = clightd.NewIdleApi()
var oldPct float64

func Subscribe() chan *dbus.Signal {
	c := api.Subscribe()
	err := api.SetTimeout(conf.DimmerTO) // 30s
	if err != nil {
		panic(err)
	}
	err = api.Start()
	if err != nil {
		panic(err)
	}
	return c
}

func Update(v *dbus.Signal) {
	if api.Update(v) {
		fmt.Println("Entering dimmed state.")
		state.DisplaySet(state.DisplayDIM)
		oldPct = state.CurBl
		backlight.Set(0.10)
	} else {
		fmt.Println("Leaving dimmed state.")
		state.DisplayClear(state.DisplayDIM)
		backlight.Set(oldPct)
	}
}

func Close() {
	err := api.Destroy()
	if err != nil {
		fmt.Println(err)
	}
}
