package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	close = "close"
	how   = 0
)

func sigDel() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGPIPE)

	go func() {
		sig := <-sigs
		fmt.Println("sig:", sig)
	}()
}
func main() {
	sigDel()
	flag.StringVar(&close, "close", close, "close or shutdown")
	flag.IntVar(&how, "how", how, "how shutdown")
	flag.Parse()
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening... ")
	for {
		c, err := l.Accept()
		// c.(*net.TCPConn).CloseWrite()
		if err != nil {
			panic(err)
		}
		tcpConn, ok := c.(*net.TCPConn)
		if !ok {
			//error handle
			panic(err)
		}
		tcpConn.SetLinger(0)
		log.Println("get conn", c.RemoteAddr())
		go func() {
			go func() {
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
			go func() {
				for i := 10000; ; i++ {
					_, err := c.Write([]byte(strconv.Itoa(i)))
					if err != nil {
						fmt.Println("write err", err)
						return
					}
					time.Sleep(time.Second)
					fmt.Println("write", i)
				}
			}()
			if close == "close" {
				time.Sleep(time.Second * 3)
				if err := c.Close(); err != nil {
					panic(err)
				}
				log.Println("closed")
			} else {
				switch how {
				case 0:
					c.(*net.TCPConn).CloseRead() // syscall.Shutdown(fd, 0)
				case 1:
					c.(*net.TCPConn).CloseWrite() // syscall.Shutdown(fd, 1)
				case 2:
					c.(*net.TCPConn).Close()
				}
				// f, _ := c.(*net.TCPConn).File()
				// if err := syscall.Shutdown(int(f.Fd()), how); err != nil {
				//  panic(err)
				// }
				log.Println("shutdown, how:", how)
			}
		}()
	}
}
