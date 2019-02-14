package main

import (
	"fmt"
	// "unsafe"
	"sync/atomic"
	// "time"
	// "errors"
	// "flag"
	"context"
	// "sync"
	"time"
)

func test3syncpool() {
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}
	v1 := pool.Get()
	fmt.Println(v1)
}
func someHandler() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doStuff(ctx)
	cancel()
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	// cancel()
	node2 := context.WithValue(ctx, "xx", "yy")
	fmt.Println(node2.Value("xx"))
}
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			fmt.Println("default")
		}
	}
}
