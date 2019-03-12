package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
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
func runcmd(cmds []*exec.Cmd)(pid string,err error){
	//reader,writer,err:=os.Pipe()
	//if err!=nil{
	//	return
	//}
	fmt.Println("runcmd")
	var readBuff bytes.Buffer
	var writeBuff bytes.Buffer
	for i,v:=range cmds{
		fmt.Println("runcmd:",i)
		fmt.Println(readBuff.String())
		fmt.Println(writeBuff.String())
		if i>0{
			v.Stdin=&readBuff
		}
		v.Stdout=&writeBuff
		if err=v.Start();err!=nil{
			return
		}
		if err=v.Wait();err!=nil{
			return
		}
		readBuff.Write(writeBuff.Bytes())
		writeBuff.Reset()
	}
	pid=writeBuff.String()
	return
}
func sendsignal(){
	time.Sleep(time.Second*10)
	cmds:=[]*exec.Cmd{
		exec.Command("ps","aux"),
		exec.Command("grep","gowallet"),
		exec.Command("grep","-v","grep"),
		exec.Command("awk","{print $2}"),
	}
	pid,err:=runcmd(cmds)
	if err!=nil{
		return
	}
	fmt.Println("pid:",pid)
	p,err:=strconv.Atoi(pid)

	proc,err:=os.FindProcess(p)
	err=proc.Signal(syscall.SIGINT)
	if err!=nil{
		return
	}
}
func main() {
	sig:=make(chan os.Signal,1)
	sigs:=[]os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	signal.Notify(sig,sigs...)

	sig2:=make(chan os.Signal,1)
	sigs2:=[]os.Signal{syscall.SIGINT}
	signal.Notify(sig2,sigs2...)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for s:=range sig{
			fmt.Print(s)
		}
		fmt.Println("done sig")
		wg.Done()
	}()
	go func() {
		for s:=range sig2{
			fmt.Print(s)
		}
		fmt.Println("done sig2")
		wg.Done()
	}()
	time.Sleep(time.Second)
	signal.Stop(sig2)
	close(sig2)
	go sendsignal()
	wg.Wait()
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
