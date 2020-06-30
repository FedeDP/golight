package state

import (
	"github.com/maltegrosse/go-geoclue2"
	"github.com/FedeDP/golight/day"
	"time"
)

var(
	NextSunrise				 	time.Time
	NextSunset 					time.Time
	Location 					geoclue2.GeoclueLocation
	DayTime 					day.Time
	AmbBr						float64
	CurBl						float64
	OnBatt						bool
 	Display						DisplayState
)
