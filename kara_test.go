package kara

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestExistSpot(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	spot := NewSpot()

	var key = "single_job_x"
	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()
		ok, e := spot.SetWhenNotExist(key)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}

func TestTimesSpot(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	spot := NewTimesSpot(10)

	var key = "limit_times_job_x"
	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()
		ok, e := spot.AddWhenNotReachedLimit(key)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}
