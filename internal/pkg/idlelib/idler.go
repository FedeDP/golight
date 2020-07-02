package idler

import (
	"fmt"
	"github.com/FedeDP/golight/pkg/go-clightd"
	"github.com/godbus/dbus/v5"
)

func Subscribe(api clightd.IdleClientApi, timeout uint) (c chan *dbus.Signal) {
	c = make(chan *dbus.Signal, 10)
	api.SubscribeIdle(c)

	if err := api.SetTimeout(timeout); err != nil {
		panic(err)
	}

	if err := api.Start(); err != nil {
		panic(err)
	}
	return
}

func Update(api clightd.IdleClientApi, v *dbus.Signal) bool {
	return api.ParseIdle(v)
}

func Close(api clightd.IdleClientApi) {
	if err := api.Destroy(); err != nil {
		fmt.Println(err)
	}
}
