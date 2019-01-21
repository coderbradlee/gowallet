package wallet

/* This pacage's function is as follows:
 *** generateSeed(secret, salt)
 *** GenerateMnemonic(seed)
 *** GenerateMasterkey
 *** GenerateAccount
 *** GenerateIndex
 ***  CreateWallet by  random and password
 */

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

const hardened = 0x80000000

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var AddressNetParams = chaincfg.MainNetParams

type WalletAccount struct {
	PrivateKey string
	PublicKey  string
	Mnemonic   string
}

// Mnemonic Import
func ImportMnemonic(mnemonic string) ([]byte, error) {
	return bip39.MnemonicToByteArray(mnemonic)

}

// Mnemonic Generation
func GenerateMnemonic(entropy []byte) (string, error) {

	if len(entropy) < 0 {
		entropy, _ = bip39.NewEntropy(256)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	return mnemonic, err
}

func GenerateMasterkey(masterSeed []byte) (string, error) {
	masterKey, err := hdkeychain.NewMaster(masterSeed, &AddressNetParams)
	return masterKey.String(), err
}

//func CreateWalletByteRandAndPwd(random, salt []byte) (masterKey, mnemonic string, err error)
// func CreateWalletByteRandAndPwd(random []byte, password string) (string,string,error) {
// 	//var seed []byte
// 	if len(random) < 0 {
// 		random, _ = bip39.NewEntropy(256)
// 	}

// 	seed, err := GenerateSeed(random, []byte(password))
// 	fmt.Println("The Real seed to byte is: #v%", seed)
// 	if err != nil {
// 		return "","",err
// 	}
// 	//Create Mnemonic
// 	mnemonic, err := GenerateMnemonic(seed)
// 	if err != nil {
// 		return "","",err
// 	}
// 	fmt.Println("The mnemonic word list is ", mnemonic)

// 	//Import Mnemonic
// 	mnemonicSeed, err := ImportMnemonic(mnemonic)
// 	if err != nil {
// 		return "","",err
// 	}
// 	fmt.Println("The mnemonic word list to byte is: #v%", mnemonicSeed)
// 	masterKeyStr, err := GenerateMasterkey(seed)
// 	if err != nil {
// 		return "","",err
// 	}
// 	//encrypt the masterKey With Password
// 	//masterKeyEncode = encodeMasterKeyWithPasword(masterKeyStr, password)
// 	return masterKeyStr, mnemonic, err
// }

//Check Password is right //TODO
func CheckPwdIsCorrect(masterKey, password string) bool {
	return true
}

// //TODO
// func encodeMasterKeyWithPasword(masterKeyStr, password string) (masterKey string) {
// 	//encode
// 	masterKey = masterKeyStr + password
// 	return
// 	//decode
// }

//CreateWallet or ImportWallet by  mnemonic
func CreateWalletByMnnicAndPwd(mnemonic string, salt []byte) (string,error) {

	//Import Mnemonic
	mnemonicSeed, err := ImportMnemonic(mnemonic)
	if err != nil {
		return "",nil
	}
	fmt.Println("The mnemonic word list to byte is: #v%", mnemonicSeed)
	masterKeyStr, err := GenerateMasterkey(mnemonicSeed)
	if err != nil {
		return "",nil
	}
	return masterKeyStr, err
}

func NewWalletAccount(secret, salt []byte) (string,error) {
	//wa = new(WalletAccount)
	var address, privateKey string
	var err error
	var seed []byte
	seed, err = GenerateSeed(secret, salt)
	fmt.Println("The Real seed to byte is: #v%", seed)
	if err != nil {
		return "",err
	}
	//Create Mnemonic
	mnemonic, err := GenerateMnemonic(seed)
	if err != nil {
		return "",err
	}
	fmt.Println("The mnemonic word list is ", mnemonic)

	//Import Mnemonic
	mnemonicSeed, err := ImportMnemonic(mnemonic)
	if err != nil {
		return "",err
	}
	fmt.Println("The mnemonic word list to byte is: #v%", mnemonicSeed)
	masterKeyStr, err := GenerateMasterkey(seed)
	//master_key, err := hdkeychain.NewMaster(seed, &AddressNetParams)
	if err != nil {
		return "",err
	}
	//err = wa.generateAccount(master_key.String(), 0)
	//err = wa.GenerateBIP44AccountPath(master_key.String(), "BTC", 0, 0, 0)
	address, privateKey,err=generateBIP44AccountWallet(masterKeyStr, "ETC", 0, 0, 0)
	return `{"addr":"`+address+`","priv":"`+privateKey+`"}`,err
}

// // Generate wallet seed from secret and salt
func GenerateSeed(secret, salt []byte) ([]byte,error) {
	// WarpWallet encryption:
	// 1. s1 ← scrypt(key=passphrase||0x1, salt=salt||0x1, N=218, r=8, p=1, dkLen=32)
	// 2. s2 ← PBKDF2(key=passphrase||0x2, salt=salt||0x2, c=216, dkLen=32)
	// 3. private_key ← s1 ⊕ s2
	// 4. Generate public_key from private_key using standard Bitcoin EC crypto
	// 5. Output (private_key, public_key)
	var seed []byte
	var err error
	if len(secret) == 0 {
		err = errors.New("empty secret")
		return seed,err
	}
	if len(salt) == 0 {
		err = errors.New("empty salt")
		return seed,err
	}

	secret1 := make([]byte, len(secret))
	secret2 := make([]byte, len(secret))
	for i, v := range secret {
		secret1[i] = byte(v | 0x01)
		secret2[i] = byte(v | 0x02)
	}

	salt1 := make([]byte, len(salt))
	salt2 := make([]byte, len(salt))
	for i, v := range salt {
		salt1[i] = byte(v | 0x01)
		salt2[i] = byte(v | 0x02)
	}

	s1, err := scrypt.Key(secret1, salt1, 16384, 8, 1, 32)
	if err != nil {
		return seed,err
	}
	s2 := pbkdf2.Key(secret2, salt2, 4096, 32, sha1.New)

	pk1, _ := btcec.PrivKeyFromBytes(btcec.S256(), s1)
	pk2, _ := btcec.PrivKeyFromBytes(btcec.S256(), s2)
	x, y := btcec.S256().Add(pk1.X, pk1.Y, pk2.X, pk2.Y)

	seed = []byte{0x04}
	seed = append(seed, x.Bytes()...)
	seed = append(seed, y.Bytes()...)

	seed_hash := sha256.Sum256(seed)
	seed = seed_hash[:]
	return seed,err
}

// // Generate BIP44 account extended private key and extended public key.
func (wa *WalletAccount) generateAccount(masterKey string, k uint32) (error) {
	var err error
	master_key, err := hdkeychain.NewKeyFromString(masterKey)
	if err != nil {
		return err
	}
	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return err
	}
	//Coin type: maybe changed by different coin type
	coin_type, err := purpose.Child(hardened + 60)
	//coin_type, err := purpose.Child(hardened + 0)
	if err != nil {
		return err
	}
	account, err := coin_type.Child(hardened + k)

	if err != nil {
		return err
	}
	account_pub, err := account.Neuter()
	if err != nil {
		return err
	}

	wa.PrivateKey = account.String()
	wa.PublicKey = account_pub.String()
	return nil
}

