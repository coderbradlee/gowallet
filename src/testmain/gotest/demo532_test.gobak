package gotest

import (
	"context"
	"errors"
	"fmt"
	// "github.com/prashantv/gostub"
	"io"
	"os"
	// "reflect"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

type student struct {
	Name string
	Age  int
}
type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i

	defer func() {
		fmt.Println("def:", t)
		t += 3
		fmt.Println("def:", t)
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1
type MyStruct struct {
	T1
	T2
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
func testfunc() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}
func testfunc2(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

// const (
// 	ERR_ELEM_EXIST error = errors.New("element already exists")

// 	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
// )
func incr(p *int) int {

	*p++

	return *p

}

type Slice []int

func NewSlice() Slice {

	return make(Slice, 0)

}

func (s *Slice) Add(elem int) *Slice {

	*s = append(*s, elem)

	fmt.Print(elem)

	return s

}
func TestTimeer(ti *testing.T) {
	query := map[string]string{}

	query["test0"] = "0"
	query["test1"] = "1"
	query["test2"] = "2"

	i := 0
	for k, v := range query {
		delete(query, "test2")
		fmt.Println(query, k, v)
		i++
	}
	// m := make(map[string]int)
	// m["1"] = 1
	// m["2"] = 2
	// m["3"] = 3
	// m["4"] = 4
	// fmt.Println(len(m))
	// for k, _ := range m {
	// 	delete(m, k)
	// }
	// fmt.Println(len(m))
	// fmt.Println(m["1"])
	// s := NewSlice()
	// defer s.Add(1).Add(2)
	// s.Add(3)
	// v := 1
	// incr(&v)
	// fmt.Println(v)
	// for i := 0; i < 5; i++ {

	// 	defer fmt.Printf("%d ", i)

	// }
	// var x chan int
	// fmt.Println(cap(x))
	// var m map[string]int

	// m = make(map[string]int)

	// m["one"] = 1
	// fmt.Println(m)
	// var s []int

	// s = append(s, 1)
	// fmt.Println(s)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("++++")
	// 		f := err.(func() string)
	// 		fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
	// 	} else {
	// 		fmt.Println("fatal")
	// 	}
	// }()

	// defer func() {
	// 	panic(func() string {
	// 		return "defer panic"
	// 	})
	// }()
	// panic("panic")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("fatal")
	// 	}
	// }()

	// defer func() {
	// 	panic("defer panic")
	// }()
	// panic("panic")
	// a, b := testfunc2(100)
	// a()
	// b()
	// funs := testfunc()
	// for _, f := range funs {
	// 	f()
	// }
	// fmt.Println(DoTheThing(true))
	// fmt.Println(DoTheThing(false))
	// my := MyStruct{}
	// my.T1.m1()
	// my.T2.m1()
	// type MyInt1 int
	// type MyInt2 = int
	// var i int = 9
	// var i1 MyInt1 = i
	// var i2 MyInt2 = i
	// fmt.Println(i1, i2)
	// fmt.Println(x, y, z, k, p)
	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// // if sm1 == sm2 {
	// // 	fmt.Println("sm1 == sm2")
	// // }
	// if reflect.DeepEqual(sm1, sm2) {
	// 	fmt.Println("sn1 ==sm")
	// } else {
	// 	fmt.Println("sn1 !=sm")
	// }
	// sn1 := struct {
	// 	age  int
	// 	name string
	// }{age: 11, name: "qq"}
	// sn2 := struct {
	// 	age  int
	// 	name string
	// }{age: 11, name: "qq"}

	// if sn1 == sn2 {
	// 	fmt.Println("sn1 == sn2")
	// }
	// println(DeferFunc1(1))
	// println(DeferFunc2(1))
	// println(DeferFunc3(1))
	// var peo People = &Stduent{}
	// think := "bitch"
	// fmt.Println(peo.Speak(think))
	// s := make([]int, 5)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)
	// s := make([]int, 0)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s) //[1 2 3]
	// runtime.GOMAXPROCS(1)
	// wg := sync.WaitGroup{}
	// wg.Add(20)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println("A: ", i)
	// 		wg.Done()
	// 	}()
	// }
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("B: ", i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// m := make(map[string]*student)
	// stus := []student{
	// 	{Name: "zhou", Age: 24},
	// 	{Name: "li", Age: 23},
	// 	{Name: "wang", Age: 22},
	// }
	// for _, stu := range stus {
	// 	// m[stu.Name] = &stu
	// 	stu.Age = stu.Age + 10
	// }
	// for i := 0; i < len(stus); i++ {
	// 	stus[i].Age += 10
	// }
	// for k, v := range stus {
	// 	fmt.Println(k, "=>", v)
	// }

	// for i := 0; i < len(stus); i++ {
	// 	m[stus[i].Name] = &stus[i]
	// }
	// for k, v := range m {
	// 	println(k, "=>", v.Name)
	// }
	// defer func() { fmt.Println("打印前") }()
	// defer func() { fmt.Println("打印中") }()
	// defer func() { fmt.Println("打印后") }()

	// panic("触发异常")
	// t := time.NewTimer(0)
	// time.Sleep(time.Second)
	// if !t.Stop() {
	// 	fmt.Println("in1")
	// 	<-t.C
	// 	fmt.Println("in2")
	// }
	// t := time.AfterFunc(1*time.Second, func() {
	// 	fmt.Println("Time has passed!")
	// })
	// // This will deadlock.
	// time.Sleep(time.Second * 3)
	// if !t.Stop() {
	// 	fmt.Println("in1")
	// 	<-t.C
	// 	fmt.Println("in2")
	// }

	// done := make(chan struct{})
	// f := func() {
	// 	fmt.Println("func1")
	// 	// time.Sleep(time.Second * 3)
	// 	close(done)
	// 	fmt.Println("func2")
	// }
	// t := time.AfterFunc(1*time.Second, f)
	// time.Sleep(time.Second * 2)
	// if !t.Stop() {
	// 	fmt.Println("in")
	// 	<-done
	// 	fmt.Println("in2")
	// }
	// time.Sleep(time.Second * 20)
	// ti := time.NewTimer(1)
	// // <-ti.C

	// time.Sleep(time.Second)
	// fmt.Println(time.Now())
	// ti.Reset(1)
	// // <-ti.C
	// fmt.Println(time.Now())
	// out := <-ti.C
	// fmt.Println("out:", out)
	// if !ti.Stop() {
	// 	out := <-ti.C
	// 	fmt.Println("in:", out)
	// }
	// fmt.Println("end")
}
func TestTimerFiresAfterStop(t *testing.T) {
	fail := make(chan struct{})
	done := make(chan struct{})
	defer close(done)
	for i := 0; i < 1000; i++ {
		tr := time.NewTimer(0)
		tr.Stop()
		// There may or may not be a value in the channel now. But definitely
		// one should not be added after we receive it.
		select {
		case <-tr.C:
			fmt.Println(i, ":")
		default:
			// fmt.Println(i, ":default")
		}
		// Now set the timer to trigger in hour. It definitely shouldn't be
		// receivable now for an hour.
		// tr.Reset(time.Hour)
		go func(iter int) {
			select {
			case <-tr.C:
				// As soon as the channel receives, notify failure.
				fail <- struct{}{}
				fmt.Println("x")
			case <-done:
				// fmt.Println(iter, ":done")
			}
		}(i)
	}
	select {
	case <-fail:
		t.FailNow()
	case <-time.After(time.Second):
	}
}
func TestString(t *testing.T) {
	fmt.Println(time.Now())
	tchan := time.After(time.Second)
	ti := <-tchan
	fmt.Println(ti)
	// tm := time.NewTimer(1)
	// // tm.Reset(1000 * time.Millisecond)
	// fmt.Println("21")
	// // <-tm.C
	// fmt.Println("22")
	// ret := tm.Stop()
	// fmt.Println(ret)
	// if !tm.Stop() {
	// 	fmt.Println("xxxx")
	// 	<-tm.C
	// 	fmt.Println("yyyy")
	// }
	// fmt.Println("23")
	// testhttp()
	// testrpc()

	// stubs := gostub.Stub(&num, 150)
	// defer stubs.Reset()
	// fmt.Println(num)
	// fmt.Println(stubs.num)

	// testnet()
	// fmt.Println(runtime.NumCPU())
	// ch := make(chan int, 1)
	// go func() {
	// 	for i := range ch {
	// 		fmt.Println(":", i)
	// 	}
	// 	// for {
	// 	// 	fmt.Println("flag3")
	// 	// 	i := <-ch
	// 	// 	fmt.Println(":", i)
	// 	// }
	// }()
	// for {
	// 	fmt.Println("flag1")
	// 	select {
	// 	case ch <- 0:
	// 		fmt.Println("0")
	// 	case ch <- 1:
	// 		fmt.Println("1")
	// 	}
	// 	fmt.Println("flag2")
	// 	time.Sleep(time.Second)
	// }

	// ch := make([]chan int, 10)
	// ch[9] = make(chan int)
	// go func() {
	// 	for i, v := range ch {
	// 		<-v
	// 		fmt.Println(i, ":")
	// 	}
	// }()
	// for i := 0; i < 9; i++ {
	// 	ch[i] = make(chan int)
	// 	go testchan(ch[i])
	// 	time.Sleep(time.Second * 2)
	// }

	// var y testinterface
	// x := &teststruct{7}
	// y = x
	// // z := &usestruct{x, 1}
	// z := &usestruct{x}
	// z.xx.add()
	// // fmt.Println(z.xx.x)
	// x.add()
	// fmt.Println(x.x)
	// if v, ok := y.(*teststruct); ok {
	// 	fmt.Println(v.x)
	// }
	// switch y.(type) {
	// case testinterface:
	// 	fmt.Println("y.(type)")
	// default:
	// 	fmt.Println("default")
	// }
	// switch {
	// case 0 > 0:
	// 	fmt.Println("0")
	// case 1 < 0:
	// 	fmt.Println("1")
	// case 2 > 0:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }
	// ret := returefunc()
	// fmt.Println(ret)
	// sli1 := []int{1, 2, 3, 4, 5}
	// sli2 := []int{6, 7, 8}
	// // copy(sli1, sli2)
	// copy(sli2, sli1)
	// fmt.Println(sli1)
	// fmt.Println(sli2)

	// sli := []int{1, 2, 3, 4}
	// fmt.Println(len(sli), ":", cap(sli))
	// sli = append(sli, 5, 6, 7)
	// fmt.Println(len(sli), ":", cap(sli))
	// ano := sli[:9]
	// fmt.Println(len(ano), ":", cap(ano))
	// arr := [5]int{1, 2, 3, 4, 5}
	// modify := func(arr [5]int) {
	// 	arr[0] = 1000
	// 	fmt.Println(arr)
	// }
	// modify(arr)
	// fmt.Println(arr)

	// str := "Hello, 世界"
	// fmt.Println(len(str))
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(i, ":", str[i])
	// }
	// for i := range str {
	// 	fmt.Println(i)
	// }
}
func TestProcessor(t *testing.T) {
	p, err := os.FindProcess(7036)
	if err != nil {
		fmt.Println(err)
	} else {
		p.Kill()
	}
}
func TestFile(t *testing.T) {
	f := os.NewFile(uintptr(syscall.Stderr), "test")
	f.WriteString("xxxx")
}
func TestBuilder(t *testing.T) {
	src := strings.NewReader("改革开放是我们党的一次伟大觉醒。在改革开放进程中，我们党不断开辟马克思主义发展新境界，创造人类历史上前所未有的发展奇迹。")
	dst := new(strings.Builder)
	writenlen, err := io.CopyN(dst, src, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(writenlen)
		fmt.Println(dst.String())
	}
}
func TestUtf(t *testing.T) {
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()

	// 示例2。
	fmt.Println("Grow the builder ...")
	builder1.Grow(10)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()

	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())
	// str := "Go爱好者"
	// for i, c := range str {
	// 	fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	// }
	// 	s := "test测试"
	// 	fmt.Printf("%s\n", s)
	// 	fmt.Printf("%q\n", []rune(s))
	// 	fmt.Printf("%x\n", []rune(s))
	// 	fmt.Printf("%x\n", []byte(s))
	// }
}
func TestContext2(t *testing.T) {
	ch := make(chan int, 1)
	// ch = nil
	// close(ch)

	go func(ch chan int) {
		// <-ch
		ch <- 100
		fmt.Println("xx")
	}(ch)
	time.Sleep(2 * time.Second)
	close(ch)
	xx := <-ch
	fmt.Println(xx)
	// close(ch)
	xx = <-ch
	fmt.Println(xx)
	// <-ch
	// // ch <- struct{}{}
	// // close(ch)
	// // <-ch
	// // ch <- struct{}{}
	// fmt.Println("yy")
	// ch = nil
	// ch <- struct{}{}
	// <-ch
	// fmt.Println("zz")
	// someHandler()
}
func TestContext(t *testing.T) {
	type myKey int
	keys := []myKey{
		myKey(0),
		myKey(1),
		myKey(2),
		myKey(3),
	}
	values := []string{
		"value 0",
		"value 1",
		"value 2",
		"value 3",
	}

	rootNode := context.Background()
	node1, cancelFunc1 := context.WithCancel(rootNode)
	// defer cancelFunc1()
	cancelFunc1()
	fmt.Println(":", node1.Err())
	<-node1.Done()

	// 示例1。
	node2 := context.WithValue(node1, keys[0], values[0])
	node3 := context.WithValue(node2, keys[1], values[1])
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[0], node3.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[1], node3.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[2], node3.Value(keys[2]))
	fmt.Println()
	<-node2.Done()
	<-node3.Done()
	// 示例2。
	node4, _ := context.WithCancel(node3)
	node5, _ := context.WithTimeout(node4, time.Hour)
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[0], node5.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[1], node5.Value(keys[1]))
	fmt.Println()

	// 示例3。
	node6 := context.WithValue(node5, keys[2], values[2])
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[0], node6.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[2], node6.Value(keys[2]))
	fmt.Println()

	// 示例4。
	node6Branch := context.WithValue(node5, keys[3], values[3])
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[1], node6Branch.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[2], node6Branch.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[3], node6Branch.Value(keys[3]))
	fmt.Println()

	// 示例5。
	node7, _ := context.WithCancel(node6)
	node8, _ := context.WithTimeout(node7, time.Hour)
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[1], node8.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[2], node8.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[3], node8.Value(keys[3]))
}
func TestReuse(t *testing.T) {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go showNum(i+j, wg.Done)
		}
		wg.Wait()
		fmt.Println("end:", i)
	}
	fmt.Println("End.")
}
func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	once := sync.Once{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("fatal error: %v\n", p)
			}
		}()
		once.Do(func() {
			fmt.Println("Do task. [4]")
			panic(errors.New("something wrong"))
			fmt.Println("Done. [4]")
		})
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [5]")
		})
		fmt.Println("Done. [5]")
	}()
	wg.Wait()
}
func TestNoRaceWaitGroupTransitive(t *testing.T) {
	x, y := 0, 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		x = 42
		wg.Done()
	}()
	go func() {
		time.Sleep(1e17)
		y = 42
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
	fmt.Println(y)
}
func TestNoRaceWaitGroupPanicRecover(t *testing.T) {
	var xx int
	xx = 0
	fmt.Println(xx)
	var wg sync.WaitGroup
	defer func() {
		err := recover()
		if err != "sync: negative WaitGroup counter" {
			t.Fatalf("Unexpected panic: %#v", err)
		}
		xx = 2
	}()
	xx = 1
	wg.Add(-1)
}
func TestNoRaceWaitGroupMultipleWait(t *testing.T) {
	c := make(chan bool, 2)
	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		c <- true
	}()
	go func() {
		wg.Wait()
		c <- true
	}()
	wg.Wait()
	<-c
	<-c
	fmt.Println("xx")
}
func TestRaceWaitGroupWrongAdd(t *testing.T) {
	c := make(chan bool, 2)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		time.Sleep(1000 * time.Millisecond)
		wg.Done()
		c <- true
	}()
	go func() {
		time.Sleep(1000 * time.Millisecond)
		wg.Add(1)

		wg.Done()
		c <- true
	}()
	time.Sleep(50 * time.Millisecond)
	wg.Wait()
	<-c
	<-c
	fmt.Println("xx")
}
func TestRaceWaitGroupWrongWait(t *testing.T) {
	c := make(chan bool, 2)
	var x int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// wg.Add(1)
		runtime.Gosched()
		x = 1
		wg.Done()
		c <- true
	}()
	go func() {
		// wg.Add(1)
		runtime.Gosched()
		x = 2
		wg.Done()
		c <- true
	}()
	wg.Wait()
	<-c
	<-c
	fmt.Println(x)
}
func TestRaceWaitGroupAsMutex(t *testing.T) {
	var x int
	var wg sync.WaitGroup
	c := make(chan bool, 2)
	go func() {
		wg.Wait()
		time.Sleep(100 * time.Millisecond)
		wg.Add(+1)
		x = 1
		wg.Add(-1)
		c <- true
	}()
	go func() {
		wg.Wait()
		time.Sleep(100 * time.Millisecond)
		wg.Add(+1)
		x = 2
		wg.Add(-1)
		c <- true
	}()
	<-c
	<-c
	fmt.Println(x)
}
func TestRaceWaitGroup(t *testing.T) {
	var x int
	var wg sync.WaitGroup
	n := 10000
	for i := 0; i < n; i++ {
		wg.Add(1)
		j := i
		go func() {
			x = j
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
}
func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
			name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil, but it should be. (name=%q)",
			name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting, but it should not be. (name=%q)",
			name)
	}
	expected := fmt.Sprintf("Hello, %s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
			greeting, name)
	}
	t.Logf("The expected greeting is %q.\n", expected)
}

func TestIntroduce(t *testing.T) {
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

func TestFail(t *testing.T) {
	t.Fail()
	// t.FailNow() // 此调用会让当前的测试立即失败。
	t.Log("Failed.")
}
func TestCC(t *testing.T) {
	cond()
}
func TestTest(t *testing.T) {
	test()
}

func BenchmarkGetPrimes(b *testing.B) {
	b.StopTimer()
	time.Sleep(time.Millisecond * 500) // 模拟某个耗时但与被测程序关系不大的操作。
	max := 1000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GetPrimes(max)
	}
}
