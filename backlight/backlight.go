package backlight

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/FedeDP/go-clightd"
	"github.com/FedeDP/golight/conf"
	"github.com/FedeDP/golight/state"
	"gonum.org/v1/gonum/stat"
	"math"
)

var api, _ = clightd.NewBacklightApi()
var BlFitParams [2]float64

func init() {
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	BlFitParams[1], BlFitParams[0] = stat.LinearRegression(x, conf.BlRegPoints, nil, false)
}

func Subscribe() chan *dbus.Call {
	return make(chan *dbus.Call, 10)
}

func Update(c *dbus.Call) {
	computeAmbBr(c)
	Set(computeNextBl())
}

func Set(val float64) {
	if	err := api.GoSetAll(val, &conf.BSmooth, nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Set %.2f backlight level.\n", val)
		state.CurBl = val
	}
}

func computeAmbBr(call *dbus.Call) {
	var Sensor string
	var Val []float64
	if err := call.Store(&Sensor, &Val); err != nil {
		fmt.Println(err.Error())
	} else {
		state.AmbBr = stat.Mean(Val, nil)
		fmt.Printf("'%s' captured %d/%d frames; Avg brightness: %.2f.\n", Sensor, len(Val), conf.NCaptures, state.AmbBr)
	}
}

func computeNextBl() float64 {
	return math.Min(BlFitParams[0] * (state.AmbBr * 10) + BlFitParams[1], 1)
}

func Close() {
	if err := api.Destroy(); err != nil {
		fmt.Println(err)
	}
}
