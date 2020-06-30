package state

type DisplayState uint8 // [ 0, 1, 2 ] -> bitmask (i want DisplayON to be 0)
const(
	DisplayON 		= iota
	DisplayDIM
	DisplayOFF
)
func DisplaySet(flag DisplayState) 		DisplayState { return Display | flag }
func DisplayClear(flag DisplayState) 	DisplayState { return Display &^ flag }