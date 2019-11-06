package kara

import "fmt"

var (
	TypeShouldOneError = fmt.Errorf("karaSpot.Type should be 1 when call SetWhenNotExist")
	TypeShouldTwoError = fmt.Errorf("karaSpot.Type should be 2 when call AddWhenNotReachLimit")

	LimitBurstError = fmt.Errorf("karaSpot.Map[key] burst its limit karaSpot.Limit")
)
