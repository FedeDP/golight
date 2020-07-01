package dimmer

import (
	"fmt"
	"github.com/FedeDP/golight/idler"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/backlight"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"github.com/FedeDP/golight/state"
)

var api, _ = clightd.NewIdleClientApi()
var oldPct float64

func Subscribe() chan *dbus.Signal {
	return idler.Subscribe(api, conf.DimmerTO[state.Ac])
}

func UpdateTimer() {
	_ = api.SetTimeout(conf.DimmerTO[state.Ac])
}

func Update(v *dbus.Signal) {
	if idler.Update(v, state.DisplayDIM) {
		fmt.Println("Entering dimmed state.")
		oldPct = state.CurBl
		backlight.Set(0.10)
	} else {
		fmt.Println("Leaving dimmed state.")
		backlight.Set(oldPct)
	}
}

func Close() {
	idler.Close(api)
}
