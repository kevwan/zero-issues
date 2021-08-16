package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

func foo() {
	panic("oops")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	done := make(chan struct{})
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- fmt.Sprintf("%+v\n\n%s", p, strings.TrimSpace(string(debug.Stack())))
			}
		}()

		foo()
		close(done)
	}()

	select {
	case p := <-panicChan:
		panic(p)
	case <-done:
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
