package location

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/maltegrosse/go-geoclue2"
	"github.com/FedeDP/golight/state"
)

var gcm geoclue2.GeoclueManager
var gcl geoclue2.GeoclueClient

func Subscribe() <-chan *dbus.Signal {
	var err error
	gcm, err = geoclue2.NewGeoclueManager()
	if err != nil {
		panic(err)
	}
	gcl, err = gcm.GetClient()
	if err != nil {
		panic(err)
	}

	_ = gcl.SetDesktopId("GoLight")
	_ = gcl.SetRequestedAccuracyLevel(geoclue2.GClueAccuracyLevelCity)
	_ = gcl.SetDistanceThreshold(50000) // 50km
	_ = gcl.Start()

	return gcl.SubscribeLocationUpdated()
}

func Update(v *dbus.Signal) {
	_, state.Location, _ = gcl.ParseLocationUpdated(v)
	logLoc()
}

func Close() {
	_ = gcm.DeleteClient(gcl)
}

func logLoc() {
	lat, _ := state.Location.GetLatitude()
	lon, _ := state.Location.GetLongitude()
	fmt.Printf("New location received: %.2f:%.2f.\n", lat, lon)
}
