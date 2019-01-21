package hdwallet

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/eoscanada/eos-go/ecc"
	"hdwallet/nuls"
	"testing"
)

func TestNewMasterKey(t *testing.T) {

	//1.创建masterkey
	masterKey, err := CreateWalletByteRandAndPwd("", "123")
	fmt.Printf("masterkey and 助记词:%s,err=%s\n", masterKey, err)
}

//导入助记词钱包"some params is empty
func TestCreateWalletByMnnicAndPwd(t *testing.T) {
	mnemonic := "enlist natural gulp launching surfer orchid exit nylon sawmill coils wagtail chrome kept dagger dyslexic scoop roster rowboat"
	masterKey, _ := CreateWalletByMnnicAndPwd(mnemonic, "123456")
	fmt.Println(masterKey)
	coinType := []string{"BTC", "LTC", "DOGE", "QTUM", "ETH", "ETF", "ETC", "NULSM"}
	fmt.Println("钱包地址为：")
	for _, v := range coinType {
		address, p := GenerateBIP44AccountWallet(masterKey, v, 0, 0, 0)
		fmt.Printf("---------------%s=%s  prik:%s\n", v, address, p)
	}

}

func TestNulsBalance(t *testing.T) {
	amount, err := NulsBalance("6HgUzobcB8pa5coQjYhWMexRjDJdtM5e")
	if err != nil {
		fmt.Println("查询余额失败：", err.Error())
	} else {
		fmt.Printf("余额:%s \n", amount)
	}
}

func TestAddress(t *testing.T) {
	master_key, _ := hdkeychain.NewMaster([]byte("aaaaaaaaaaaaaaaa"), &BtcNetParams)

	drivedCoinType, _ := master_key.Child(1)
	nulsAddress := nuls.Address(drivedCoinType)
	fmt.Printf("nuls Address:%s \n", nulsAddress)
	//前面两个00 标示转为有符号的数字
	ecPrivKey, _ := drivedCoinType.ECPrivKey()
	fmt.Printf("nuls privateKey:00%x \n", ecPrivKey.Serialize())
	fmt.Printf("nuls publicKey:00%x \n", ecPrivKey.PubKey().SerializeCompressed())

	fmt.Printf("%x", base58.Decode(nulsAddress))

}

func TestTransfer(t *testing.T) {
	mnemonic := "enlist natural gulp launching surfer orchid exit nylon sawmill coils wagtail chrome kept dagger dyslexic scoop roster rowboat"
	masterKey, _ := CreateWalletByMnnicAndPwd(mnemonic, "123456")
	fmt.Println(masterKey)

	private_key, _ := CheckAuthAndGetPrivateKey(masterKey, "123456", "NULS", 0, 0, 0)

	fee, err := NulsTransferFee("6HgjScviL1LWubo7U79v2bb4sE3mQ9UF", "6HgeU17cGgKqGBqLdjJTowqDifEgC4Tm", "0.9", "", "")
	if err != nil {
		fmt.Println("失败：", err.Error())
	} else {
		fmt.Printf("手续费:%s \n", fee)

		sign, err := NulsTransfer(private_key, "6HgeU17cGgKqGBqLdjJTowqDifEgC4Tm", "0.9", "", "")
		if err != nil {
			fmt.Println("签名失败：", err.Error())
		} else {
			fmt.Printf("签名:%s \n", sign)
			//nulsBroadcast(sign)
		}
	}
}

func nulsBroadcast(txHex string) {
	v, err := NulsBroadcast(txHex)
	if err != nil {
		fmt.Println("广播失败：", err.Error())
	} else {
		fmt.Printf("广播成功返回hash:%s \n", v)
	}
}

func TestCreateRawTransactionNew(t *testing.T) {
	tx, err := createRawTransactionNew("17sgoUddjm3oHoYeyxjHr2W8RQzAJ5cuBG", "1LRPG7nyW8od1omTjFZkz6MxwWr1B6PxUR", 0.06820755, 0.0001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestEos(t *testing.T) {
	pri, _ := ecc.NewRandomPrivateKey()

	priKeyE := pri.String()
	fmt.Println(priKeyE)
	fmt.Println(pri.PublicKey().String())

}
func TestImport(t *testing.T) {
	//imtoken 助记词 测试例子
	//crisp bus ordinary fossil cliff inmate night program song patient elevator shallow
	//eth 地址0xd73eab1b58a8f7936ce5a9eccdd9bad472ab6d28
	encryptedmk, err := CreateWalletByMnnicAndPwd("crisp bus ordinary fossil cliff inmate night program song patient elevator shallow", "123password")
	if err != nil {
		fmt.Println("encryptedmk:", err)
		return
	}
	fmt.Println("encryptedmk:", encryptedmk)
	mk, err := decryptMasterkey(encryptedmk)
	if err != nil {
		fmt.Println("decryptMasterkey:", encryptedmk)
		return
	}
	fmt.Println("real master key:", mk)
	// GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int) (address string, err error)
	addr, err := GenerateBIP44AccountWallet(mk, "ETH", 0, 0, 0)
	if err != nil {
		fmt.Println("generate addr:", err)
		return
	}
	fmt.Println(addr)
}
