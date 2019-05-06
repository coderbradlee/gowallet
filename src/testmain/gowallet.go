package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var counter = 0

// var Testvar = "main"

// type sliceValue []string

// //new一个存放命令行参数值的slice
// func newSliceValue(vals []string, p *[]string) *sliceValue {
// 	*p = vals
// 	return (*sliceValue)(p)
// }

// /*
// Value接口：
// type Value interface {
//     String() string
//     Set(string) error
// }
// 实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
// */
// func (s *sliceValue) Set(val string) error {
// 	*s = sliceValue(strings.Split(val, ","))
// 	return nil
// }

// //flag为slice的默认值default is me,和return返回值没有关系
// func (s *sliceValue) String() string {
// 	*s = sliceValue(strings.Split("default is me", ","))
// 	return "It's none of my business"
// }

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
func runcmd(cmds []*exec.Cmd) (pids []string, err error) {
	//reader,writer,err:=os.Pipe()
	//if err!=nil{
	//	return
	//}
	fmt.Println("runcmd")
	var readBuff bytes.Buffer
	var writeBuff bytes.Buffer
	for i, v := range cmds {
		fmt.Println("runcmd:", i)
		fmt.Println(readBuff.String())
		fmt.Println(writeBuff.String())
		if i > 0 {
			v.Stdin = &readBuff
		}
		v.Stdout = &writeBuff
		if err = v.Start(); err != nil {
			return
		}
		if err = v.Wait(); err != nil {
			return
		}
		readBuff.Write(writeBuff.Bytes())
		writeBuff.Reset()
	}
	temp := readBuff.String()
	pids = strings.Split(temp, "\n")
	//if len(pids)>0{
	//	pid=pids[0]
	//}
	return
}
func sendsignal() {
	time.Sleep(time.Second * 10)
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "gowallet"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("awk", "{print $2}"),
	}
	pids, err := runcmd(cmds)
	if err != nil {
		return
	}
	fmt.Println("pid:", pids)
	for _, pid := range pids {
		p, err := strconv.Atoi(pid)
		if err != nil {
			return
		}
		fmt.Println("p:", p)
		proc, err := os.FindProcess(p)
		err = proc.Signal(syscall.SIGKILL)
		if err != nil {
			fmt.Println("sendsignal:", err)
			return
		}
	}
}
func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(time.Second * 10))
		readBytes, isprefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println("server error:", err)
			break
		}
		if !isprefix {
			fmt.Println("server read end:", string(readBytes))

			var buffer bytes.Buffer
			buffer.Write(readBytes[:len(readBytes)])
			buffer.WriteByte('\n')
			n, err := conn.Write(buffer.Bytes())
			if err != nil {
				fmt.Println("server error:", err)
				break
			}
			fmt.Println("server write:", n)

			break
		}
	}
}
func server() {
	listener, err := net.Listen("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server error:", err)
			break
		}
		go handleConn(conn)
	}
}
func client() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:80", time.Second*2)
	if err != nil {
		fmt.Println("client conn:", err)
		return
	}
	defer conn.Close()
	for {
		var buffer bytes.Buffer
		buffer.Write([]byte("string test\n"))
		//buffer.WriteByte('\n')
		n, err := conn.Write(buffer.Bytes())
		if err != nil {
			fmt.Println("client write:", err)
			break
		}
		fmt.Println("client write:", n)
		reader := bufio.NewReader(conn)
		content, isprefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println("client read:", err)
			break
		}
		if !isprefix {
			fmt.Println("client end read:", string(content))

			break
		}

	}
}
func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
func f() {
	container := make([]int, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			traceMemStats()
		}
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}
func test() {
	log.Println("start.")
	traceMemStats()
	f()
	traceMemStats()

	log.Println("force gc.")
	runtime.GC()

	log.Println("done.")
	traceMemStats()

	go func() {
		for {
			traceMemStats()
			time.Sleep(10 * time.Second)
		}
	}()
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()
}
func main() {
	test()
	//go server()
	//time.Sleep(time.Second)
	//go client()
	//sig:=make(chan os.Signal,1)
	//sigs:=[]os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	//signal.Notify(sig,sigs...)
	//
	//sig2:=make(chan os.Signal,1)
	//sigs2:=[]os.Signal{syscall.SIGINT}
	//signal.Notify(sig2,sigs2...)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	for s:=range sig{
	//		fmt.Print("ssss:",s)
	//	}
	//	fmt.Println("done sig")
	//	wg.Done()
	//}()
	//go func() {
	//	for s:=range sig2{
	//		fmt.Print(s)
	//	}
	//	fmt.Println("done sig2")
	//	wg.Done()
	//}()
	//time.Sleep(time.Second)
	//signal.Stop(sig2)
	//close(sig2)
	//go sendsignal()
	//wg.Wait()
	//signal.Stop(sig)
	//defer func() {
	//	if p := recover(); p != nil {
	//		fmt.Println(p)
	//	}
	//}()
	//x := 1
	//fmt.Println(x)     //prints 1
	//{
	//	fmt.Println(x) //prints 1
	//	x := 2
	//	fmt.Println(x) //prints 2
	//}
	//fmt.Println(x)
	// test()
	// someHandler()
	// test3syncpool()
	// test3forstrings()
	// test3cpuprofile()
	//test333333()
	//fmt.Println(http.ListenAndServe(":80", nil))
	//
	//fmt.Println("before wait")
	ch := make(chan int, 0)
	<-ch
	// var name = flag.String("name", "everyone", "name for test")
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// // flag.Parse()
	// var languages []string
	// flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	// flag.Parse()

	// //打印结果slice接收到的值
	// fmt.Println(languages)

	// fmt.Println(":", *name)
	// // var Testvar = "iner"
	// fmt.Println(testvar)
	// dynamic()
	// testbfs()
	// var choice []int
	// for i := 0; i < 16; i++ {
	// 	choice = append(choice, i)
	// }
	// var results []int
	// combine(choice, results, 14)
	// fmt.Println("counter:", counter)
	// choice := []string{"a", "b", "c", "d", "e"}
	// {
	// 	newchoice := rmstring(choice, "a")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "b")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "c")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "d")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "e")
	// 	fmt.Println(newchoice)
	// }
	// var results []string
	// permutate(choice, results)
	// fmt.Println("imtoken address:", ethaddress)
	// fmt.Printf("\n")
	// testbtcsign()
	// testfunc()
	// var results []int
	// getReward(10, results)
	// left := []int{1, 4, 8}
	// right := []int{2, 10, 20, 33, 55, 66, 77}
	// ret := mergeSort(left, right)
	// fmt.Println(ret)
	// test := []int{2, 9, 4, 10, 3, 20, 5, 33, 1, 88, 11, 25, 6}
	// ret := merges(test)
	// fmt.Println(ret)
	// to_sort := []int{3434, 3356, 67, 123, 111, 890}
	// sorted := merges(to_sort)
	// fmt.Println(sorted)
	// dv(8, results)
	// fmt.Println("===================")
	// dv(15, results)
	// dv(16, results)
	// dv(15, results)
	// test()
	// testipfs()
}
