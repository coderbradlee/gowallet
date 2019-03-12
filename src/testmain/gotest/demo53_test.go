package gotest

import (
	//"encoding/json"
	// "context"
	// "errors"
	"fmt"
	"math/rand"
	"sync"
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
func runwait(i int,wq <-chan interface{},done <-chan struct{},group *sync.WaitGroup){
	//fmt.Println(i," start")
	defer group.Done()
	//fmt.Println(i," done")
	for {
		select {
		case m := <- wq:
			fmt.Printf("%v do %v\n",i,m)
		case <- done:
			fmt.Printf("%v is done\n",i)
			return
		}
	}
}
type data struct {
	num int
	key *string
	items map[string]bool
}

func (this *data) pmethod() {
	this.num = 7
}

func (this data) vmethod() {
	this.num = 8
	*this.key = "v.key"
	this.items["vmethod"] = true
}
//type config struct {
//	Data []byte `json:"data"`
//}
//type config struct {
//	Data string `json:"data"`
//}
type datass struct {
	num int                //ok
	//checks [10]func() bool //not comparable
	doit func() bool       //not comparable
	m map[string] string   //not comparable
	bytes []byte           //not comparable
}
func doRecover() {
	fmt.Println("recovered =>",recover()) //prints: recovered => <nil>
}
func get() []byte {
	raw := make([]byte,10)
	fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
	raw[2]=1
	fmt.Println(raw)
	res := make([]byte,3)
	res[0]=2
	fmt.Println(res)
	//copy(res,raw[:3])
	//fmt.Println(res[0])
	return res
}
//type myLocker struct {
//	xx sync.Mutex
//}
type myLocker sync.Locker
type field struct {
	name string
}

func (p *field) print() {

	fmt.Println(p,":",p.name)
}
func (p field) print2() {

	fmt.Println(p,":",p.name)
}
type datadata struct {

}
func TestAll(ti *testing.T) {
	a := &data{}
	b := &data{}

	if a == b {
		fmt.Printf("same address - a=%p b=%p\n",a,b)
		//prints: same address - a=0x1953e4 b=0x1953e4
	}
	//data:=[]field{{"11"},{"22"},{"33"}}
	//for _,vv:=range data{
	//	go vv.print2()
	//	//go func(){
	//	//	fmt.Println(vv)
	//	//	vv.print()
	//	//}()
	//
	//}
	//data:=[]string{"11","22","33"}
	//for _,vv:=range data{
	//	v:=vv
	//	go func(){
	//		fmt.Println(v)
	//	}()
	//}
	//ch:=make(chan struct{})
	//<-ch
	//loop:
	//	for {
	//		switch {
	//		case true:
	//			fmt.Println("breaking out...")
	//			break loop
	//		}
	//	}
	//
	//fmt.Println("out!")
	//var lock myLocker = new(sync.Mutex)
	//lock.Lock() //ok
	//lock.Unlock() //ok
	//var lock myLocker
	//lock.xx.Lock() //ok
	//lock.xx.Unlock() //ok
	//path := []byte("AAAA/BBBBBBBBB")
	//sepIndex := bytes.IndexByte(path,'/')
	////dir1 := path[:sepIndex:sepIndex] //full slice expression
	//dir1:=make([]byte,sepIndex)
	//copy(dir1,path[:sepIndex])
	//dir2 := path[sepIndex+1:]
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	//
	//dir1 = append(dir1,"suffix"...)
	//path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	//
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)
	//
	//fmt.Println("new path =>",string(path))
	//path := []byte("AAAA/BBBBBBBBB")
	//sepIndex := bytes.IndexByte(path,'/')
	//dir1 := path[:sepIndex]
	//dir2 := path[sepIndex+1:]
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	//
	//dir1 = append(dir1,"suffix"...)
	//path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	//
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB (not ok)
	//
	//fmt.Println("new path =>",string(path))
	//data := get()
	//fmt.Println(len(data),cap(data),&data[0])
	//fmt.Println(data)
	//data := []*struct{num int} {{1},{2},{3}}
	//
	//for _,v := range data {
	//	v.num *= 10
	//}
	//
	//fmt.Println(data[0],data[1],data[2])
	//data := []int{1,2,3}
	//for i,_ := range data {
	//	data[i] *= 10
	//}
	//
	//fmt.Println("data:",data)
	//defer func() {
	//	//doRecover() //panic is not recovered
	//	fmt.Println("recovered =>",recover())
	//}()
	//
	//panic("not good")
	//var str string = "one"
	//var in interface{} = "one"
	//fmt.Println("str == in:",str == in,reflect.DeepEqual(str, in))
	////prints: str == in: true true
	//
	//v1 := []string{"one","two"}
	//v2 := []interface{}{"one","two"}
	//fmt.Println("v1 == v2:",reflect.DeepEqual(v1, v2))
	////prints: v1 == v2: false (not ok)
	//
	//data := map[string]interface{}{
	//	"code": 200,
	//	"value": []string{"one","two"},
	//}
	//encoded, _ := json.Marshal(data)
	//var decoded map[string]interface{}
	//json.Unmarshal(encoded, &decoded)
	//fmt.Println(decoded)
	//fmt.Println("data == decoded:",reflect.DeepEqual(data, decoded))
	//v1 := datass{}
	//v2 := datass{}
	//fmt.Println("v1 == v2:",v1 == v2)
	//raw := []byte(`{"data":"\\xc2"}`)
	//var decoded config
	//
	//if err := json.Unmarshal(raw, &decoded); err != nil {
	//	fmt.Println(err)
	//	//prints: invalid character 'x' in string escape code
	//}
	//fmt.Println(decoded.Data)
	//raw := []byte(`{"data":"w"}`)
	//var decoded config
	//
	//if err := json.Unmarshal(raw, &decoded); err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("%x\n",decoded.Data)
	//var data = []byte(`{"status": 200}`)
	//
	//var result map[string]interface{}
	//var decoder = json.NewDecoder(bytes.NewReader(data))
	//decoder.UseNumber()
	//
	//if err := decoder.Decode(&result); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//
	//var status,_ = result["status"].(json.Number).Int64() //ok
	//fmt.Println("status value:",status)
	//data := "x < y"
	//
	//raw,_ := json.Marshal(data)
	//fmt.Println(string(raw))
	////prints: "x \u003c y" <- probably not what you expected
	//
	//var b1 bytes.Buffer
	//json.NewEncoder(&b1).Encode(data)
	//fmt.Println(b1.String())
	////prints: "x \u003c y" <- probably not what you expected
	//
	//var b2 bytes.Buffer
	//enc := json.NewEncoder(&b2)
	//enc.SetEscapeHTML(false)
	//enc.Encode(data)
	//fmt.Println(b2.String())
	//data := map[string]int{"key": 1}
	//
	//var b bytes.Buffer
	//json.NewEncoder(&b).Encode(data)
	//
	//raw,_ := json.Marshal(data)
	//
	//if b.String() == string(raw) {
	//	fmt.Println("same encoded data")
	//} else {
	//	fmt.Printf("'%s' != '%s'\n",raw,b.String())
	//	//prints:
	//	//'{"key":1}' != '{"key":1}\n'
	//}
	//key := "key.1"
	//d := data{1,&key,make(map[string]bool)}
	//
	//fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	////prints num=1 key=key.1 items=map[]
	//
	//d.pmethod()
	//fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	////prints num=7 key=key.1 items=map[]
	//
	//d.vmethod()
	//fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//data:=make(chan string)
	//done:=make(chan struct{})
	//for i:=0;i<3;i++{
	//	temp:=fmt.Sprintf("%d",i)
	//	go func(i string){
	//		select{
	//			case data<-i:
	//				fmt.Println("put:",i)
	//			case <-done:
	//				fmt.Println("done:",i)
	//		}
	//	}(temp)
	//}
	//fmt.Println(<-data)
	//close(done)
	//done<- "1"
	//done<- "2"
	//close(done)
	fmt.Println("yy")
	//time.Sleep(time.Second * 3)
	//<-done
	//fmt.Println("xx")
	//var wg sync.WaitGroup
	//done := make(chan struct{})
	//wq := make(chan interface{})
	//for i:=0;i<2;i++{
	//	wg.Add(1)
	//	go runwait(i,wq,done,&wg)
	//}
	//for i := 0; i < 2; i++ {
	//	wq <- i
	//}
	//
	//close(done)
	//wg.Wait()
	//fmt.Println("all done")
	//fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n",0x2 & 0x2 + 0x4)
	//var d uint8 = 2
	//fmt.Printf("%08b\n",^d)
	//m := map[string]int{"one":1,"two":2,"three":3,"four":4}
	//
	//for k,v := range m {
	//	fmt.Println(k,v)
	//}
	//data := "A\xfe\x02\xff\x04"
	//for _,v := range data {
	//	fmt.Printf("%#x ",v)
	//}
	////prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)
	//
	//fmt.Println()
	//for _,v := range []byte(data) {
	//	fmt.Printf("%#x ",v)
	//}
	//data := "é"
	//fmt.Println(len(data))                    //prints: 3
	//fmt.Println(utf8.RuneCountInString(data))
	//user:=map[string]int{
	//	"a":55,
	//	"b":11,
	//	"c":33,
	//	"d":88,
	//	"e":6,
	//}
	//userArray:=make([]string,0)
	//for k,_:=range user{
	//	userArray=append(userArray,k)
	//}
	//offsetUser:=make(map[string]int)
	//sum:=0
	//for _,v:=range userArray{
	//	sum+=user[v]
	//	offsetUser[v]=sum
	//}
	//for k,v:=range offsetUser{
	//	fmt.Println(k,":",v)
	//}
	//stat:=make(map[string]int)
	//for i:=0;i<1000;i++{
	//	us:=getRewardUser(userArray,sum,offsetUser)
	//	stat[us]++
	//}
	//for k,v:=range stat{
	//	fmt.Println(k,":",v)
	//}
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
	//time.Sleep(time.Second * 3)
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
