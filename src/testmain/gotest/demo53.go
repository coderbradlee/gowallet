package gotest

import (
	"errors"
	// "flag"
	"fmt"
	"math"
	"sync"
	// "time"
)

var name string

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
