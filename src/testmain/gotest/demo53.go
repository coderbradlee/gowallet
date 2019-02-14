package gotest

import (
	"errors"
	// "flag"
	"context"
	"fmt"
	"math"
	"sync"
	"time"
	// "io"
	// "os"
	"reflect"
	"sync/atomic"
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

var name string

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}
func showNum(id int, deferfunc func()) {
	defer func() {
		deferfunc()
	}()
	fmt.Println(id)
}
func addNum(num *int32, id, max int32, deferfunc func()) {
	defer func() {
		deferfunc()
	}()
	for i := 0; ; i++ {
		numnow := atomic.LoadInt32(num)
		if numnow >= max {
			break
		}
		add := numnow + 2
		if atomic.CompareAndSwapInt32(num, numnow, add) {
			fmt.Println(id, ":changeto", add)
		} else {
			fmt.Println(id, ":fail:", add)
		}
	}
}
func test() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	go addNum(&num, 1, 10, wg.Done)
	go addNum(&num, 2, 10, wg.Done)
	wg.Wait()
}
func testx() {
	ch := make(chan struct{}, 2)
	num := int32(0)
	go addNum(&num, 1, 10, func() {
		ch <- struct{}{}
		fmt.Println("defer:", 1)
	})
	go addNum(&num, 2, 10, func() {
		ch <- struct{}{}
		fmt.Println("defer:", 2)
	})
	<-ch
}

// func test33() {
// 	var box atomic.Value
// 	fmt.Println("Copy box to box2.")
// 	box2 := box // 原子值在真正使用前可以被复制。
// 	v1 := [...]int{1, 2, 3}
// 	fmt.Printf("Store %v to box.\n", v1)
// 	box.Store(v1)
// 	fmt.Printf("The value load from box is %v.\n", box.Load())
// 	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
// 	fmt.Println()

// 	// 示例2。
// 	v2 := "123"
// 	fmt.Printf("Store %q to box2.\n", v2)
// 	box2.Store(v2) // 这里并不会引发panic。
// 	fmt.Printf("The value load from box is %v.\n", box.Load())
// 	fmt.Printf("The value load from box2 is %q.\n", box2.Load())
// 	fmt.Println()

// 	// 示例3。
// 	fmt.Println("Copy box to box3.")
// 	box3 := box // 原子值在真正使用后不应该被复制！
// 	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
// 	v3 := 123
// 	fmt.Printf("Store %d to box2.\n", v3)
// 	// box3.Store(v3) // 这里会引发一个panic，报告存储值的类型不一致。
// 	_ = box3
// 	fmt.Println()

// 	// 示例4。
// 	var box4 atomic.Value
// 	v4 := errors.New("something wrong")
// 	fmt.Printf("Store an error with message %q to box4.\n", v4)
// 	box4.Store(v4)
// 	v41 := io.EOF
// 	fmt.Println("Store a value of the same type to box4.")
// 	box4.Store(v41)
// 	v42, ok := interface{}(&os.PathError{}).(error)
// 	if ok {
// 		fmt.Printf("Store a value of type %T that implements error interface to box4.\n", v42)
// 		// box4.Store(v42) // 这里会引发一个panic，报告存储值的类型不一致。
// 	}
// 	fmt.Println()

// 	// 示例5。
// 	box5, err := NewAtomicValue(v4)
// 	if err != nil {
// 		fmt.Printf("error: %s\n", err)
// 	}
// 	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
// 	fmt.Println("Store a value of the same type to box5.")
// 	err = box5.Store(v41)
// 	if err != nil {
// 		fmt.Printf("error: %s\n", err)
// 	}
// 	fmt.Printf("Store a value of type %T that implements error interface to box5.\n", v42)
// 	err = box5.Store(v42)
// 	if err != nil {
// 		fmt.Printf("error: %s\n", err)
// 	}
// 	fmt.Println()