// // Generate BIP44 account extended private key and extended public key by the path.
// //(wa *WalletAccount) generateAccount(masterKey string, k uint32)
func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int64) (string,error) {
	var address, privateKey string
	var err error
	master_key, err := hdkeychain.NewKeyFromString(masterKey)
	var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return "",err
	}

	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return "",err
	}

	var flag uint32
	//Coin type: maybe changed by different coin type
	if coinType == "BTC" {
		drivedCoinType, err = purpose.Child(hardened + 0)
		flag = 0
	} else if coinType == "ETH" {
		drivedCoinType, err = purpose.Child(hardened + 60)
		flag = 1
	} else if coinType == "ETC" {
		drivedCoinType, err = purpose.Child(hardened + 60)
		flag = 1
	} else {
		return "", errors.New("The Coin Type is not support!!!")
	}
	if err != nil {
		return "",err
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + uint32(account))
	if err != nil {
		return "",err
	}

	//Change(T/F:1,0)
	change = 0
	drivedChange, err := drivedAccount.Child(uint32(change))
	if err != nil {
		return "",err
	}
	//create change Index
	index = 0

	address, privateKey,err=createChangeIndex(drivedChange, int32(index), int32(flag))
	// return address, privateKey,err
	ret:=`{"addr":"`+address+`","priv":"`+privateKey+`"}`
	return ret,err
}
func generateBIP44AccountWallet(masterKey string, coinType string, account, change, index int64) (string,string,error) {
	var address, privateKey string
	var err error
	master_key, err := hdkeychain.NewKeyFromString(masterKey)
	var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return address, privateKey,err
	}

	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return address, privateKey,err
	}

	var flag uint32
	//Coin type: maybe changed by different coin type
	if coinType == "BTC" {
		drivedCoinType, err = purpose.Child(hardened + 0)
		flag = 0
	} else if coinType == "ETH" {
		drivedCoinType, err = purpose.Child(hardened + 60)
		flag = 1
	} else if coinType == "ETC" {
		drivedCoinType, err = purpose.Child(hardened + 60)
		flag = 1
	} else {
		return address, privateKey, errors.New("The Coin Type is not support!!!")
	}
	if err != nil {
		return address, privateKey,err
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + uint32(account))
	if err != nil {
		return address, privateKey,err
	}

	//Change(T/F:1,0)
	change = 0
	drivedChange, err := drivedAccount.Child(uint32(change))
	if err != nil {
		return address, privateKey,err
	}
	//create change Index
	index = 0

	address, privateKey,err=createChangeIndex(drivedChange, int32(index), int32(flag))
	return address, privateKey,err
	// ret:=`{"addr":"`+address+`","priv":"`+privateKey+`"}`
	// return ret,err
}

