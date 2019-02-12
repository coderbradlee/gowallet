package main

import (
	btswallets "btswallet"
	"fmt"
	wallet "hdwallet"
	// "math/big"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	cfg "github.com/ipfs/go-ipfs-config"
	ci "github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"
	// "time"
)

var (
	// imtoken测试用例
	// 助记词 inject kidney empty canal shadow pact comfort wife crush horse wife sketch
	// ipfs QmWqwovhrZBMmo32BzY83ZMEBQaP7YRMqXNmMc8mgrpzs6
	// eth 6031564e7b2f5cc33737807b2e58daff870b590b
	// 私钥 cce64585e3b15a0e4ee601a467e050c9504a0db69a559d7ec416fa25ad3410c2

	// btc 0地址 12z6UzsA3tjpaeuvA2Zr9jwx19Azz74D6g
	// 1地址1962gsZ8PoPUYHneFakkCTrukdFMVQ4i4T
	// 私钥,此私钥与golang生成的不一致，xprv9yrdwPSRnvomqFK4u1y5uW2SaXS2Vnr3pAYTjJjbyRZR8p9BwoadRsCxtgUFdAKeRPbwvGRcCSYMV69nNK4N2kadevJ6L5iQVy1SwGKDTHQ

	// bitcoin的memo 用go包试下
	// 如何利用助记词 对ipfs 生成peerid及私钥
	// ipfs的imtoken使用方式
	imtokenmnemonic = "inject kidney empty canal shadow pact comfort wife crush horse wife sketch"
	ethaddress      = "6031564e7b2f5cc33737807b2e58daff870b590b"
	btcaddress      = "12z6UzsA3tjpaeuvA2Zr9jwx19Azz74D6g"
	btc0            = "xprvA2veQkqHmgXwTSh9pCfyUPo8SEmpqPgqTBhE8KXQLNf76jbZUCAWT7JsyN3iDWfWFJbt3BqeMigLg5hJEkpNm6WvbmXViRC9zyubgR2eF3T"
	btc0address     = "mhW3n3x8rvB5MmPXsbYDyfAGs8mhw9GGaW"
	btc1            = "xprvA2veQkqHmgXwWChJWsD7mXphrhjumvLgF2o82dE9UeLGYcHoTYa8hi7A2ndChg8xbkPTEq7277cPL2qPTFLQH4AhRYT7nLqKVRr2Prwej3z"
	btc1address     = "mobyyve7CppjKQGFy9j82P5Eccr4PxHeqS"
)

