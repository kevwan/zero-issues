package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	var val *redis.IntCmd
	rds := redis.New("localhost:6379")
	err := rds.Pipelined(func(pipeliner redis.Pipeliner) error {
		val = pipeliner.Exists(context.Background(), "foo", "bar")
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val.Val())
}
