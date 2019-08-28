package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func sigDel() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
	}()
}
func main() {
	sigDel()
	c, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			time.Sleep(time.Second * 1)
			_, err := c.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				fmt.Println("write err", err)
				return
			}
			fmt.Println("write", i)
		}
	}()
	go func() {
		defer wg.Done()
		b := make([]byte, 1024)
		for {
			_, err := c.Read(b)
			if err != nil {
				fmt.Println("read err", err)
				return
			}
			fmt.Println("read", string(b))
		}
	}()
	wg.Wait()
}
