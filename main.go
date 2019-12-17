package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lzxm160/gowallet/src/hdwallet"
)

func main() {
	f, err := os.Open("output.txt")
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
		balance, err = hdwallet.GetLTCBalanceByAddr(addr)
		if err != nil {
			continue
		}
		if balance != "0" {
			f.WriteString(addr + ":" + pri + ":" + balance)
		}
		///////////////////
		//addr, pri, err = hd.GenerateAddress(60, 0, 0, 0)
		//if err == nil {
		//	continue
		//}
		//fmt.Println("eth:", addr)
		//
		//balance, err = hdwallet.(addr)
		//if err == nil {
		//	continue
		//}
		//if balance != "0" {
		//	f.WriteString(addr + ":" + pri + ":" + balance)
		//}
		time.Sleep(time.Millisecond * 100)
	}

}
