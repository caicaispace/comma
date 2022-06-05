package util

import (
	"time"
)

func Ticker(second int64, fn func()) {
	d := time.Duration(second) * time.Second
	t := time.NewTicker(d)
	defer t.Stop()
	for {
		<-t.C
		fn()
	}
}

func Timer(second int64, fn func()) {
	d := time.Duration(second) * time.Second
	t := time.NewTimer(d)
	defer t.Stop()
	for {
		<-t.C
		fn()
		// need reset
		t.Reset(d)
	}
}