// 	// 示例6。
// 	var box6 atomic.Value
// 	v6 := []int{1, 2, 3}
// 	fmt.Printf("Store %v to box6.\n", v6)
// 	box6.Store(v6)
// 	v6[1] = 4 // 注意，此处的操作不是并发安全的！
// 	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
// 	// 正确的做法如下。
// 	v6 = []int{1, 2, 3}
// 	store := func(v []int) {
// 		replica := make([]int, len(v))
// 		copy(replica, v)
// 		box6.Store(replica)
// 	}
// 	fmt.Printf("Store %v to box6.\n", v6)
// 	store(v6)
// 	v6[2] = 5 // 此处的操作是安全的。
// 	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
// }
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
func cond() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.RWMutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)

	recvCond := sync.NewCond(&lock)

	// send 代表用于发信的函数。
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		fmt.Printf("put [%d-%d].\n",
			id, index)
		mailbox = 1

		lock.Unlock()
		recvCond.Broadcast()
		fmt.Printf("Broadcast [%d-%d].\n",
			id, index)
	}

	// recv 代表用于收信的函数。
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		fmt.Printf("take [%d-%d].\n",
			id, index)
		mailbox = 0

		lock.Unlock()
		sendCond.Signal() // 确定只会有一个发信的goroutine。
		fmt.Printf("Signal [%d-%d].\n",
			id, index)
	}

	// sign 用于传递演示完成的信号。
	// sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) { // 用于发信。
		for i := 1; i <= max; i++ {
			send(id, i)
		}
	}(0, max)
	go func(id, max int) { // 用于收信。
		for j := 1; j <= max; j++ {
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { // 用于收信。
		for k := 1; k <= max; k++ {
			recv(id, k)
		}
	}(2, max/2)

	// recvCond 代表专用于收信的条件变量。
	// recvCond := sync.NewCond(lock.RLocker())

	// // sign 用于传递演示完成的信号。
	// // sign := make(chan struct{}, 3)
	// max := 5
	// go func(max int) { // 用于发信。
	// 	// defer func() {
	// 	// 	sign <- struct{}{}
	// 	// }()
	// 	for i := 1; i <= max; i++ {
	// 		// time.Sleep(time.Millisecond * 500)
	// 		lock.Lock()
	// 		fmt.Println("put:", i)
	// 		for mailbox == 1 {
	// 			sendCond.Wait()
	// 		}
	// 		fmt.Printf("[%d]: the mailbox is empty.\n", i)
	// 		mailbox = 1
	// 		lock.Unlock()
	// 		recvCond.Signal()
	// 		fmt.Printf("[%d]: put 1 mail.\n", i)
	// 	}
	// }(max)
	// go func(max int) { // 用于发信。
	// 	// defer func() {
	// 	// 	sign <- struct{}{}
	// 	// }()
	// 	for i := 1; i <= max; i++ {
	// 		// time.Sleep(time.Millisecond * 500)
	// 		lock.Lock()
	// 		fmt.Println("put:", i+max)
	// 		for mailbox == 1 {
	// 			sendCond.Wait()
	// 		}
	// 		fmt.Printf("[%d]: the mailbox is empty.\n", i+max)
	// 		mailbox = 1
	// 		lock.Unlock()
	// 		recvCond.Signal()
	// 		fmt.Printf("[%d]: put 1 mail.\n", i+max)
	// 	}
	// }(max)
	// go func(max int) { // 用于收信。
	// 	// defer func() {
	// 	// 	sign <- struct{}{}
	// 	// }()
	// 	for j := 1; j <= max; j++ {
	// 		// time.Sleep(time.Millisecond * 500)
	// 		lock.RLock()
	// 		fmt.Println("take:", j)
	// 		for mailbox == 0 {
	// 			recvCond.Wait()
	// 		}
	// 		fmt.Printf("[%d]: the mailbox is full.\n", j)
	// 		mailbox = 0

	// 		lock.RUnlock()
	// 		sendCond.Signal()
	// 		fmt.Printf("[%d]: take 1 mail.\n", j)
	// 	}
	// }(max * 2)

	// <-sign
	// <-sign
	ch := make(chan struct{}, 0)
	<-ch
}

// func init() {
// 	flag.StringVar(&name, "name", "everyone", "The greeting object.")
// }

// func main() {
// 	flag.Parse()
// 	greeting, err := hello(name)
// 	if err != nil {
// 		fmt.Printf("error: %s\n", err)
// 		return
// 	}
// 	fmt.Println(greeting, introduce())
// }

// hello 用于生成问候内容。
func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

// introduce 用于生成介绍内容。
func introduce() string {
	return "Welcome to my Golang column."
}
