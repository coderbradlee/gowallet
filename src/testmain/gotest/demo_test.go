package gotest

import (
	"container/heap"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"sort"

	//"runtime"

	//"math/rand"
	"sync/atomic"
	//"os"
	"testing"
	"time"
	//"sys"
	//"sync"
	"context"
)

type KeyNotFoundError struct {
	Name string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Errorf("taildb: key %q not found", e.Name).Error()
}
func retError() error {
	return KeyNotFoundError{"test"}
}
func passArrayByValue(arr [3]int) {
	arr = [3]int{}
}
func TestAll9(t *testing.T) {
	arr := [3]int{1, 2, 3}
	passArrayByValue(arr)
	fmt.Println(arr)
	//golang_arr := []int{1, 2, 3, 4, 5}
	//fmt.Println(golang_arr)
	//golang_arr2 := []int{4: -1}
	//fmt.Println(golang_arr2)
	//fmt.Println(retError())
	//arr := make([]int, 0)
	//arr = append(arr, 1, 2, 3, 4)
	//arr1 := arr
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//arr[0] = 10
	//fmt.Println("///////////////")
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//funcarr := func(arr []int) (ret []int) {
	//	arr[0] = 100
	//	return arr
	//}
	//arr2 := funcarr(arr[:2])
	//fmt.Println("///////////////")
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println(arr2)

}
func TestAll8(t *testing.T) {

	for i := 0; i < 100; i++ {
		actionMap := make(map[string][]string)
		actionMap["tt"] = []string{"asfdafd"}
		fmt.Println(len(actionMap["tt"]))
	}
}
func TestAll7(t *testing.T) {
	retBytes, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000019000000000000000000000000000000000000000000000000000000000000001f4")
	if err != nil {
		fmt.Println(err)
		return
	}
	len := len(retBytes) / 32
	fmt.Println("/////////////////////", len)

	retBigs := []*big.Int{}
	for i := 2; i < len; i++ {
		b := retBytes[i*32 : (i+1)*32]
		fmt.Println(b)
		retBig := new(big.Int).SetBytes(b)
		retBigs = append(retBigs, retBig)
	}
	fmt.Println(("//////////////////////"), retBigs)
}

