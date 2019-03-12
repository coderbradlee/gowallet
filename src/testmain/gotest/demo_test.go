package gotest

import (
	"bufio"
	//"bufio"
	// "strings"
	// "sync"
	// "syscall"
	//"math"
	"fmt"
	"os/exec"
	"testing"
	"time"

	//"sys"
	//"sync"
)

type testinterface interface {
	xx(int)
	yy(int)
	print()
}
type teststruct struct {
	i int
}
func (t *teststruct)xx(v int){
	t.i=v
}
func (t teststruct)yy(v int){
	t.i=v
}
func (t teststruct)print(){
	fmt.Println(t.i)
}
type ttstruct struct {
	*teststruct
}
func TestAll2(ti *testing.T) {
	cmd0:=exec.Command("echo","-n","xxx")
	if err:=cmd0.Start();err!=nil{
		fmt.Println("x:",err)
		return
	}
	stdout,err:=cmd0.StdoutPipe()
	if err!=nil{
		fmt.Println("y:",err)
		return
	}
	output:=bufio.NewReader(stdout)
	out,e,err:=output.ReadLine()
	fmt.Println(out,":",e,":",err)
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
