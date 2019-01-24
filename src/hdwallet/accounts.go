package hdwallet

import (
	// "crypto/sha1"
	// "crypto/sha256"
	// "crypto/sha512"

	// "encoding/json"
	"errors"
	"fmt"
	// "github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	// "golang.org/x/crypto/pbkdf2"
	// "golang.org/x/crypto/scrypt"
	"hdwallet/nuls"
)

type Hdwallet struct {
	mnemonic  string
	masterKey string
	seed      []byte
}

func NewHdwallet() *Hdwallet {
	return &Hdwallet{}
}

func (hd *Hdwallet) generateMnemonic(entropy []byte) (err error) {
	if len(entropy) < 0 {
		entropy, err = bip39.NewEntropy(128)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	hd.mnemonic, err = bip39.NewMnemonic(entropy)
	return
}

func (hd *Hdwallet) mnemonicToSeed() (err error) {
	hd.seed, err = bip39.NewSeedWithErrorChecking(hd.mnemonic, "")
	return
}

func (hd *Hdwallet) generateMasterkey() (err error) {
	mk, err := bip32.NewMasterKey(hd.seed)
	hd.masterKey = mk.String()
	return
}

func (hd *Hdwallet) GenerateMnemonicAndMasterKey() (err error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}
	hd.mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return
	}
	err = hd.mnemonicToSeed()
	if err != nil {
		return
	}
	if len(hd.seed) < 2 {
		err = errors.New("The mnemonicSeed byte len is two low!!")
		return
	}
	err = hd.generateMasterkey()
	if err != nil {
		return
	}
	return
}

//CreateWallet or ImportWallet by  mnemonic
func (hd *Hdwallet) ImportMnemonic(mnemonic string) (err error) {
	if len(mnemonic) == 0 {
		err = errors.New("some params is empty!!!")
		return
	}
	hd.mnemonic = mnemonic
	err = hd.mnemonicToSeed()

	if len(hd.seed) < 2 {
		err = errors.New("The mnemonicSeed byte len is two low!!")
		return
	}
	err = hd.generateMasterkey()
	if err != nil {
		return
	}
	return
}

func (hd *Hdwallet) GenerateAddress(coinType, account, change, index int) (address, private string, err error) {
	err = hd.GenerateMnemonicAndMasterKey()
	if err != nil {
		return
	}
	if len(hd.masterKey) == 0 {
		err = errors.New("some params is empty!!!")
		return
	}
	return hd.GenerateAddressWithMnemonic(coinType, account, change, index)
}
func (hd *Hdwallet) GenerateAddressWithMnemonic(coinType, account, change, index int) (address, private string, err error) {

	master_key, err := hdkeychain.NewKeyFromString(hd.masterKey)
	// var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return
	}
	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return
	}
	drivedCoinType, err := purpose.Child(uint32(hardened + coinType))
	if err != nil {
		return
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + (uint32)(account))
	if err != nil {
		return
	}
	//Change(T/F:1,0)
	//change = 0
	drivedChange, err := drivedAccount.Child((uint32)(change))
	if err != nil {
		return
	}
	//create change Index
	//index = 0
	address, private, err = hd.createChangeIndex(drivedChange, index, coinType)
	return
}
func (hd *Hdwallet) Mnemonic() (mnemonic string) {
	return hd.mnemonic
}
func (hd *Hdwallet) createChangeIndex(change *hdkeychain.ExtendedKey, index int, coinType int) (address, privateKey string, err error) {
	child, err := change.Child((uint32)(index))
	if err != nil {
		return
	}
	private_key, err := child.ECPrivKey()

	if err != nil {
		return
	}
	switch coinType {
	case 0:
		address, privateKey, err = hd.btcAddress(private_key, child)
	case 60, 61, 63:
		address, privateKey, err = hd.ethAddress(private_key, child)
	case 2:
		//LTC
		address, privateKey, err = hd.ltcAddress(private_key, child)
	case 3:
		private_wif, err := btcutil.NewWIF(private_key, &dogeAddressNetParams, true)
		private_str := private_wif.String()
		address_str, err := child.Address(&dogeAddressNetParams)
		if err != nil {
			return "", "", err
		}
		address = address_str.String()
		fmt.Println("The DOGE private wif key is ", private_str)
		fmt.Println("The DOGE address is ", address)
	case 4:
		//private_wif, err := btcutil.NewWIF(private_key, &qtumAddressNetParams, true)
		//private_str := private_wif.String()
		address_str, err := child.Address(&qtumAddressNetParams)
		if err != nil {
			return "", "", err
		}
		address = address_str.String()
		//fmt.Println("The QTUM private wif key is ", private_str)
		fmt.Println("The QTUM address is ", address)
	case 6:
		address = nuls.Address(child)
		if len(address) <= 0 {
			return "", "", err
		}
		fmt.Println("The Nuls address is ", address)
	default:
		err = errors.New("not support")
	}
	if err != nil {
		return
	}
	privateKeyStr := child.String()
	return address, privateKeyStr, nil
}
