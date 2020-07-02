package capture

import (
	"fmt"
	"github.com/FedeDP/golight/internal/app/golight/conf"
	"github.com/FedeDP/golight/internal/app/golight/state"
	"github.com/FedeDP/golight/pkg/go-clightd"
	"github.com/godbus/dbus/v5"
	"time"
)

var api, _ = clightd.NewSensorApi(clightd.SensAny)

func Subscribe() <- chan time.Time {
	m := time.Duration(conf.CaptureTO[state.Ac]) * time.Second
	return time.Tick(m)
}

func Update(blC chan *dbus.Call) {
	if err := api.GoCapture(blC, "", conf.NCaptures[state.Ac], ""); err != nil {
		fmt.Println(err)
	}
}

func Close() {
	if err := api.Destroy(); err != nil {
		fmt.Println(err)
	}
}
