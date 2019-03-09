package gotest

import (
	// "context"
	// "errors"
	"fmt"
	"math/rand"

	// "github.com/prashantv/gostub"
	// "io/reader"
	// "bufio"
	// "os"
	// "reflect"
	"runtime"
	// "strings"
	// "sync"
	// "syscall"
	"math"
	"testing"
	"time"
)

var intMap map[int]int
var cnt = 8192

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

type student2 struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student2:
		fmt.Println("*:", msg)
	case student2:
		// msg.Name = "qq"
		fmt.Println(msg)
	}
}
func SpiltList(list []int, size int) {
	lens := len(list)
	mod := math.Ceil(float64(lens) / float64(size))
	spliltList := make([][]int, 0)
	for i := 0; i < int(mod); i++ {
		tmpList := make([]int, 0, size)
		fmt.Println("i=", i)
		if i == int(mod)-1 {
			tmpList = list[i*size:]
		} else {
			tmpList = list[i*size : i*size+size]
		}
		spliltList = append(spliltList, tmpList)
	}
	for i, sp := range spliltList {
		fmt.Println(i, " ==> ", sp)
	}
}
func MutilParam(p ...interface{}) {
	fmt.Println("MutilParam=", p)
}
func testslice(b []int) {
	// b = make([]int, 0)
	// b = append(b, 1)
	b[0] = 1
	fmt.Println(b)
}
func getRewardUser(user []string,sum int,offsetUser map[string]int)(name string){
	//limit:=len(user)
	num:=rand.Intn(sum)
	//fmt.Println(num)
	for _,v:=range user{
		//fmt.Println(i,"::::::",v,"$$$$$",offsetUser[v],"****",num)

		if offsetUser[v]>=num{
			name=v
			//fmt.Println("name:",name)
			return
		}
	}
	return
}
func TestAll(ti *testing.T) {
	user:=map[string]int{
		"a":55,
		"b":11,
		"c":33,
		"d":88,
		"e":6,
	}
	userArray:=make([]string,0)
	for k,_:=range user{
		userArray=append(userArray,k)
	}
	offsetUser:=make(map[string]int)
	sum:=0
	for _,v:=range userArray{
		sum+=user[v]
		offsetUser[v]=sum
	}
	for k,v:=range offsetUser{
		fmt.Println(k,":",v)
	}
	stat:=make(map[string]int)
	for i:=0;i<1000;i++{
		us:=getRewardUser(userArray,sum,offsetUser)
		stat[us]++
	}
	for k,v:=range stat{
		fmt.Println(k,":",v)
	}
	//b := make([]int, 10)
	//c := []int{1, 2, 3, 4, 5, 6, 7}
	//copy(b, c)
	//fmt.Println(b)
	// reader := bufio.NewReader(strings.NewReader("ABCDEFG\nHIJKLMN\n"))
	// var bts = make([]byte, 100) // 注意点：需要初始化
	// _, err := reader.Read(bts)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// b := make([]int, 10)
	// // var b []int
	// testslice(b)
	// fmt.Println(b)
	// v := make([]int, 10)
	// fmt.Println(len(v))
	// fmt.Println(cap(v))
	// MutilParam("ssss", 1, 2, 3, 4) //[ssss 1 2 3 4]
	// iis := []int{1, 2, 3, 4}
	// MutilParam("ssss", iis)
	// temp := make([]interface{}, 0)
	// temp = append(temp, "ssss")
	// for _, v := range iis {
	// 	temp = append(temp, v)
	// }
	// MutilParam(temp...)
	// ff := reflect.ValueOf(MutilParam)
	// temp := make([]reflect.Value, 0)
	// temp = append(temp, reflect.ValueOf("ssss"))
	// for _, v := range iis {
	// 	temp = append(temp, reflect.ValueOf(v))
	// }
	// ff.Call(temp)
	// Annie := "Annie"
	// Betty := "Betty"
	// Charley := "Charley"
	// m := []*string{&Annie, &Betty, &Charley}
	// // for _, v := range m {
	// // 	temp := "xx"
	// // 	v = &temp
	// // }
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	// for _, v := range five {
	// 	fmt.Printf("v[%s]\n", v)
	// 	v = "xx"
	// }
	// for _, v := range five {
	// 	fmt.Println(v)
	// }
	// alpha := make(chan int)
	// num := make(chan int)
	// go func() {
	// 	for i := 65; i < 97; i++ {

	// 		fmt.Println(string(rune(i)))
	// 		num <- i
	// 		<-alpha
	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < 32; i++ {
	// 		<-num
	// 		fmt.Println(i)
	// 		alpha <- i
	// 	}
	// }()
	// alpha <- 1
	time.Sleep(time.Second * 3)
	// lenth := 11
	// size := 5
	// list := make([]int, 0, lenth)
	// for i := 0; i < lenth; i++ {
	// 	list = append(list, i)
	// }
	// SpiltList(list, size)
	// var s student2
	// zhoujielun(s)
	// zhoujielun(&s)
	// printMemStats()
	// runtime.GC()
	// printMemStats()
	// initMap()
	// runtime.GC()
	// printMemStats()

	// fmt.Println(len(intMap))
	// for i := 0; i < cnt; i++ {
	// 	delete(intMap, i)
	// }
	// fmt.Println(len(intMap))

	// runtime.GC()
	// printMemStats()

	// intMap = nil
	// runtime.GC()
	// printMemStats()
	// time.Sleep(time.Minute)
}
