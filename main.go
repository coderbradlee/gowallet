package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lzxm160/gowallet/src/hdwallet"
)

func main() {
	f, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		hd := hdwallet.Hdwallet{}
		addr, pri, err := hd.GenerateAddress(0, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("btc:", addr)
		balance, err := hdwallet.GetBTCBalanceByAddr(addr)
		if err != nil {
			fmt.Println("btc:", err)
			continue
		}
		if balance != "0" {
			f.WriteString(addr + ":" + pri + ":" + balance)
		}

		/////////
		addr, pri, err = hd.GenerateAddress(2, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("ltc:", addr)
		//balance, err = hdwallet.GetLTCBalanceByAddr(addr)
		balance, err = hdwallet.GetLTCBalanceByAddr("LZ43jcFdxNVpJWJ6o3neYEsnqEGxQTsP9M")
		if err != nil {
			fmt.Println("ltc:", err)
			continue
		}
		if balance != "0" {
			fmt.Println("ltc balance:", balance)
			f.WriteString(addr + ":" + pri + ":" + balance)
		}
		///////////////////
		addr, pri, err = hd.GenerateAddress(60, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("eth:", addr)
		balance, err = hdwallet.GetBalance(addr)
		if err != nil {
			fmt.Println("eth:", err)
			continue
		}
		if balance != "0" {
			f.WriteString(addr + ":" + pri + ":" + balance)
		}
		time.Sleep(time.Second)
	}

}
