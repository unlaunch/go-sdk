package util

import (
	"time"
)

func RunImmediatelyAndSchedule(what func() error, delay time.Duration) chan bool {
	if delay <= 0 {
		panic("delay must be greater than 0")
	}

	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func Schedule(what func() error, delay time.Duration) chan bool {
	if delay <= 0 {
		panic("delay must be greater than 0")
	}

	ticker := time.NewTicker(delay)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <- ticker.C:
				what()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}
