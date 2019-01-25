package main

import (
	"fmt"
	wallet "hdwallet"
	// "math/big"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
	"crypto/rand"
	"encoding/base64"
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
	// 私钥xprv9yrdwPSRnvomqFK4u1y5uW2SaXS2Vnr3pAYTjJjbyRZR8p9BwoadRsCxtgUFdAKeRPbwvGRcCSYMV69nNK4N2kadevJ6L5iQVy1SwGKDTHQ
	// 公钥BdgvWHN/Uh/K526q/+CdpGwEPZ41SvZHHGSgiSqhFesjErdbo6UnJMIoDOHV94qW8fd2KBW18UG3nTzDwS7a5oArqPtv+2aE9+1bNvCdtYoAx3979N3vbX4Xxn/najTABykXrJDjgpoaXxSo/xTktQ==

	imtokenmnemonic = "inject kidney empty canal shadow pact comfort wife crush horse wife sketch"
	ethaddress      = "6031564e7b2f5cc33737807b2e58daff870b590b"
	btcaddress      = "12z6UzsA3tjpaeuvA2Zr9jwx19Azz74D6g"
)

//Just for test
func main() {
	fmt.Println("imtoken address:", ethaddress)
	fmt.Printf("\n")
	test()
	testipfs()
}
func testipfs() {
	c := cfg.Config{}
	priv, pub, err := ci.GenerateKeyPairWithReader(ci.RSA, 1024, rand.Reader)
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

	c.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)
	fmt.Println(c.Identity.PeerID)
	fmt.Println(c.Identity.PrivKey)
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
	fmt.Printf("\nbtc:\n")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(0, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
		fmt.Println(hd.MasterKey())
	}
	fmt.Printf("\nbtc:\n")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(0, 0, 0, 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
		fmt.Println(hd.MasterKey())
	}
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
