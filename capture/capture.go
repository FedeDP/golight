package capture

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/golight/clightd"
	"github.com/FedeDP/golight/conf"
	"time"
)

var api, _ = clightd.NewSensorApi(clightd.SensAny)

func Subscribe() <- chan time.Time {
	m := time.Duration(conf.CaptureTO) * time.Second
	return time.Tick(m)
}

func Update(blC chan *dbus.Call) {
	err := api.GoCapture(blC, conf.NCaptures)
	if err != nil {
		fmt.Println(err)
	}
}

func Close() {
	err := api.Destroy()
	if err != nil {
		fmt.Println(err)
	}
}
