package day

import "time"

type Time uint8
const(
	Day 		= iota
	Night
)

func Subscribe() <- chan time.Time {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day() + 1, 0, 0, 5, 0, t.Location())
	return time.After(n.Sub(t))
}
