package conf

import (
	"github.com/FedeDP/golight/internal/app/golight/state"
	"github.com/FedeDP/golight/pkg/go-clightd"
)

var (
	BSmooth		=				 	clightd.BacklightSmooth{ Smooth: true, Step: 0.05, Timeout: 30 }
	GSmooth 	=					clightd.GammaSmooth { Smooth: true, Step: 50, Timeout: 300 }
	NCaptures 	= 					[state.AcSize]int32{5, 5}
	Temps		=					[state.TimeSize]int32{ 6500, 4000 }
	BlRegPoints =					[]float64{0.0, 0.15, 0.29, 0.45, 0.61, 0.74, 0.81, 0.88, 0.93, 0.97, 1.0}
	DimmerTO	= 					[state.AcSize]uint{45, 30}
	DpmsTO		=					[state.AcSize]uint{600, 300}
	CaptureTO	= 					[state.AcSize]uint{300, 600}
)
