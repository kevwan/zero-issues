package main

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/stores/redis"
)

func printVal(red *redis.Redis, key string) {
	val, err := red.Get(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func main() {
	const (
		addr = "localhost:6379"
		key  = "lock"
	)
	red := redis.New(addr)
	lock := redis.NewRedisLock(red, key)
	lock.Acquire()
	fmt.Println("locked once")
	lock.Acquire()
	fmt.Println("locked twice")
	printVal(red, key)
	lock.Release()
	fmt.Println("released once")
	printVal(red, key)
	lock.Release()
	fmt.Println("released twice")
	printVal(red, key)
}