func SendInt(ch chan<- int) {
	ch <- 100
}
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
func test() {
	// intChan1 := make(chan int, 3)
	// SendInt(intChan1)
	// xx := <-intChan1
	// fmt.Println(xx)
	ch := getIntChan()
	xx := <-ch
	fmt.Println(xx)
	xx = <-ch
	fmt.Println(xx)
	// var channels = [3]chan int{
	// 	nil,
	// 	make(chan int),
	// 	nil,
	// }
	// for {
	// 	select {
	// 	case channels[0] <- 0:
	// 		fmt.Println("The first candidate case is selected.")
	// 	case channels[1] <- 1:
	// 		fmt.Println("The second candidate case is selected.")
	// 	case channels[2] <- 2:
	// 		fmt.Println("The third candidate case is selected")
	// 	default:
	// 		fmt.Println("No candidate case is selected!")
	// 	}
	// }

	// intChan := make(chan int, 1)
	// // 一秒后关闭通道。
	// time.AfterFunc(time.Second, func() {
	// 	close(intChan)
	// })
	// select {
	// case _, ok := <-intChan:
	// 	if !ok {
	// 		fmt.Println("The candidate case is closed.")
	// 		break
	// 	}
	// 	fmt.Println("The candidate case is selected.")
	// }

	// 准备好几个通道。
	// intChannels := [3]chan int{
	// 	make(chan int, 1),
	// 	make(chan int, 1),
	// 	make(chan int, 1),
	// }
	// // 随机选择一个通道，并向它发送元素值。
	// index := 0
	// intChannels[index] <- index
	// index = 1
	// intChannels[index] <- index
	// index = 2
	// intChannels[index] <- index
	// // 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	// select {
	// case <-intChannels[0]:
	// 	fmt.Println("The first candidate case is selected.")
	// case <-intChannels[1]:
	// 	fmt.Println("The second candidate case is selected.")
	// case elem := <-intChannels[2]:
	// 	fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	// default:
	// 	fmt.Println("No candidate case is selected!")
	// }

}
func test3() {
	{
		s := []int{1, 2, 3, 4, 5}
		s1 := make([]int, 5)
		copy(s1, s)
		fmt.Println(s)
		fmt.Println(s1)
		s[0] = 100
		fmt.Println(s)
		fmt.Println(s1)
	}
	{
		// s5 := make([]int)
		// fmt.Printf("The capacity of s5: %d\n", cap(s5))

		s6 := make([]int, 0)
		fmt.Printf("The capacity of s6: %d\n", cap(s6))
		fmt.Printf("The len of s6: %d\n", len(s6))
		for i := 1; i <= 5; i++ {
			s6 = append(s6, i)
			fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
		}
		fmt.Println()

		// 示例2。
		s7 := make([]int, 1024)
		fmt.Printf("The capacity of s7: %d\n", cap(s7))
		s7e1 := append(s7, make([]int, 200)...)
		fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
		s7e2 := append(s7, make([]int, 400)...)
		fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
		s7e3 := append(s7, make([]int, 600)...)
		fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
		fmt.Println()

		// 示例3。
		s8 := make([]int, 10)
		fmt.Printf("The capacity of s8: %d\n", cap(s8))
		s8a := append(s8, make([]int, 11)...)
		fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
		s8b := append(s8a, make([]int, 23)...)
		fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
		s8c := append(s8b, make([]int, 45)...)
		fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))

	}
	// {
	// 	arr := make([]int, 5)
	// 	fmt.Println(len(arr))
	// 	fmt.Println(cap(arr))
	// 	fmt.Println(arr)
	// }

	// {
	// 	arr := make([]int, 5, 8)
	// 	fmt.Println(len(arr))
	// 	fmt.Println(cap(arr))
	// 	fmt.Println(arr)
	// }
}
func testfunc() {
	var xx uint
	xx = 16
	var yy int
	yy = -16
	fmt.Println(xx << 1)
	fmt.Println(yy << 1)
	fmt.Println(xx >> 1)
	fmt.Println(yy >> 1)
}

var (
	coin = []int{1, 2, 5, 10}
	// results = [...]int{1}

)

func getReward(sum int, result []int) {
	if sum == 0 {
		fmt.Println(result)
		return
	} else if sum < 0 {
		return
	} else {
		for i := 0; i < 4; i++ {
			newRet := append(result, coin[i])
			getReward(sum-coin[i], newRet)
		}
	}

}
func existone(num int, result []int) (numofone int) {
	numofone = 0
	for _, v := range result {
		if v == num {
			numofone++
		}
	}
	return
}
func dv(final int, result []int) {
	numofone := existone(1, result)
	if numofone > 1 {
		return
	}
	temp := 1
	for _, v := range result {
		temp *= v
	}
	if temp == final {
		numoffinal := existone(final, result)
		if numoffinal == 1 && numofone < 1 {
			result = append(result, 1)
		}
		fmt.Println(result)
		return
	} else if temp > final {
		return
	} else {
		for i := 1; i <= final; i++ {
			newRet := append(result, i)
			dv(final, newRet)
		}
	}
}
func merges(array []int) (ret []int) {
	fmt.Println("all:", array)
	length := len(array)
	if length < 2 {
		ret = array[:1]
		return
	} else if length == 2 {
		ret = mergeSort(array[:1], array[1:])
		return
	}
	leftone := array[:len(array)/2]
	rightone := array[len(array)/2:]
	left := merges(leftone)
	right := merges(rightone)
	ret = mergeSort(left, right)
	return
}
func mergeSort(left []int, right []int) (ret []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else if left[i] == right[j] {
			ret = append(ret, left[i])
			i++
			j++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}
	if i < len(left) {
		ret = append(ret, left[i:]...)
	}
	if j < len(right) {
		ret = append(ret, right[j:]...)
	}
	return
	// fmt.Println("left:", left)
	// fmt.Println("right:", right)
	// leftLength := len(left)
	// rightLength := len(right)
	// if leftLength < 1 || rightLength < 1 {
	// 	return
	// }
	// if leftLength == 1 && rightLength == 1 {
	// 	if left[0] > right[0] {
	// 		ret = append(ret, right[0])
	// 		ret = append(ret, left[0])
	// 	} else {
	// 		ret = append(ret, left[0])
	// 		ret = append(ret, right[0])
	// 	}
	// 	return
	// }
	// splitLeft := leftLength / 2
	// splitRight := rightLength / 2
	// leftMerge := left
	// rightMerge := right
	// if splitLeft > 0 {
	// 	leftMerge = mergeSort(left[:splitLeft], left[splitLeft:])
	// }
	// if splitRight > 0 {
	// 	rightMerge = mergeSort(right[:splitRight], right[splitRight:])
	// }
	// return mergeSort(leftMerge, rightMerge)
}

