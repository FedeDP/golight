package capture

import (
	"fmt"
	"github.com/FedeDP/golight/state"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"time"
)

var api, _ = clightd.NewSensorApi(clightd.SensAny)

func Subscribe() <- chan time.Time {
	m := time.Duration(conf.CaptureTO[state.Ac]) * time.Second
	return time.Tick(m)
}

func Update(blC chan *dbus.Call) {
	if err := api.GoCapture(blC, conf.NCaptures[state.Ac]); err != nil {
		fmt.Println(err)
	}
}

func Close() {
	if err := api.Destroy(); err != nil {
		fmt.Println(err)
	}
}
