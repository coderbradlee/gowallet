package main

import (
	"fmt"
	wallet "hdwallet"
	// "math/big"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
)

var (
	// imtoken
	// crisp bus ordinary fossil cliff inmate night program song patient elevator shallow
	//eth 地址0xd73eab1b58a8f7936ce5a9eccdd9bad472ab6d28
	// mnemonic   = "crisp bus ordinary fossil cliff inmate night program song patient elevator shallow"
	// ethaddress = "0xd73eab1b58a8f7936ce5a9eccdd9bad472ab6d28"

	//metamask
	// velvet bid mask thank joke educate edit business advance valley book surround
	//0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5
	//0x6356908ACe09268130DEE2b7de643314BBeb3683
	// mnemonic   = "velvet bid mask thank joke educate edit business advance valley book surround"
	imtokenmnemonic = "crisp bus ordinary fossil cliff inmate night program song patient elevator shallow"
	ethaddress      = "0xd73eab1b58a8f7936ce5a9eccdd9bad472ab6d28"
)

//Just for test
func main() {
	fmt.Println("imtoken address:", ethaddress)
	fmt.Println("\n")
	test()
}

func test() {
	{
		hd := wallet.NewHdwallet()
		address, err := hd.GenerateAddress(60, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(address)
	}
	fmt.Println("\n")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, err := hd.GenerateAddressWithMnemonic(60, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(address)
	}
	fmt.Println("\n")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, err := hd.GenerateAddress(0, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(address)
	}
}
