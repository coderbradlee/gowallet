package hdwallet

import (
	// "bytes"
	// "crypto/aes"
	// "crypto/cipher"
	// "crypto/ecdsa"
	// "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	// "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	// "io"
	// "io/ioutil"
	// "math/big"

	// "encoding/binary"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/debian-go/golang-github-nebulouslabs-entropy-mnemonics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"hdwallet/nuls"
)

// Key derives a key from the password, salt, and cost parameters, returning
// a byte slice of length keyLen that can be used as cryptographic key.
//
// N is a CPU/memory cost parameter, which must be a power of two greater than 1.
// r and p must satisfy r * p < 2³⁰. If the parameters do not satisfy the
// limits, the function returns a nil byte slice and an error.
//
// For example, you can get a derived key for e.g. AES-256 (which needs a
// 32-byte key) by doing:
//
//      dk, err := scrypt.Key([]byte("some password"), salt, 16384, 8, 1, 32)
//
// The recommended parameters for interactive logins as of 2009 are N=16384,
// r=8, p=1. They should be increased as memory latency and CPU parallelism
// increases. Remember to get a good random salt.
func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error) {
	if N <= 1 || N&(N-1) != 0 {
		return nil, errors.New("scrypt: N must be > 1 and a power of 2")
	}
	if uint64(r)*uint64(p) >= 1<<30 || r > maxInt/128/p || r > maxInt/256 || N > maxInt/128/r {
		return nil, errors.New("scrypt: parameters are too large")
	}

	xy := make([]uint32, 64*r)
	v := make([]uint32, 32*N*r)
	b := pbkdf2.Key(password, salt, 1, p*128*r, sha256.New)

	for i := 0; i < p; i++ {
		smix(b[i*128*r:], r, N, v, xy)
	}

	return pbkdf2.Key(password, b, 1, keyLen, sha256.New), nil
}

// Mnemonic Import
func importMnemonic(mnemonic string) ([]byte, error) {
	return mnemonics.FromString(mnemonic, mnemonics.English)
}

// Mnemonic Generation
func generateMnemonic(entropy []byte) (string, error) {
	if len(entropy) < 0 {
		entropy, _ = bip39.NewEntropy(192)
	}
	mnemonic, err := mnemonics.ToPhrase(entropy, mnemonics.English)
	return mnemonic.String(), err
}
func generateMasterkey(masterSeed []byte) (string, error) {
	masterKey, err := hdkeychain.NewMaster(masterSeed, &btcAddressNetParams)
	return masterKey.String(), err
}

//func CreateWalletByteRandAndPwd(random []byte, password string) (masterKey, mnemonic string, err error))
func CreateWalletByteRandAndPwd(rand string, password string) (masterKeyWithmnemonic string, err error) {
	//var seed []byte
	random := []byte(rand)
	if len(random) <= 0 {
		random, _ = bip39.NewEntropy(192)
	}

	seed, err := generateSeed(random, []byte(password))
	seedLen = len(seed)
	//fmt.Println("The Real seed len is :", seedLen)
	//fmt.Println("The Real seed to byte is: #v%", seed)
	if err != nil {
		return "", err
	}
	//Create Mnemonic
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
	//fmt.Println("The origianl masterky is---->", masterKeyStr)
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

func newWalletAccount(secret, salt []byte) (addressWithPrivateKey string, err error) {
	//wa = new(WalletAccount)
	var seed []byte
	seed, err = generateSeed(secret, salt)
	fmt.Println("The Real seed to byte is: #v%", seed)
	if err != nil {
		return
	}
	//Create Mnemonic
	mnemonic, err := generateMnemonic(seed)
	if err != nil {
		return "", err
	}
	fmt.Println("The mnemonic word list is ", mnemonic)

	//Import Mnemonic
	mnemonicSeed, err := importMnemonic(mnemonic)
	if err != nil {
		return "", err
	}
	fmt.Println("The mnemonic word list to byte is: #v%", mnemonicSeed)
	masterKeyStr, err := generateMasterkey(seed)
	//master_key, err := hdkeychain.NewMaster(seed, &AddressNetParams)
	if err != nil {
		return "", err
	}
	//err = wa.generateAccount(master_key.String(), 0)
	//err = wa.GenerateBIP44AccountPath(master_key.String(), "BTC", 0, 0, 0)
	return GenerateBIP44AccountWallet(masterKeyStr, "ETC", 0, 0, 0)

}

// Generate wallet seed from secret and salt
func generateSeed(secret, salt []byte) (seed []byte, err error) {
	// WarpWallet encryption:
	// 1. s1 ← scrypt(key=passphrase||0x1, salt=salt||0x1, N=218, r=8, p=1, dkLen=32)
	// 2. s2 ← PBKDF2(key=passphrase||0x2, salt=salt||0x2, c=216, dkLen=32)
	// 3. private_key ← s1 ⊕ s2
	// 4. Generate public_key from private_key using standard Bitcoin EC crypto
	// 5. Output (private_key, public_key)

	if len(secret) == 0 {
		err = errors.New("empty secret")
		return
	}
	if len(salt) == 0 {
		err = errors.New("empty salt")
		return
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
		return
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
	seed_hash = sha256.Sum256(seed_hash[:])
	seed = seed_hash[:24]
	return
}

// Generate BIP44 account extended private key and extended public key by the path.
//(wa *WalletAccount) generateAccount(masterKey string, k uint32)
//func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int) (address, privateKey string, err error)
func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int) (address string, err error) {
	if (len(masterKey) == 0) || (len(coinType) == 0) {
		return "", errors.New("some params is empty!!!")
	}
	//Decrypt the masterkey
	decMasterkey, err := decryptMasterkey(masterKey)
	if err != nil {
		return "", err
	}
	//fmt.Println("The decrypt masterky is---->", decMasterkey)
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
		fmt.Println("The ETH/ETC/ETF address is ", addressStr)
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
