package main

import (
	"time"

	"github.com/tal-tech/go-zero/core/proc"
)

const (
	wrapInterval     = time.Second * 10
	shutdownInterval = time.Second * 3
)

func main() {
	start := time.Now()
	proc.AddWrapUpListener(func() {
		time.Sleep(wrapInterval)
		println(time.Since(start)/time.Second, 1)
	})
	proc.AddWrapUpListener(func() {
		time.Sleep(wrapInterval)
		println(time.Since(start)/time.Second, 2)
	})
	proc.AddShutdownListener(func() {
		time.Sleep(shutdownInterval)
		println(time.Since(start)/time.Second, 3)
	})
	proc.AddShutdownListener(func() {
		time.Sleep(shutdownInterval)
		println(time.Since(start)/time.Second, 4)
	})
	select {}
}