func TestAll6(t *testing.T) {
	//ret := math.Log(1.2) * 100
	//fmt.Println(ret)
	remaintime := 24*14*time.Hour - time.Hour*24
	cel := math.Ceil(remaintime.Seconds() / 86400)
	fmt.Println(cel)
	weight := float64(1)
	weight += math.Log(math.Ceil(remaintime.Seconds()/86400)) / math.Log(1.2) / 100
	fmt.Println(weight)
	amount := new(big.Float).SetInt(big.NewInt(3000000))
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	fmt.Println(weightedAmount)
}
func TestAll5(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func TestAll4(ti *testing.T) {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
}
func TestAll3(t *testing.T) {
	var n int = 10
	var chs = make([]chan int, n)

	var worker = func(n int, c chan int) {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}

	//不定数量的channel数组
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go worker(3+i, chs[i])
	}

	var selectCase = make([]reflect.SelectCase, n)
	//将channel绑定到SelectCase
	for i := 0; i < n; i++ {
		selectCase[i].Dir = reflect.SelectRecv //设置信道是接收,可以为下面值之一
		/*
			const (
				SelectSend    SelectDir // case Chan <- Send
				SelectRecv              // case <-Chan:
				SelectDefault           // default
			)
		*/
		selectCase[i].Chan = reflect.ValueOf(chs[i])
	}

	numDone := 0
	//从所有channel中取出最先到达的N个值
	for numDone < n {
		chosen, recv, recvOk := reflect.Select(selectCase)
		if recvOk {
			fmt.Println(chosen, recv.Int(), recvOk)
			numDone++
		} else {
			fmt.Println("recv error")
		}
	}
}
func TestAll2(ti *testing.T) {
	//runtime.GOMAXPROCS(1)
	//var quit chan int = make(chan int)
	//
	//loop := func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Printf("%d \n", i)
	//	}
	//	quit <- 0
	//}
	//go loop()
	//go loop()
	//
	//for i := 0; i < 2; i++ {
	//	<-quit
	//}
	//1000 0000 0000 0000
	//num := 1<<31 - 1
	//num32 := int32(num)
	//fmt.Println(num32)
	//n := runtime.GOMAXPROCS(1)
	//fmt.Println(n)
	//
	//n = runtime.GOMAXPROCS(n)
	//fmt.Println(n)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(time.Second)
	//		fmt.Println(i)
	//	}
	//}()
	//go func() {
	//	for i := 100; i > 20; i-- {
	//		time.Sleep(time.Second * 2)
	//		fmt.Println(i)
	//	}
	//}()
	//num := (30 + ^uint(0)>>63)
	//fmt.Println(num)
	//old := debug.SetMaxThreads(1 << (num + 10))
	//fmt.Println(old)
	//old = debug.SetMaxThreads(old)
	//fmt.Println(old)
	//var tchan chan int
	//chanVal := reflect.ValueOf(interface{}(tchan))
	//fmt.Println(chanVal.Type().ChanDir())
	//fmt.Println(reflect.SendDir)
	//if chanVal.Kind() != reflect.Chan || chanVal.Type().ChanDir()&reflect.SendDir == 0 {
	//	fmt.Println("first argument to Subscribe must be a writable channel")
	//}
	//if chanVal.IsNil() {
	//	fmt.Println("channel given to Subscribe must not be nil")
	//}
	//var t test
	//ty := reflect.TypeOf(interface{}(&t))
	//fmt.Println(ty)
	////tvalue := reflect.ValueOf(interface{}(t))
	//fmt.Println(ty.MethodByName("Te"))
	//method := ty.Method(0)
	//mtype := method.Name
	//fmt.Println(mtype)
	//fmt.Println(new(big.Int))
	//threshold := new(big.Int).Div(new(big.Int).Mul(big.NewInt(2), big.NewInt(100+int64(100))), big.NewInt(100))
	//fmt.Println(threshold)
	//testHeap()
	//type node []byte
	//n := node{1, 2, 3}
	//testbytes(n)
	//m := make(map[int]int, 10)
	//fmt.Println(len(m))
	//b := make([]int, 10, 0)
	//fmt.Println(cap(b))
	//fmt.Println(len(b))
	//m := make(map[string]string)
	//m["0"] = "0s"
	//fmt.Println(cap(m))//nop cap only have len
	//testContext4()
	//var buf bytes.Buffer
	//p := []byte{'s', 'q', 't'}
	//fmt.Println(p)
	//buf.Write(p)
	//fmt.Println(buf.String())
	//p[0] = 'q'
	//fmt.Println(p)
	//fmt.Println(buf.String())
	//type xx struct {
	//	YY []int `json:"yy"`
	//}
	//j, err := json.Marshal(&xx{YY: []int{}})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(j))
	//fmt.Println((^uintptr(0) >> 63))
	//fmt.Println(3 << 1)
	//fmt.Println(4 << (^uintptr(0) >> 63))
	//test := unsafe.Offsetof(struct {
	//	b uint8
	//	c uint64
	//	v int64
	//}{}.v)
	//fmt.Println(test)
	//v := make(map[string]int)
	//byte
	//v["ss"] = 1
	//debug.SetGCPercent(-1)
	//var count int32
	//newFunc := func() interface{} {
	//	return atomic.AddInt32(&count, 1)
	//}
	//pool := sync.Pool{New: newFunc}
	//v1 := pool.Get()
	//fmt.Println(v1)
	//pool.Put(10)
	//pool.Put(11)
	//pool.Put(12)
	//pool.Put(13)
	//v2 := pool.Get()
	//fmt.Println(v2)
	//fmt.Println(pool.Get())
	//debug.SetGCPercent(100)
	//runtime.GC()
	//fmt.Println(pool.Get())
	//pool.New = nil
	//fmt.Println(pool.Get())
	//var v atomic.Value
	//v.Store(10)
	//testAtomicValue(&v)
	//xx := v.Load()
	//fmt.Println(xx)
	//x := v.Load()
	//fmt.Println(x)
	//v.Store(100)
	//y := v.Load()
	//fmt.Println(y)
	//v.Store(10.1)
	//num := 10
	//i := unsafe.Pointer(&num)
	//fmt.Println(i)
	//delta := uintptr(16) + uintptr(i)
	//
	//fmt.Println(unsafe.Pointer(delta))
	//j := atomic.LoadPointer(&i)
	//fmt.Println(j)
	//var x struct {
	//	before uintptr
	//	i      uintptr
	//	after  uintptr
	//}
	//var m uint64 = 0xdeddeadbeefbeef
	//magicptr := uintptr(m)
	//x.before = magicptr
	//x.after = magicptr
	//var j uintptr
	//for delta := uintptr(1); delta+delta > delta; delta += delta {
	//	k := atomic.SwapUintptr(&x.i, delta)
	//	if x.i != delta || k != j {
	//		ti.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
	//	}
	//	j = delta
	//}
	//if x.before != magicptr || x.after != magicptr {
	//	ti.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
	//}
	//ch := make(chan int, 1)
	//ticker := time.NewTicker(time.Second)
	//go func() {
	//
	//	for t := range ticker.C {
	//		fmt.Println(t)
	//		select {
	//		case ch <- 1:
	//		case ch <- 2:
	//		case ch <- 3:
	//		}
	//	}
	//	fmt.Println("xx")
	//}()
	//go func() {
	//	timer := time.After(time.Second * 5)
	//loop:
	//	for {
	//		select {
	//		case i := <-ch:
	//			fmt.Println("received:", i)
	//		case <-timer:
	//			ticker.Stop()
	//			fmt.Println("its time")
	//			break loop
	//		}
	//	}
	//}()

	time.Sleep(time.Second * 30)
	//var stopPoint chan struct {}
	//<-stopPoint
	//ch1:=make(chan int,1)
	//ch2:=make(chan int,2)
	//chs:=[]chan int{ch1,ch2}
	//nums:=[]int{1,2,3,4}
	//select {
	//	case getChan(chs,0)<-getNumber(nums,0):
	//		fmt.Println("1")
	//	case getChan(chs,1)<-getNumber(nums,1):
	//		fmt.Println("2")
	//default:
	//	fmt.Println("default")
	//}
	//ch:=make(chan<- int,1)
	//_,ok:=interface{}(ch).(chan int)
	//fmt.Println(ok)
	//
	//_,ok=interface{}(ch).(<-chan int)
	//fmt.Println(ok)
	//close(ch)
	//ch=nil
	//close(ch)
	//fmt.Println(^(-200))
	//x-200=x+^(-200)
	//var ch2 chan int
	////close(ch2)
	//fmt.Println(len(ch2),":",cap(ch2))
	//wg:=sync.WaitGroup{}
	//wg.Add(2)
	//type counter struct {
	//	count int
	//}
	//ch:=make(chan map[string]*counter,3)
	//fmt.Println(len(ch),":",cap(ch))
	//go func() {
	//	sendmap:=map[string]*counter{
	//		"count":&counter{},
	//	}
	//	for i:=0;i<5;i++{
	//		ch<-sendmap
	//	}
	//	close(ch)
	//	fmt.Println("close chan")
	//	wg.Done()
	//}()
	//go func() {
	//	time.Sleep(time.Second*4)
	//	for re:=range ch{
	//		c:=re["count"]
	//		fmt.Println(c.count)
	//		c.count++
	//	}
	//	fmt.Print("stop read")
	//	wg.Done()
	//}()
	//wg.Wait()

	//from:=make(chan int,1)
	//from<-1+2
	//out:=<-from
	//fmt.Print(out)
	//s := []string{"a", "b", "c", "d"}
	//defer fmt.Println(s) // [a x y d]
	//// defer append(s[:1], "x", "y") // error
	//defer func() {
	//	_=append(s[:1], "x", "y")
	//}()
	//s=append(s,"xxxxxx")
	//fmt.Println(s)
	//names:=[]string{"1","2","3","4"}
	//for _,v:=range names{
	//	defer func(name string) {
	//		fmt.Println(name)
	//	}(v)
	//}
	//time.Sleep(time.Second)
	//name:="start"
	//defer func(){
	//	fmt.Println(name)
	//}()
	//name="end"

	//sig:=make(chan os.Signal,1)
	//sigs:=[]os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	//signal.Notify(sig,sigs...)
	//for s:=range sig{
	//	fmt.Print(s)
	//}
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
	//	time.Sleep(time.Second)
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

func getNumber(nums []int, index int) int {
	fmt.Println("getnumber:", index)
	return nums[index]
}
func getChan(chs []chan int, index int) chan int {
	fmt.Println("getchan:", index)
	return chs[index]
}
func testAtomicValue(v *atomic.Value) {
	v.Store(20)
}
func testContext() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case x := <-ctx.Done(): //只有撤销函数被调用后，才会触发
					fmt.Println(x, " done")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //调用返回的cancel方法来让 context声明周期结束

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
func testContext2() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel() //时间超时会自动调用

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
func testContext3() {
	type favContextKey string

	f := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)

	}

	//k := favContextKey("language")
	//k1 := favContextKey("Chinese")
	k := "language"
	k1 := "Chinese"
	ctx := context.WithValue(context.Background(), k, "Go")
	ctx1 := context.WithValue(ctx, k1, "Go1")

	f(ctx1, k1)
	f(ctx1, k)

}
func testContext4() {
	ctx, cancel := context.WithCancel(context.Background())
	f := func(name string, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, " quit")
				return
			case <-time.After(time.Second * 2):
				fmt.Println(name)
			}
		}
	}
	go f("1", ctx)
	<-time.After(time.Second * 10)
	fmt.Println("main quit:")
	cancel()
	<-time.After(time.Second * 1)
}
func testbytes(thisok []byte) {
	fmt.Println(thisok)
}

type nonceHeap []uint64

func (h nonceHeap) Len() int           { return len(h) }
func (h nonceHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h nonceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nonceHeap) Push(x interface{}) {
	*h = append(*h, x.(uint64))
}

func (h *nonceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func testHeap() {
	var index = &nonceHeap{}

	for i := 10; i > 0; i-- {
		//nonce := uint64(rand.Intn(10))
		//fmt.Println(nonce)
		//heap.Push(index, nonce)
		*index = append(*index, uint64(i))
	}

	fmt.Println(index)
	//nonce := heap.Pop(index).(uint64)
	//fmt.Println("poped nonce:", nonce)
	//fmt.Println(index)
	heap.Init(index)
	fmt.Println(index)
	//sort.Sort(*index)
	//fmt.Println(index)

	/////iterator
	fmt.Println("////////////////////////")
	for next := (*index)[0]; index.Len() > 0 && (*index)[0] == next; next++ {
		//fmt.Println(index.Len())
		fmt.Println(next)
		ret := heap.Pop(index).(uint64)
		fmt.Println("poped:", ret)
		fmt.Println(index)
	}
}

type test struct {
	xx int
}

func (t test) Te() {

}
