package state

import (
	"github.com/maltegrosse/go-geoclue2"
	"time"
)

var(
	NextSunrise				 	time.Time
	NextSunset 					time.Time
	Location 					geoclue2.GeoclueLocation
	DayTime 					Time
	AmbBr						float64
	CurBl						float64
	Ac							AcState
 	Display						DisplayState
)
