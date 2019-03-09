package gotest

import (
	// "context"
	// "errors"
	"fmt"
	// "github.com/prashantv/gostub"
	// "io"
	// "os"
	// "reflect"
	"runtime"
	// "strings"
	// "sync"
	// "syscall"
	"testing"
	"time"
)

var intMap map[int]int
var cnt = 8192

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
func TestAll(ti *testing.T) {
	printMemStats()
	runtime.GC()
	printMemStats()
	initMap()
	runtime.GC()
	printMemStats()

	fmt.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	fmt.Println(len(intMap))

	runtime.GC()
	printMemStats()

	intMap = nil
	runtime.GC()
	printMemStats()
	time.Sleep(time.Minute)
}
