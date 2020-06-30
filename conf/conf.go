package conf

type BacklightSmooth struct {
	Smooth 	bool
	Step  	float64
	Timeout	uint32
}

type GammaSmooth struct {
	Smooth 	bool
	Step  	uint32
	Timeout	uint32
}

var (
	BSmooth		=				 	BacklightSmooth{ true, 0.05, 30 }
	GSmooth 	=					GammaSmooth { true, 50, 300 }
	NCaptures 	= 					int32(5)
	Temps		=					[]int32{ 6500, 4000 }
	BlRegPoints =					[]float64{0.0, 0.15, 0.29, 0.45, 0.61, 0.74, 0.81, 0.88, 0.93, 0.97, 1.0}
	DimmerTO	= 					uint(30)
	DpmsTO		=					uint(600)
	CaptureTO	= 					300
)