func createChangeIndex(change *hdkeychain.ExtendedKey, index int32, flag int32) (string,string,error) {
	var address, privateKey string
	var err error
	// var addressStr string
	child, err := change.Child(uint32(index))
	if err != nil {
		return address, privateKey,err
	}
	private_key, err := child.ECPrivKey()
	if err != nil {
		return address, privateKey,err
	}

	//ETC ETH
	if flag == 1 {
		//Private key
		privateKeyBytes := private_key.Serialize()
		private_str := hex.EncodeToString(privateKeyBytes)
		fmt.Println("The ETH privateKeyBytes is ", private_str)
		ethaddress_key, err1 := AddressforEth(child) //child.AddressforEth()
		if err1 != nil {
			err=err1
			return address, privateKey,err
		}
		address = hex.EncodeToString(ethaddress_key)
		// fmt.Println("The ETH address is ", addressStr)
	} else {
		//BTC
		//private_wif, err := btcutil.NewWIF(private_key, &AddressNetParams, true)
		if err != nil {
			return address, privateKey,err
		}
		//private_str := private_wif.String()
		address_str, err1 := child.Address(&AddressNetParams)
		if err1 != nil {
			err=err1
			return address, privateKey,err
		}
		address = address_str.String()
		// fmt.Println("The BTC address is ", addressStr)
	}
	privateKey = child.String()
	return address, privateKey,err
	// return `{"addr":"`+address+`","priv":"`+privateKey+`"}`,err
}

func getChildByPath(masterKey string, coinType string, account, change, index int32) (child *hdkeychain.ExtendedKey, err error) {

	master_key, err := hdkeychain.NewKeyFromString(masterKey)
	var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return
	}
	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return
	}

	//Coin type: maybe changed by different coin type
	if coinType == "BTC" {
		drivedCoinType, err = purpose.Child(hardened + 0)
	} else if coinType == "ETH" {
		drivedCoinType, err = purpose.Child(hardened + 60)
	} else if coinType == "ETC" {
		drivedCoinType, err = purpose.Child(hardened + 60)
	} else {
		return nil, errors.New("The Coin Type is not support!!!")
	}
	if err != nil {
		return
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + uint32(account))
	if err != nil {
		return
	}

	//Change(T/F:1,0)
	change = 0
	drivedChange, err := drivedAccount.Child(uint32(change))
	if err != nil {
		return nil, err
	}
	//create change Index
	index = 0
	child,err=drivedChange.Child(uint32(index))
	return
}

