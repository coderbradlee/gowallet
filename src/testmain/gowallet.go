package main

import (
	"fmt"
	wallet "hdwallet"
	// "math/big"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
	"crypto/rand"
	cfg "github.com/ipfs/go-ipfs-config"
	ci "github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"
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
	//对应的eth
	//        PrivateKey=038d3e74691601d876e11badc1c9e7d00631f0ef708768baa53f367520788d5f
	//        Mnemonic=blouse hedgehog during sun vehicle calm panther scene wash grid mimic divorce
	//        Keystore={"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"ddeb8e7503b651c3d523c972f49417c3"},"ciphertext":"9847b471f201c025e817a4850146ef5a58b454e317f48fb53760dc1b483c7a05","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"f8b363a1a7573f40d84d007059efcbe38e3b2f35d1de1741fc6df697c7a6628d"},"mac":"4c7480ce04705d8bdd1ca49acc4ffc9aaa18d80479f570362db46c76a5ba90b9"},"id":"5d18bb08-2019-4292-95bb-722a8223f1e8","version":3,"address":"1eff9574401979e2d94c652d9069a6c250d65bb5"}
	// }
	//metamask
	// velvet bid mask thank joke educate edit business advance valley book surround
	//0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5
	//0x6356908ACe09268130DEE2b7de643314BBeb3683
	// mnemonic   = "velvet bid mask thank joke educate edit business advance valley book surround"
	imtokenmnemonic = "blouse hedgehog during sun vehicle calm panther scene wash grid mimic divorce"
	ethaddress      = "1eff9574401979e2d94c652d9069a6c250d65bb5"
	btcaddress      = "1QFVGastye8fonetC1VC6MQ16NAUmsJDsz"
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

	}

	fmt.Printf("\nltc:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(2, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
	fmt.Printf("\ndoge:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(3, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
	fmt.Printf("\nqtum:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(4, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
	fmt.Printf("\nnuls:")
	{
		hd := wallet.NewHdwallet()
		err := hd.ImportMnemonic(imtokenmnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, private, err := hd.GenerateAddressWithMnemonic(6, 0, 0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(private)
		fmt.Println(address)
	}
}
