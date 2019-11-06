package kara

import (
	"fmt"
	"sync"
)

const (
	Version = "v1.0.1"
)

var (
	KaraPool sync.Map
)

const (
	NOT_EXIST = 1
	EXIST     = 2
)

type KaraSpot struct {
	Map map[string]int

	// 1- if exist, 2- times increment
	// When Type=1, karaSpot.Map will has its value ranged [1,2], 1-Not Exist, 2-Exist
	// When Type=2, karaSpot.Map will has its value ranged [0,1,2,3...]
	Type int

	// servers when karaSpot.Type=2, it will limit karaSpot.Map's values maximum
	limit int
	L     *sync.RWMutex
}

// New a exist-type spot
func NewSpot() *KaraSpot {
	return &KaraSpot{
		Map:   make(map[string]int, 0),
		Type:  1,
		limit: 0,
		L:     &sync.RWMutex{},
	}
}

// New a range-type spot with limit
func NewTimesSpot(limit int) *KaraSpot {
	return &KaraSpot{
		Map:   make(map[string]int, 0),
		Type:  2,
		limit: limit,
		L:     &sync.RWMutex{},
	}
}

func (ks *KaraSpot) SetWhenNotExist(key string) (bool, error) {
	if ks.Type != 1 {
		return false, TypeShouldOneError
	}

	ks.L.Lock()
	defer ks.L.Unlock()

	v, ok := ks.Map[key]

	if !ok || v == NOT_EXIST {
		ks.Map[key] = EXIST
	} else {
		// existed
		return false, nil
	}
	return true, nil
}

func (ks *KaraSpot) AddWhenNotReachedLimit(key string) (bool, error) {
	if ks.Type != 2 {
		return false, TypeShouldTwoError
	}

	ks.L.Lock()
	defer ks.L.Unlock()

	v, ok := ks.Map[key]

	if !ok {
		if ks.limit <= 0 {
			return false, nil
		}
		ks.Map[key] = 1
		return true, nil
	}

	if v == ks.limit {
		return false, nil
	}

	if v < ks.limit {
		ks.Map[key] = v + 1
		return true, nil
	}

	if v > ks.limit {
		return false, LimitBurstError
	}
	return false, fmt.Errorf("unexpected trace")
}