func getChildByPrivatekeyStr(privateKey string) (child *hdkeychain.ExtendedKey, err error) {
	child, err = hdkeychain.NewKeyFromString(privateKey)
	return
}

// //Send ETH RawTransaction
func SendETHRawTxByPath(masterKey string, coinType string, account, change, index int32, nonce uint64, toAddr string, amount, gasLimit, gasPrice *big.Int, data []byte) (string,error) {
	var signedParam string
	var err error
	child, err := getChildByPath(masterKey, coinType, account, change, index)
	address := make([]byte, 20)
	//toAddr need to check the prefix 0x
	find := (strings.HasPrefix(toAddr, "0x")) || (strings.HasPrefix(toAddr, "0X"))
	var tempAddr string
	if find {
		tempAddr = toAddr[2:len(toAddr)]
	}
	address, err = hex.DecodeString(tempAddr)
	if err != nil {
		return signedParam,err
	}
	var toAddress common.Address
	for i := 0; i < len(address); i++ {
		toAddress[i] = address[i]
	}
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	private_key, _ := child.ECPrivKey()
	signed, err := types.SignTx(tx, types.HomesteadSigner{}, (*ecdsa.PrivateKey)(private_key))
	if err != nil {
		return signedParam,err
	}
	if err != nil {
		return signedParam,err
	}

	encodeData, err := rlp.EncodeToBytes(signed)
	if err != nil {
		return signedParam,err
	}
	signedParam= common.ToHex(encodeData)

	// fmt.Println("The real ETH sig is ", signed.String())
	// fmt.Println("The real signed hex string is ", hexString)
	return signedParam,err

}

// //Send ETH RawTransaction
func SendETHRawTxByPrivateKey(privateKey string, nonce uint64, toAddr string, amount, gasLimit, gasPrice *big.Int, data []byte) (string,error) {
	var signedParam string
	var err error
	
	child, err := hdkeychain.NewKeyFromString(privateKey)
	address := make([]byte, 20)
	//toAddr need to check the prefix 0x
	find := (strings.HasPrefix(toAddr, "0x")) || (strings.HasPrefix(toAddr, "0X"))
	var tempAddr string
	if find {
		tempAddr = toAddr[2:len(toAddr)]
	}
	address, err = hex.DecodeString(tempAddr)
	if err != nil {
		return signedParam,err
	}
	var toAddress common.Address
	for i := 0; i < len(address); i++ {
		toAddress[i] = address[i]
	}
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	private_key, _ := child.ECPrivKey()
	signed, err := types.SignTx(tx, types.HomesteadSigner{}, (*ecdsa.PrivateKey)(private_key))
	if err != nil {
		return signedParam,err
	}

	encodeData, err := rlp.EncodeToBytes(signed)
	if err != nil {
		return signedParam,err
	}
	signedParam= common.ToHex(encodeData)

	// fmt.Println("The real ETH sig is ", signed.String())
	// fmt.Println("The real signed hex string is ", hexString)
	return signedParam,err

	//Test with ETH endpoint
	//RUL := "http" + "://" + "172.16.2.130" + ":" + "8080"
	//	client, err := ethclient.Dial(RUL)
	//	if err != nil {
	//		fmt.Println("The ethclient Dial is wrong!!!")
	//	})
	//	client.SendTransactionTest(context.Background(), hexString)

}

func AddressforEth(k *hdkeychain.ExtendedKey) ([]byte, error) {
	publickey, _ := k.ECPubKey()
	var p *ecdsa.PublicKey
	p = (*ecdsa.PublicKey)(publickey)
	pubBytes := crypto.FromECDSAPub(p)
	pkPrv := common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
	pkHash := pkPrv[:]
	return pkHash, nil
}