func testbtcsign() {
	hd := wallet.NewHdwallet()
	err := hd.ImportMnemonic(imtokenmnemonic)
	if err != nil {
		fmt.Println(err)
		return
	}
	{
		address, private, err := hd.GenerateAddressWithMnemonic(0, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
	{
		address, private, err := hd.GenerateAddressWithMnemonic(0, 0, 0, 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}

	// SendBTCRawTxByPrivateKey(privateKey string, toAddress string, amount float64, txFee float64) (signedParam string, err error)

	sign, err := wallet.SendBTCRawTxByPrivateKey(btc0, btc1address, 0.001, 0.00004, "测试汉字测试汉字测试汉字测试汉字测试试汉字测试汉字测子")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sign)
}
func testipfs() {
	fmt.Printf("\nipfs:\n")
	c := cfg.Config{}
	priv, pub, err := ci.GenerateKeyPairWithReader(ci.ECDSA, 1024, rand.Reader)
	if err != nil {
		return
	}

	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return
	}

	privkeyb, err := priv.Bytes()
	if err != nil {
		return
	}

	c.Bootstrap = cfg.DefaultBootstrapAddresses
	c.Addresses.Swarm = []string{"/ip4/0.0.0.0/tcp/4001"}
	c.Identity.PeerID = pid.Pretty()
	// fmt.Println(string(privkeyb))
	c.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)
	fmt.Println(c.Identity.PeerID)
	fmt.Println(c.Identity.PrivKey)
	private_str := hex.EncodeToString(privkeyb)
	fmt.Println(private_str)
}
func test2() {
	btswallets.GetBtsKey("xx", "tt")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	address, private, err := hd.GenerateAddress(60, 0, 0, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(hd.Mnemonic())
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// }
	fmt.Printf("\neth:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(60, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
		fmt.Println(hd.MasterKey())
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("\nbtc:\n")
			{
				hd := wallet.NewHdwallet()
				err := hd.ImportMnemonic(imtokenmnemonic)
				if err != nil {
					fmt.Println(err)
					return
				}
				address, private, err := hd.GenerateAddressWithMnemonic(0, 0, i, j)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(private)
				fmt.Println(address)
				// fmt.Println(hd.MasterKey())
			}
		}
	}

	// fmt.Printf("\nbtc:\n")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	err := hd.ImportMnemonic(imtokenmnemonic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	address, private, err := hd.GenerateAddressWithMnemonic(0, 0, 0, 1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// 	// fmt.Println(hd.MasterKey())
	// }
	// fmt.Printf("\nltc:")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	err := hd.ImportMnemonic(imtokenmnemonic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	address, private, err := hd.GenerateAddressWithMnemonic(2, 0, 0, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// }
	// fmt.Printf("\ndoge:")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	err := hd.ImportMnemonic(imtokenmnemonic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	address, private, err := hd.GenerateAddressWithMnemonic(3, 0, 0, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// }
	// fmt.Printf("\nqtum:")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	err := hd.ImportMnemonic(imtokenmnemonic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	address, private, err := hd.GenerateAddressWithMnemonic(4, 0, 0, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// }
	// fmt.Printf("\nnuls:")
	// {
	// 	hd := wallet.NewHdwallet()
	// 	err := hd.ImportMnemonic(imtokenmnemonic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	address, private, err := hd.GenerateAddressWithMnemonic(6, 0, 0, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(private)
	// 	fmt.Println(address)
	// }
}
func rmstring(origin []string, elem string) (ret []string) {
	// length := len(origin)
	// ret = make([]string, length)

	// for i := 0; i < length; i++ {
	// 	if origin[i] != elem {
	// 		ret = append(ret, origin[i])
	// 	}
	// }
	for _, v := range origin {
		if v != elem {
			ret = append(ret, v)
		}
	}
	return
}
func existstring(result []string, a string) bool {
	for _, v := range result {
		if v == a {
			return true
		}
	}
	return false
}
func permutate(choice []string, result []string) {
	// fmt.Println("choice:", choice)
	// fmt.Println("result:", result)
	if len(choice) == 0 {
		fmt.Println(result)
		return
	}
	for _, v := range choice {
		if !existstring(result, v) {
			newresult := append(result, v)
			newchoice := rmstring(choice, v)
			// fmt.Println("newchoice:", choice)
			permutate(newchoice, newresult)
			// return
		}
	}
}
func combine(choice []int, result []int, num int) {
	// fmt.Println("choice:", choice)
	// fmt.Println("result:", result)
	if num < 1 {
		fmt.Println(result)
		counter++
		return
	}
	// if len(result) == num {
	// 	fmt.Println(result)
	// 	return
	// }
	if len(choice) < num {
		return
	}
	index := len(result) - 1
	var start int
	if index < 0 {
		start = 0
	} else {
		start = result[index] + 1
	}
	for i := start; i < len(choice); i++ {
		newresult := append(result, choice[i])
		combine(choice, newresult, num-1)
	}
}
func dynamic() {
	coin := []int{1, 2, 5, 10}
	var c [101]int //和为i的最小组合的钱币个数
	c[0] = 0
	for i := 1; i < 101; i++ {
		minValue := 100
		for _, value := range coin {
			temp := i - value
			if temp >= 0 {
				if c[temp] < minValue {
					minValue = c[temp]
				}
			}
		}
		c[i] = minValue + 1
	}
	fmt.Println(c)
}

