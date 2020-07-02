package dimmer

import (
	"fmt"
	"github.com/FedeDP/golight/internal/app/golight/backlight"
	"github.com/FedeDP/golight/internal/app/golight/conf"
	"github.com/FedeDP/golight/internal/app/golight/state"
	idler "github.com/FedeDP/golight/internal/pkg/idlelib"
	"github.com/FedeDP/golight/pkg/go-clightd"
	"github.com/godbus/dbus/v5"
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
	if idler.Update(api, v) {
		state.DisplaySet(state.DisplayDIM)
		fmt.Println("Entering dimmed state.")
		oldPct = state.CurBl
		backlight.Set(0.10)
	} else {
		state.DisplayClear(state.DisplayDIM)
		fmt.Println("Leaving dimmed state.")
		backlight.Set(oldPct)
	}
}

func Close() {
	idler.Close(api)
}
