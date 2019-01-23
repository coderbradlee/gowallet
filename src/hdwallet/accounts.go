package hdwallet

import (
	// "crypto/sha1"
	// "crypto/sha256"
	// "crypto/sha512"
	"encoding/hex"
	"encoding/json"
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

// Mnemonic Import
func importMnemonic(mnemonic string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, "")
}

// Mnemonic Generation
func generateMnemonic(entropy []byte) (ret string, err error) {
	if len(entropy) < 0 {
		entropy, err = bip39.NewEntropy(128)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// fmt.Println("len(entropy):", len(entropy))
	return bip39.NewMnemonic(entropy)
}
func generateMasterkey(masterSeed []byte) (string, error) {
	// masterKey, err := hdkeychain.NewMaster(masterSeed, &btcAddressNetParams)

	masterKey, err := bip32.NewMasterKey(masterSeed)
	return masterKey.String(), err
}

//func CreateWalletByteRandAndPwd(random []byte, password string) (masterKey, mnemonic string, err error))
func CreateWalletByteRandAndPwd(rand string, password string) (masterKeyWithmnemonic string, err error) {
	//var seed []byte
	// random := []byte(rand)
	// if len(random) <= 0 {
	// 	random, _ = bip39.NewEntropy(128)
	// }

	// seed, err := generateSeed(random, []byte(password))
	// seedLen = len(seed)
	//fmt.Println("The Real seed len is :", seedLen)
	//fmt.Println("The Real seed to byte is: #v%", seed)
	// if err != nil {
	// 	return "", err
	// }
	//Create Mnemonic
	// seed := random
	seed, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}
	mnemonic, err := generateMnemonic(seed)
	if err != nil {
		return "", err
	}
	fmt.Println("The mnemonic word list is:", mnemonic)

	masterKeyStr, err := generateMasterkey(seed)
	if err != nil {
		return "", err
	}
	//fmt.Println("The origianl masterky is---->", masterKeyStr)
	//Add the MasterKeyWith the seed
	masterKeyStr = masterKeyStr + string(seed)

	//Encrpt the masterKey with password
	encryptMasterkey, err := encryptMastkeyWithPwd(masterKeyStr, password)
	if err != nil {
		return "", err
	}
	//fmt.Println("According by the seed ,The encrypt masterkey is", encryptMasterkey)
	//Convert to JSON
	var waAccount WalletAccount
	waAccount.MasterKey = encryptMasterkey
	waAccount.Mnemonic = mnemonic
	temp, err := json.Marshal(waAccount)
	return string(temp), err
}
func CreateNewMnemonicAndMasterKey(rand string, password string) (mnemonic, mk string, err error) {
	//var seed []byte
	// random := []byte(rand)
	// if len(random) <= 0 {
	// 	random, _ = bip39.NewEntropy(128)
	// }

	// seed, err := generateSeed(random, []byte(password))
	// seedLen = len(seed)
	//fmt.Println("The Real seed len is :", seedLen)
	//fmt.Println("The Real seed to byte is: #v%", seed)

	// seed := random
	seed, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}
	//Create Mnemonic
	mnemonic, err = generateMnemonic(seed)
	if err != nil {
		return
	}
	fmt.Println("The mnemonic word list is:", mnemonic)

	masterKeyStr, err := generateMasterkey(seed)
	if err != nil {
		return
	}
	//fmt.Println("The origianl masterky is---->", masterKeyStr)
	//Add the MasterKeyWith the seed
	masterKeyStr = masterKeyStr + string(seed)

	//Encrpt the masterKey with password
	mk, err = encryptMastkeyWithPwd(masterKeyStr, password)
	if err != nil {
		return
	}
	return
}

//CreateWallet or ImportWallet by  mnemonic
func CreateWalletByMnnicAndPwd(mnemonic string, password string) (masterKey string, err error) {
	if (len(mnemonic) == 0) || (len(password) == 0) {
		return "", errors.New("some params is empty!!!")
	}
	//Import Mnemonic
	mnemonicSeed, err := importMnemonic(mnemonic)
	if err != nil {
		return "", err
	}
	if len(mnemonicSeed) < 2 {
		return "", errors.New("The mnemonicSeed byte len is two low!!")
	}
	seedLen = len(mnemonicSeed)
	masterKeyStr, err := generateMasterkey(mnemonicSeed)
	if err != nil {
		return "", err
	}
	// fmt.Println("The origianl masterky is---->", masterKeyStr)
	//Add the MasterKeyWith the seed
	masterKeyStr = masterKeyStr + string(mnemonicSeed)

	//Encrpt the masterKey with password
	encryptMasterkey, err := encryptMastkeyWithPwd(masterKeyStr, password)
	if err != nil {
		return "", err
	}
	//fmt.Println("According by the mnemonic ,The encrypt masterkey is", encryptMasterkey)
	return encryptMasterkey, err
}