type TreeNode struct {
	sons    map[string]*TreeNode
	auxList *list.List
}

func NewTreeNode() (ret *TreeNode) {
	ret = &TreeNode{
		sons:    make(map[string]*TreeNode),
		auxList: list.New(),
	}
	return
}
func (t *TreeNode) BreadthFirstSearch() {

	for k, v := range t.sons {
		fmt.Println(k)
		t.auxList.PushBack(v)
	}
	for t.auxList.Len() > 0 {
		front := t.auxList.Front()
		if es, ok := front.Value.(*TreeNode); ok {
			es.BreadthFirstSearch()
		}
		t.auxList.Remove(front)
	}
}
func testbfs() {
	root := NewTreeNode()

	second1 := NewTreeNode()
	second2 := NewTreeNode()
	root.sons["1"] = second1
	root.sons["2"] = second2
	third1 := NewTreeNode()
	third2 := NewTreeNode()
	third3 := NewTreeNode()
	third4 := NewTreeNode()

	second1.sons["3"] = third1
	second1.sons["4"] = third2
	second2.sons["5"] = third3
	second2.sons["6"] = third4
	root.BreadthFirstSearch()
}

// func dijkstra() {
// 	weight := [][]int{
// 		{1000, 5, 8, 1000, 2},
// 		{5, 1000, 1, 3, 1000},
// 		{8, 1, 1000, 1000, 1000},
// 		{1000, 3, 1000, 1000, 7},
// 		{1000, 2, 1000, 7, 1000},
// 	}
// 	used := []bool{false, false, false, false, false}
// 	dis := []int{0, 1000, 1000, 1000, 1000}
// 	min := 1000
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if weight[i][j] < min {
// 				dis[j]
// 			}
// 		}
// 	}
// }
