package main

import (
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
)

func main() {
	group := threading.NewRoutineGroup()
	for i := 0; i < 100; i++ {
		group.Run(func() {
			for j := 0; j < 1000; j++ {
				logx.WithDuration(time.Second).Info("hello")
			}
		})
	}
	group.Wait()
}
