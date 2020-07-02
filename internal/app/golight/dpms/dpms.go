package dpms

import (
	"fmt"
	"github.com/FedeDP/golight/internal/app/golight/conf"
	"github.com/FedeDP/golight/internal/app/golight/state"
	idler "github.com/FedeDP/golight/internal/pkg/idlelib"
	"github.com/FedeDP/golight/pkg/go-clightd"
	"github.com/godbus/dbus/v5"
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
	if idler.Update(api, v) {
		state.DisplaySet(state.DisplayOFF)
		fmt.Println("Entering DPMS state.")
		setDpms(1)
	} else {
		state.DisplayClear(state.DisplayOFF)
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
