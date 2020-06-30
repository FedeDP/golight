package dpms

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"github.com/FedeDP/golight/state"
)

var api, _ = clightd.NewIdleApi()
var dpmsApi, _ = clightd.NewDpmsApi()

func Subscribe() chan *dbus.Signal {
	c := api.Subscribe()
	err := api.SetTimeout(conf.DpmsTO)
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
		fmt.Println("Entering DPMS state.")
		state.DisplaySet(state.DisplayOFF)
		setDpms(1)
	} else {
		fmt.Println("Leaving DPMS state.")
		state.DisplayClear(state.DisplayOFF)
		setDpms(0)
	}
}

func Close() {
	err := api.Destroy()
	if err != nil {
		fmt.Println(err)
	}
	_ = dpmsApi.Destroy()
}

func setDpms(val int32) {
	_ = dpmsApi.GoSetDpms(val, nil)
}
