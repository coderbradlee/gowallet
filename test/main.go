package test

import (
	"log"
	"runtime"
	"time"
)

func f() {
	container := make([]int, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}
func testfmt() {
	log.Println("start.")
	f()

	log.Println("force gc.")
	runtime.GC() // 调用强制gc函数

	log.Println("done.")
	time.Sleep(1 * time.Hour) // 保持程序不退出
}
func main() {
	testfmt()
}
