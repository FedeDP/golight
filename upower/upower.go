package upower

import (
	"github.com/FedeDP/golight/state"
	"fmt"
	"github.com/godbus/dbus/v5"
)

var conn, _ = dbus.ConnectSystemBus()
var upobj = conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

func Subscribe() chan *dbus.Signal {
	_ = conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/org/freedesktop/UPower"),
		dbus.WithMatchInterface("org.freedesktop.DBus.Properties"),
		dbus.WithMatchSender("org.freedesktop.UPower"),
		dbus.WithMatchMember("PropertiesChanged"))

	Update()
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	return c
}

func Update() {
	val, err := upobj.GetProperty("org.freedesktop.UPower.OnBattery")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Upower unavailable. Fallbacking to AC state.")
	} else {
		err = val.Store(&state.OnBatt)
		if state.OnBatt {
			fmt.Println("Current AC state: on Batt.")
		} else {
			fmt.Println("Current AC state: on AC.")
		}
	}
}

func Close() {
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
