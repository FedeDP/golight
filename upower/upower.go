package upower

import (
	"fmt"
	"github.com/FedeDP/golight/state"
	"github.com/godbus/dbus/v5"
)

var conn, _ = dbus.ConnectSystemBus()
var upobj = conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

func Subscribe() (c chan *dbus.Signal) {
	err := conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/org/freedesktop/UPower"),
		dbus.WithMatchInterface("org.freedesktop.DBus.Properties"),
		dbus.WithMatchSender("org.freedesktop.UPower"),
		dbus.WithMatchMember("PropertiesChanged"))
	if err == nil {
		_, err = Update()
		if err == nil {
			if state.Ac == state.OnBatt {
				fmt.Println("Initial AC state: on Batt.")
			} else {
				fmt.Println("Initial AC state: on AC.")
			}
			c := make(chan *dbus.Signal, 10)
			conn.Signal(c)
		}
	}
	return
}

func Update() (bool, error) {
	val, err := upobj.GetProperty("org.freedesktop.UPower.OnBattery")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Upower unavailable. Fallbacking to AC state.")
		return false, err
	}

	var onBatt bool
	err = val.Store(&onBatt)
	/* Same value */
	if err != nil || onBatt == (state.Ac == state.OnBatt) {
		return false, err
	}
	return true, err
}

func Close() {
	if err := conn.Close(); err != nil {
		fmt.Println(err)
	}
}