func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int) (address string, err error) {
	if (len(masterKey) == 0) || (len(coinType) == 0) {
		return "", errors.New("some params is empty!!!")
	}
	//Decrypt the masterkey
	decMasterkey, err := decryptMasterkey(masterKey)
	if err != nil {
		return "", err
	}
	fmt.Println("The decrypt masterky is---->", decMasterkey)
	master_key, err := hdkeychain.NewKeyFromString(decMasterkey)
	var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return "", err
	}

	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return "", err
	}

	var flag int
	//Coin type: maybe changed by different coin type
	if coinType == "BTC" {
		drivedCoinType, err = purpose.Child(hardened + 0)
		flag = 0
	} else if coinType == "ETH" {
		drivedCoinType, err = purpose.Child(hardened + 60)
		flag = 1
	} else if coinType == "ETC" {
		drivedCoinType, err = purpose.Child(hardened + 61)
		flag = 1
	} else if coinType == "ETF" {
		drivedCoinType, err = purpose.Child(hardened + 62)
		flag = 1
	} else if coinType == "LTC" {
		drivedCoinType, err = purpose.Child(hardened + 2)
		flag = 2
	} else if coinType == "DOGE" {
		drivedCoinType, err = purpose.Child(hardened + 3)
		flag = 3
	} else if coinType == "QTUM" {
		drivedCoinType, err = purpose.Child(hardened + 4)
		flag = 4
	} else if coinType == "NULSM" {
		drivedCoinType, err = purpose.Child(hardened + 6)
		flag = 5
	} else {
		return "", errors.New("The Coin Type is not support!!!")
	}
	if err != nil {
		return "", err
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + (uint32)(account))
	if err != nil {
		return "", err
	}

	//Change(T/F:1,0)
	//change = 0
	drivedChange, err := drivedAccount.Child((uint32)(change))
	if err != nil {
		return "", err
	}
	//create change Index
	//index = 0
	address, _, err = createChangeIndex(drivedChange, index, flag)

	return address, err

}
func createChangeIndex(change *hdkeychain.ExtendedKey, index int, flag int) (address, privateKey string, err error) {
	var addressStr string
	child, err := change.Child((uint32)(index))
	if err != nil {
		return "", "", err
	}
	private_key, err := child.ECPrivKey()

	if err != nil {
		return "", "", err
	}

	if flag == 0 { //BTC
		//private_wif, err := btcutil.NewWIF(private_key, &AddressNetParams, true)
		//private_str := private_wif.String()
		address_str, err := child.Address(&btcAddressNetParams)
		if err != nil {
			return "", "", err
		}
		addressStr = address_str.String()
		fmt.Println("The BTC address is ", addressStr)
	} else if flag == 1 { //ETC ETH ETF
		//Private key
		privateKeyBytes := private_key.Serialize()
		private_str := hex.EncodeToString(privateKeyBytes)
		fmt.Println("The ETH/ETC privateKeyBytes is ", private_str)
		ethaddress_key, err := addressforEth(child) //child.AddressforEth()
		if err != nil {
			return "", "", err
		}
		addressStr = hex.EncodeToString(ethaddress_key)
		// fmt.Println("The ETH/ETC/ETF address is ", addressStr)
	} else if flag == 2 {
		//LTC
		private_wif, err := btcutil.NewWIF(private_key, &ltcAddressNetParams, true)
		if err != nil {
			return "", "", err
		}
		private_str := private_wif.String()
		address_str, err := child.Address(&ltcAddressNetParams)
		if err != nil {
			return "", "", err
		}
		addressStr = address_str.String()
		fmt.Println("The LTC private wif key is ", private_str)
		fmt.Println("The LTC address is ", addressStr)
	} else if flag == 3 {
		private_wif, err := btcutil.NewWIF(private_key, &dogeAddressNetParams, true)
		private_str := private_wif.String()
		address_str, err := child.Address(&dogeAddressNetParams)
		if err != nil {
			return "", "", err
		}
		addressStr = address_str.String()
		fmt.Println("The DOGE private wif key is ", private_str)
		fmt.Println("The DOGE address is ", addressStr)
	} else if flag == 4 {
		//private_wif, err := btcutil.NewWIF(private_key, &qtumAddressNetParams, true)
		//private_str := private_wif.String()
		address_str, err := child.Address(&qtumAddressNetParams)
		if err != nil {
			return "", "", err
		}
		addressStr = address_str.String()
		//fmt.Println("The QTUM private wif key is ", private_str)
		fmt.Println("The QTUM address is ", addressStr)
	} else if flag == 5 {
		addressStr = nuls.Address(child)
		if len(addressStr) <= 0 {
			return "", "", err
		}
		fmt.Println("The Nuls address is ", addressStr)
		//fmt.Printf("The Nuls priKey：%x ", private_key.Serialize())
	} else {
		//not support
		return "", "", errors.New("The CoinType is not supported!!")
	}
	privateKeyStr := child.String()
	return addressStr, privateKeyStr, nil
}
