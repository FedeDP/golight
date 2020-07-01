package dpms

import (
	"fmt"
	"github.com/FedeDP/golight/idler"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"github.com/FedeDP/golight/state"
)

var api, _ = clightd.NewIdleClientApi()
var dpmsApi, _ = clightd.NewDpmsApi()

func Subscribe() chan *dbus.Signal {
	return idler.Subscribe(api, conf.DpmsTO[state.Ac])
}

func UpdateTimer() {
	_ = api.SetTimeout(conf.DimmerTO[state.Ac])
}

func Update(v *dbus.Signal) {
	if idler.Update(v, state.DisplayOFF) {
		fmt.Println("Entering DPMS state.")
		setDpms(1)
	} else {
		fmt.Println("Leaving DPMS state.")
		setDpms(0)
	}
}

func Close() {
	idler.Close(api)
	_ = dpmsApi.Destroy()
}

func setDpms(val int32) {
	_ = dpmsApi.GoSetDpms(val, nil)
}
