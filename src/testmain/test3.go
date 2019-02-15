package main

import (
	"fmt"
	// "unsafe"
	"sync/atomic"
	// "time"
	// "errors"
	// "flag"
	"context"
	"runtime/pprof"
	// "runtime/debug"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func test3cpuprofile() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(dir, "cpuprofile.out")
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	pprof.StartCPUProfile(f)
	for i := 0; i < 100000; i++ {
		_ = i * i
		time.Sleep(time.Millisecond)
	}

	pprof.StopCPUProfile()
}
func test3forstrings() {

	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))

	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	// 示例2。
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	// 示例3。
	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	n, _ = reader1.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))

	// var builder1 strings.Builder
	// builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	// fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	// fmt.Println()
	// builder1.WriteByte(' ')
	// builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	// builder1.Write([]byte{'\n', '\n'})
	// builder1.WriteString("Do not copy a non-zero Builder.")
	// fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	// fmt.Println()

	// // 示例2。
	// fmt.Println("Grow the builder ...")
	// builder1.Grow(10)
	// fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	// fmt.Println()
	// b := &builder1
	// b.WriteString("================================")
	// fmt.Printf("b(%d):\n\"%s\"\n", b.Len(), b.String())

	// // 示例3。
	// fmt.Println("Reset the builder ...")
	// builder1.Reset()
	// fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())

}
func test3syncpool() {
	// defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}
	v1 := pool.Get()
	fmt.Println(v1)
	pool.Put(100)
	pool.Put(200)
	pool.Put(300)
	v2 := pool.Get()
	fmt.Println(v2)
	// runtime.GC()
	time.Sleep(3 * time.Minute)
	v3 := pool.Get()
	fmt.Println(v3)

	pool.New = nil
	v4 := pool.Get()
	fmt.Println(v4)
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
