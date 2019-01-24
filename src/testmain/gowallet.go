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

	// {//imtoken btc
	// 	PrivateKey2=e2beee031b183c9f0d0efe946152497cd9dabc16e55ee8ba12e830bb1aaff7f7
	//     Mnemonic2=blouse hedgehog during sun vehicle calm panther scene wash grid mimic divorce
	//     Keystore2={"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"6c2460fbda35b4dfd61e3ff16a0c7fa6"},"ciphertext":"cdd0a2c939c60e7bc0dadc74a3b3ed2c1757e97baf9d0da329b63ce11c0e775f","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"d5bed3edfebaace94046ad26d708fd33fad0c1e43f105d1b2d62286b41002bd3"},"mac":"49ec75e49ef6059f4e1269abc3037b811c633f7ae01536c7151fa0afc53ae3a3"},"id":"932e090f-45f1-4cdf-8ab8-0daf1696d0fd","version":3,"address":"1QFVGastye8fonetC1VC6MQ16NAUmsJDsz"}
	// }
	//metamask
	// velvet bid mask thank joke educate edit business advance valley book surround
	//0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5
	//0x6356908ACe09268130DEE2b7de643314BBeb3683
	// mnemonic   = "velvet bid mask thank joke educate edit business advance valley book surround"
	imtokenmnemonic = "blouse hedgehog during sun vehicle calm panther scene wash grid mimic divorce"
	ethaddress      = "0xd73eab1b58a8f7936ce5a9eccdd9bad472ab6d28"
)

//Just for test
func main() {
	fmt.Println("imtoken address:", ethaddress)
	fmt.Printf("\n")
	test()
}

func test() {
	{
		hd := wallet.NewHdwallet()
		address, private, err := hd.GenerateAddress(60, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(hd.Mnemonic())
		fmt.Println(private)
		fmt.Println(address)
	}
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
	}
	fmt.Printf("\nbtc:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddress(0, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
	fmt.Printf("\nltc:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddress(2, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
}
