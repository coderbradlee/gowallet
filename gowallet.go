package main

import (
	"fmt"
	wallet "hdwallet"
	// "math/big"
)

var (
	nonece uint64 = 0
)

//Just for test
func main() {
	number := 1
	err := GenerateWallets(uint32(number))
	if err != nil {
		println(err.Error())
		return
	}
}

func GenerateWallets(number uint32) (err error) {
	{
		// velvet bid mask thank joke educate edit business advance valley book surround
		//0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5
		//0x6356908ACe09268130DEE2b7de643314BBeb3683
		encryptedmk, err := wallet.CreateWalletByMnnicAndPwd("velvet bid mask thank joke educate edit business advance valley book surround", "123password")

		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println("encryptedmk", encryptedmk)
		addr, err := wallet.GenerateBIP44AccountWallet(encryptedmk, "ETH", 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("addr", addr)
	}
	{
		// encryptedmk, err := wallet.CreateWalletByMnnicAndPwd("much vacant moral dumb marble now require radio later there broccoli vapor", "123password")

		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("encryptedmk", encryptedmk)
		// addr, err := wallet.GenerateBIP44AccountWallet(encryptedmk, "ETH", 0, 0, 0)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("addr", addr)
	}

	{
		// masterKeyWithmnemonic, err := wallet.CreateWalletByteRandAndPwd("test", "test")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// fmt.Println("masterKeyWithmnemonic", masterKeyWithmnemonic)
	}
	// var secret, salt string
	// secret = "ShowSplashViewShowSplashViewShowSplashViewShowSplashView"
	// salt = "1234567890"
	// var byteSecret []byte = []byte(secret)
	// var byteSalt []byte = []byte(salt)
	// //wa, err := wallet.NewWalletAccount(wp.SecretBytes(), wp.SaltBytes())
	// addr, privateKey, err := wallet.NewWalletAccount(byteSecret, byteSalt)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("The Private key is  " + privateKey)
	// fmt.Println(" The Address is  " + addr)
	// testAddr := "0x4661dbc978fd123e2250a33c9eedcfeec3746ec5"
	// signedData, _ := wallet.SendETHRawTxByPrivateKey(privateKey, nonece+3, testAddr, big.NewInt(1000000000), big.NewInt(21000), big.NewInt(18000000000), nil)
	// fmt.Println("The real signed hex string is ", signedData)
	return err
}
