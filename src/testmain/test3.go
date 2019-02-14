package main

import (
	"fmt"
	// "unsafe"
	// "sync/atomic"
	// "time"
	"errors"
	// "flag"
	"context"
	"sync"
	"time"
)

func someHandler() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

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
