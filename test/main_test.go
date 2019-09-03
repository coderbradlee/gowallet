package test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"golang.org/x/net/ipv4"
)

func testDijkstra() {
	var name = make(map[string]int)
	name["s"] = 0
	name["a"] = 1
	name["b"] = 2
	name["c"] = 3
	name["d"] = 4
	name["e"] = 5
	name["f"] = 6
	name["g"] = 7
	name["h"] = 8
	var index = make(map[int]string)
	index[0] = "s"
	index[1] = "a"
	index[2] = "b"
	index[3] = "c"
	index[4] = "d"
	index[5] = "e"
	index[6] = "f"
	index[7] = "g"
	index[8] = "h"

	var weight [9][9]float64

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i == j {
				weight[i][j] = 0
			} else {
				weight[i][j] = 1000
			}
		}
	}
	weight[name["s"]][name["a"]] = 0.5
	weight[name["s"]][name["b"]] = 0.3
	weight[name["s"]][name["c"]] = 0.2
	weight[name["s"]][name["d"]] = 0.4

	weight[name["a"]][name["e"]] = 0.3

	weight[name["b"]][name["a"]] = 0.2
	weight[name["b"]][name["f"]] = 0.1

	weight[name["c"]][name["f"]] = 0.4
	weight[name["c"]][name["h"]] = 0.8

	weight[name["d"]][name["c"]] = 0.1
	weight[name["d"]][name["h"]] = 0.6

	weight[name["e"]][name["g"]] = 0.1

	weight[name["f"]][name["e"]] = 0.1
	weight[name["f"]][name["h"]] = 0.2

	weight[name["h"]][name["g"]] = 0.4
	//settled := make(map[int]int)
	settledWei := make(map[int]float64)
	mw := make(map[int]float64)
	for i := 0; i < 9; i++ {
		mw[i] = weight[name["s"]][i]
	}
	dijkstra(weight, mw, settledWei)
	//for k, v := range settled {
	//	fmt.Println(k, ":", v)
	//}
	for k, v := range settledWei {
		fmt.Printf("%s:%0.2f ", index[k], v)
	}
	fmt.Println()
}
func dijkstra(weight [9][9]float64, mw map[int]float64, settledWei map[int]float64) {
	//if len(settledWei) >= 9 {
	//	return
	//}
	if len(mw) == 0 {
		return
	}
	minDis := 1000.0
	node := 0
	// 找出距离最小的节点
	for k, v := range mw {
		if v < minDis {
			minDis = v
			node = k
		}
	}
	// update min weight
	for k, v := range mw {
		if v > mw[node]+weight[node][k] {
			mw[k] = mw[node] + weight[node][k]
		}
	}
	settledWei[node] = minDis
	delete(mw, node)

	dijkstra(weight, mw, settledWei)

}

type TokenType string
type Token interface {
	Type() TokenType
	Lexeme() string
}

type Match struct {
	toktype TokenType
	lexeme  string
}

type IntegerConstant struct {
	Match
	value uint64
}

func (m *Match) Type() TokenType {
	return m.toktype
}

func (m *Match) Lexeme() string {
	return m.lexeme
}

func (i *IntegerConstant) Value() uint64 {
	return i.value
}

type Pet struct {
	name string
}

type Dog struct {
	Pet
	Breed string
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak2() string {
	return fmt.Sprintf("%v and I am a %v", d.Speak(), d.Breed)
}
func testudp() {
	go func() {
		time.Sleep(time.Second * 2)
		sip := net.ParseIP("127.0.0.1")
		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
		dstAddr := &net.UDPAddr{IP: sip, Port: 9981}
		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		conn.Write([]byte("hello"))
		fmt.Printf("<%s>\n", conn.RemoteAddr())

		data := make([]byte, 1024)
		n, err := conn.Read(data)
		fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
	}()
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	//data := make([]byte, 1024)
	//remoteAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 19982}
	for {
		//n, err := listener.Read(data)
		//if err != nil {
		//	fmt.Printf("error during read: %s", err)
		//}
		//fmt.Printf(" %s\n", data[:n])
		//n, remoteAddr, err := listener.ReadFromUDP(data)
		//if err != nil {
		//	fmt.Printf("error during read: %s", err)
		//}
		//
		//fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		//_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		//if err != nil {
		//	fmt.Printf(err.Error())
		//}
		_, err = listener.Write([]byte("hello"))
		if err != nil {
			fmt.Println("server write:", err.Error())
			break
		}
	}
}
func testmulti() {
	go func() {
		time.Sleep(time.Second * 3)
		ip := net.ParseIP("224.0.0.250")
		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
		dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		conn.Write([]byte("hello"))
		fmt.Printf("<%s>\n", conn.RemoteAddr())
	}()
	//1. 得到一个interface
	en4, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println(err)
	}
	group := net.IPv4(224, 0, 0, 250)
	//2. bind一个本地地址
	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	//3.
	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(en4, &net.UDPAddr{IP: group}); err != nil {
		fmt.Println(err)
	}
	//4.更多的控制
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		fmt.Println(err)
	}
	//5.接收消息
	b := make([]byte, 1500)
	for {
		n, cm, src, err := p.ReadFrom(b)
		if err != nil {
			fmt.Println(err)
		}
		if cm.Dst.IsMulticast() {
			if cm.Dst.Equal(group) {
				fmt.Printf("received: %s from <%s>\n", b[:n], src)
				n, err = p.WriteTo([]byte("world"), cm, src)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown group")
				continue
			}
		}
	}
}
func testopenfile() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak2())
	fmt.Println(d.name)
	//t := IntegerConstant{Match{"xx", "wizard"}, 2}
	//fmt.Println(t.Type(), t.Lexeme(), t.Value())
	//x := Token(t)
	//fmt.Println(x.Type(), x.Lexeme())

	//db, err := bolt.Open("my.db", 0600, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(db)
	//opt := bolt.Options{
	//	ReadOnly: true,
	//}
	//db, err = bolt.Open("my.db", 0600, &opt)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(db)

	//xx, err := os.OpenFile("./xx", os.O_RDWR|os.O_CREATE, 0666)
	//fmt.Println(xx, ":", err)
	//
	////flag := syscall.LOCK_SH
	//flag := syscall.LOCK_EX
	//
	//// Otherwise attempt to obtain an exclusive lock.
	//err = syscall.Flock(int(xx.Fd()), flag|syscall.LOCK_NB)
	//fmt.Println(err)
	//
	//xx2, err := os.OpenFile("./xx", os.O_RDWR, 0666)
	//fmt.Println(xx2, ":", err)
}
func testbroard() {
	go func() {
		time.Sleep(time.Second * 4)
		ip := net.ParseIP("172.17.255.255")
		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
		dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
		conn, err := net.ListenUDP("udp", srcAddr)
		if err != nil {
			fmt.Println(err)
		}
		n, err := conn.WriteToUDP([]byte("hello"), dstAddr)
		if err != nil {
			fmt.Println(err)
		}
		data := make([]byte, 1024)
		n, _, err = conn.ReadFrom(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("client read %s from <%s>\n", data[:n], conn.RemoteAddr().String())
	}()
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("error during read: ", err)
		}
		fmt.Printf("server read: <%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
func TestXx(t *testing.T) {
	//testDijkstra()
	//testopenfile()
	//testudp()
	testmulti()
}
