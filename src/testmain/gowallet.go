package main

import (
	"fmt"
	wallet "hdwallet"
	// "math/big"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	cfg "github.com/ipfs/go-ipfs-config"
	ci "github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"
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
	if len(array) < 2 {
		return
	}
	left := array[:len(array)/2]
	right := array[len(array)/2:]
	fmt.Println(left)
	fmt.Println(right)
	ret = mergeSort(left, right)
	return
}
func mergeSort(left []int, right []int) (ret []int) {
	for i, j := 0, 0; i < len(left) && j < len(right); {
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
func main() {
	// fmt.Println("imtoken address:", ethaddress)
	// fmt.Printf("\n")
	// testbtcsign()
	// testfunc()
	// var results []int
	// getReward(10, results)
	left := []int{1, 8, 4}
	right := []int{2, 20, 33, 10}
	ret := mergeSort(left, right)
	fmt.Println(ret)
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
func test() {
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
