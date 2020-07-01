package state

type AcState uint8
const(
	OnAc AcState		= iota
	OnBatt
	AcSize
)
