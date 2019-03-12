package gotest

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	//"os"
	"testing"
	"time"

	//"sys"
	//"sync"
)


func TestAll2(ti *testing.T) {
	sig:=make(chan os.Signal,1)
	sigs:=[]os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	signal.Notify(sig,sigs...)
	for s:=range sig{
		fmt.Print(s)
	}
	//reader,writer,err:=os.Pipe()
	//reader,writer:=io.Pipe()
	//go func() {
	//	var buffread =make([]byte,100)
	//	fmt.Println("before read")
	//	n,err:=reader.Read(buffread)
	//	fmt.Println("after read")
	//	if err!=nil{
	//		fmt.Print(err)
	//		return
	//
	//	}
	//	fmt.Println("read:",n,":",string(buffread))
	//}()
	//
	//var buffwrite []byte=[]byte("xxxxinput")
	//fmt.Println("before write")
	//n2,err:=writer.Write(buffwrite)
	//fmt.Println("after write")
	//if err!=nil{
	//	fmt.Print(err)
	//	return
	//}
	//fmt.Println("write:",n2)




	//cmd1:=exec.Command("ps","aux")
	//cmd2:=exec.Command("grep","go")
	//var buff bytes.Buffer
	//var buff2out bytes.Buffer
	//cmd1.Stdout=&buff
	//cmd2.Stdin=&buff
	//cmd2.Stdout=&buff2out
	//if err:=cmd1.Start();err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//if err:=cmd1.Wait();err!=nil{
	//	fmt.Println(err)
	//}
	//if err:=cmd2.Start();err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//if err:=cmd2.Wait();err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(buff2out.String())
	//cmd0:=exec.Command("echo","-n","xxx")
	//stdout,err:=cmd0.StdoutPipe()
	//if err:=cmd0.Start();err!=nil{
	//	fmt.Println("x:",err)
	//	return
	//}
	//
	//output:=bufio.NewReader(stdout)
	//out,e,err:=output.ReadLine()
	//fmt.Println(string(out),":",e,":",err)
	//for i:=0;i<5;i++{
	//	defer func(n int) {
	//		fmt.Println(n)
	//	}(i)
	//}
	//s:=[5]int{1,2,3,4,5}
	//go func() {
	//	s[4]=1000
	//}()
	//for i,v:=range s{
	//	fmt.Println(i,":",v)
		time.Sleep(time.Second)
	//}
	//i:=0
	//for i!=0{
	//	fmt.Println("x")
	//}
	//var v interface{}=5
	//switch v.(type){
	//case int,string:
	//	fmt.Println("int string")
	//default:
	//	fmt.Println("default")
	//}
	//var ts *ttstruct=&ttstruct{&teststruct{}}
	////var ts testinterface =teststruct{i:10}
	//ts.teststruct.xx(20)
	//ts.print()
	//ts.yy(30)
	//ts.print()
}
