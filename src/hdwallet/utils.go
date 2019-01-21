package hdwallet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"

	"encoding/binary"
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

type WalletAccount struct {
	MasterKey string
	Mnemonic  string
}

type addressToKey struct {
	key        *btcec.PrivateKey
	compressed bool
}

// AuthoredTx holds the state of a newly-created transaction and the change
// output (if one was added).
type AuthoredTx struct {
	Tx          *wire.MsgTx
	PrevScripts []string
}

//Get BTC balance Info
type ChainBalanceInfo struct {
	Status string           `json:"status"`
	Data   ChainBalanceData `json:"data"`
}

type ChainBalanceData struct {
	Network             string `json:"network"`
	Address             string `json:"address"`
	Confirmed_balance   string `json:"confirmed_balance"`
	Unconfirmed_balance string `json:"unconfirmed_balance"`
}

// VinPrevOut is like Vin except it includes PrevOut.  It is used by searchrawtransaction
type ChainUnspentInfo struct {
	Status string           `json:"status"`
	Data   ChainUnspentData `json:"data"`
}
type ChainUnspentData struct {
	Network string            `json:"network"`
	Address string            `json:"address"`
	Txs     []TxChainDataInfo `json:"txs"`
}

type QutmDataInfo struct {
	Address       string  `json:"address"`
	Txid          string  `json:"txid"`
	Vout          int64   `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Satoshis      int64   `json:"satoshis"`
	IsStake       bool    `json:"isStake"`
	Height        int64   `json:"height"`
	Confirmations int64   `json:"confirmations"`
}

type TxChainDataInfo struct {
	Txid          string `json:"txid"`
	Vout          int64  `json:"output_no"`
	Script_asm    string `json:"script_asm"`
	Script_hex    string `json:"script_hex"`
	Value         string `json:"value"`
	Confirmations int64  `json:"confirmations"`
	Time          int64  `json:"time"`
}

func flashBackInt(data []byte) []byte {
	len := len(data)
	s := make([]byte, len)
	i := 0
	for index := len - 1; index >= 0; index-- {
		if data[index] != 0 {
			s[i] = data[index]
			i++
		}
	}
	rec := make([]byte, i)
	for index := 0; index < i; index++ {
		rec[index] = s[index]
	}
	return rec
}
func flashBackString(data []byte) []byte {
	len := len(data)
	s := make([]byte, len)
	i := 0
	for index := len - 1; index >= 0; index-- {
		s[i] = data[index]
		i++
	}
	fmt.Println(data)
	fmt.Println(s)
	return s
}
func mkGetKey(keys map[string]addressToKey) txscript.KeyDB {
	if keys == nil {
		return txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey,
			bool, error) {
			return nil, false, errors.New("nope")
		})
	}
	return txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey,
		bool, error) {
		a2k, ok := keys[addr.EncodeAddress()]
		if !ok {
			return nil, false, errors.New("nope")
		}
		return a2k.key, a2k.compressed, nil
	})
}

func mkGetScript(scripts map[string][]byte) txscript.ScriptDB {
	if scripts == nil {
		return txscript.ScriptClosure(func(addr btcutil.Address) ([]byte, error) {
			return nil, errors.New("nope")
		})
	}
	return txscript.ScriptClosure(func(addr btcutil.Address) ([]byte, error) {
		script, ok := scripts[addr.EncodeAddress()]
		if !ok {
			return nil, errors.New("nope")
		}
		return script, nil
	})
}

// blockCopy copies n numbers from src into dst.
func blockCopy(dst, src []uint32, n int) {
	copy(dst, src[:n])
}

// blockXOR XORs numbers from dst with n numbers from src.
func blockXOR(dst, src []uint32, n int) {
	for i, v := range src[:n] {
		dst[i] ^= v
	}
}

// salsaXOR applies Salsa20/8 to the XOR of 16 numbers from tmp and in,
// and puts the result into both both tmp and out.
func salsaXOR(tmp *[16]uint32, in, out []uint32) {
	w0 := tmp[0] ^ in[0]
	w1 := tmp[1] ^ in[1]
	w2 := tmp[2] ^ in[2]
	w3 := tmp[3] ^ in[3]
	w4 := tmp[4] ^ in[4]
	w5 := tmp[5] ^ in[5]
	w6 := tmp[6] ^ in[6]
	w7 := tmp[7] ^ in[7]
	w8 := tmp[8] ^ in[8]
	w9 := tmp[9] ^ in[9]
	w10 := tmp[10] ^ in[10]
	w11 := tmp[11] ^ in[11]
	w12 := tmp[12] ^ in[12]
	w13 := tmp[13] ^ in[13]
	w14 := tmp[14] ^ in[14]
	w15 := tmp[15] ^ in[15]

	x0, x1, x2, x3, x4, x5, x6, x7, x8 := w0, w1, w2, w3, w4, w5, w6, w7, w8
	x9, x10, x11, x12, x13, x14, x15 := w9, w10, w11, w12, w13, w14, w15

	for i := 0; i < 8; i += 2 {
		u := x0 + x12
		x4 ^= u<<7 | u>>(32-7)
		u = x4 + x0
		x8 ^= u<<9 | u>>(32-9)
		u = x8 + x4
		x12 ^= u<<13 | u>>(32-13)
		u = x12 + x8
		x0 ^= u<<18 | u>>(32-18)

		u = x5 + x1
		x9 ^= u<<7 | u>>(32-7)
		u = x9 + x5
		x13 ^= u<<9 | u>>(32-9)
		u = x13 + x9
		x1 ^= u<<13 | u>>(32-13)
		u = x1 + x13
		x5 ^= u<<18 | u>>(32-18)

		u = x10 + x6
		x14 ^= u<<7 | u>>(32-7)
		u = x14 + x10
		x2 ^= u<<9 | u>>(32-9)
		u = x2 + x14
		x6 ^= u<<13 | u>>(32-13)
		u = x6 + x2
		x10 ^= u<<18 | u>>(32-18)

		u = x15 + x11
		x3 ^= u<<7 | u>>(32-7)
		u = x3 + x15
		x7 ^= u<<9 | u>>(32-9)
		u = x7 + x3
		x11 ^= u<<13 | u>>(32-13)
		u = x11 + x7
		x15 ^= u<<18 | u>>(32-18)

		u = x0 + x3
		x1 ^= u<<7 | u>>(32-7)
		u = x1 + x0
		x2 ^= u<<9 | u>>(32-9)
		u = x2 + x1
		x3 ^= u<<13 | u>>(32-13)
		u = x3 + x2
		x0 ^= u<<18 | u>>(32-18)

		u = x5 + x4
		x6 ^= u<<7 | u>>(32-7)
		u = x6 + x5
		x7 ^= u<<9 | u>>(32-9)
		u = x7 + x6
		x4 ^= u<<13 | u>>(32-13)
		u = x4 + x7
		x5 ^= u<<18 | u>>(32-18)

		u = x10 + x9
		x11 ^= u<<7 | u>>(32-7)
		u = x11 + x10
		x8 ^= u<<9 | u>>(32-9)
		u = x8 + x11
		x9 ^= u<<13 | u>>(32-13)
		u = x9 + x8
		x10 ^= u<<18 | u>>(32-18)

		u = x15 + x14
		x12 ^= u<<7 | u>>(32-7)
		u = x12 + x15
		x13 ^= u<<9 | u>>(32-9)
		u = x13 + x12
		x14 ^= u<<13 | u>>(32-13)
		u = x14 + x13
		x15 ^= u<<18 | u>>(32-18)
	}
	x0 += w0
	x1 += w1
	x2 += w2
	x3 += w3
	x4 += w4
	x5 += w5
	x6 += w6
	x7 += w7
	x8 += w8
	x9 += w9
	x10 += w10
	x11 += w11
	x12 += w12
	x13 += w13
	x14 += w14
	x15 += w15

	out[0], tmp[0] = x0, x0
	out[1], tmp[1] = x1, x1
	out[2], tmp[2] = x2, x2
	out[3], tmp[3] = x3, x3
	out[4], tmp[4] = x4, x4
	out[5], tmp[5] = x5, x5
	out[6], tmp[6] = x6, x6
	out[7], tmp[7] = x7, x7
	out[8], tmp[8] = x8, x8
	out[9], tmp[9] = x9, x9
	out[10], tmp[10] = x10, x10
	out[11], tmp[11] = x11, x11
	out[12], tmp[12] = x12, x12
	out[13], tmp[13] = x13, x13
	out[14], tmp[14] = x14, x14
	out[15], tmp[15] = x15, x15
}

func blockMix(tmp *[16]uint32, in, out []uint32, r int) {
	blockCopy(tmp[:], in[(2*r-1)*16:], 16)
	for i := 0; i < 2*r; i += 2 {
		salsaXOR(tmp, in[i*16:], out[i*8:])
		salsaXOR(tmp, in[i*16+16:], out[i*8+r*16:])
	}
}

func integer(b []uint32, r int) uint64 {
	j := (2*r - 1) * 16
	return uint64(b[j]) | uint64(b[j+1])<<32
}

func smix(b []byte, r, N int, v, xy []uint32) {
	var tmp [16]uint32
	x := xy
	y := xy[32*r:]

	j := 0
	for i := 0; i < 32*r; i++ {
		x[i] = uint32(b[j]) | uint32(b[j+1])<<8 | uint32(b[j+2])<<16 | uint32(b[j+3])<<24
		j += 4
	}
	for i := 0; i < N; i += 2 {
		blockCopy(v[i*(32*r):], x, 32*r)
		blockMix(&tmp, x, y, r)

		blockCopy(v[(i+1)*(32*r):], y, 32*r)
		blockMix(&tmp, y, x, r)
	}
	for i := 0; i < N; i += 2 {
		j := int(integer(x, r) & uint64(N-1))
		blockXOR(x, v[j*(32*r):], 32*r)
		blockMix(&tmp, x, y, r)

		j = int(integer(y, r) & uint64(N-1))
		blockXOR(y, v[j*(32*r):], 32*r)
		blockMix(&tmp, y, x, r)
	}
	j = 0
	for _, v := range x[:32*r] {
		b[j+0] = byte(v >> 0)
		b[j+1] = byte(v >> 8)
		b[j+2] = byte(v >> 16)
		b[j+3] = byte(v >> 24)
		j += 4
	}
}

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

// 该方法已经删除
func getChildByPath(masterKey string, coinType string, account, change, index int) (child *hdkeychain.ExtendedKey, err error) {
	master_key, err := hdkeychain.NewKeyFromString(masterKey)
	var drivedCoinType *hdkeychain.ExtendedKey
	if err != nil {
		return drivedCoinType, err
	}
	purpose, err := master_key.Child(hardened + 44)
	if err != nil {
		return drivedCoinType, err
	}

	//Coin type: maybe changed by different coin type
	if coinType == "BTC" {
		drivedCoinType, err = purpose.Child(hardened + 0)
	} else if coinType == "ETH" {
		drivedCoinType, err = purpose.Child(hardened + 60)
	} else if coinType == "ETC" {
		drivedCoinType, err = purpose.Child(hardened + 61)
	} else if coinType == "ETF" {
		drivedCoinType, err = purpose.Child(hardened + 62)
	} else if coinType == "LTC" {
		drivedCoinType, err = purpose.Child(hardened + 2)
	} else {
		return drivedCoinType, errors.New("The Coin Type is not support!!!")
	}
	if err != nil {
		return drivedCoinType, err
	}
	//account
	drivedAccount, err := drivedCoinType.Child(hardened + (uint32)(account))
	if err != nil {
		return drivedCoinType, err
	}

	//Change(T/F:1,0)
	change = 0
	drivedChange, err := drivedAccount.Child((uint32)(change))
	if err != nil {
		return drivedCoinType, err
	}
	//create change Index
	index = 0
	return drivedChange.Child((uint32)(index))
}

func getChildByPrivatekeyStr(privateKey string) (child *hdkeychain.ExtendedKey, err error) {
	private_key, err := hdkeychain.NewKeyFromString(privateKey)
	return private_key, err
}

func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func decodeBase64(b []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		// fmt.Printf("Error: Bad Key!\n")
		// os.Exit(0)
		return []byte(""), err
	}
	return data, nil
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte(""), err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext, nil
}
func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	if len(text) < aes.BlockSize {
		return []byte(""), errors.New("len(text) < aes.BlockSize")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	// fmt.Println("55")
	cfb.XORKeyStream(text, text)
	// fmt.Println("55")
	return decodeBase64(text)
}
func byteSliceEqual(a, b []byte) bool {
	//fmt.Println("a:", a)
	//fmt.Println("b:", b)
	if len(a) != len(b) {
		// fmt.Println("len")
		return false
	}

	if (a == nil) != (b == nil) {
		// fmt.Println("nil")
		return false
	}

	for i, v := range a {
		if v != b[i] {
			// fmt.Println("range")
			return false
		}
	}

	return true
}

//用密码对masterkey加密，对加密后的文本在app端保存
func encryptMastkeyWithPwd(masterKeyStr, passwordStr string) (string, error) {
	//fmt.Print("The masterkey is->", masterKeyStr)
	//fmt.Println("--->The Password is ", passwordStr)
	key := []byte(passwordStr)
	text := []byte(masterKeyStr)
	hashKey := sha256.Sum256(key)
	encryptStr := make([]byte, len(text)+len(hashKey))

	copy(encryptStr[:len(text)], text[:])
	copy(encryptStr[len(text):], hashKey[:])

	suffix := sha256.Sum256(encryptStr)
	constEncKey := sha256.Sum256([]byte(constEncKeyStr))
	prefix, err := encrypt(constEncKey[:], encryptStr)
	if err != nil {
		return "", err
	}
	ret := make([]byte, len(prefix)+len(suffix))
	copy(ret[:len(prefix)], prefix[:])
	copy(ret[len(prefix):], suffix[:])
	return hex.EncodeToString(ret[:]), nil
}

//用密码对密文解密返回masterkey对应的byte数组
func decryptMasterkey(encryptMasterkeyStr string) (string, error) {
	text, err := hex.DecodeString(encryptMasterkeyStr)
	if err != nil {
		return "", err
	}
	constEncKey := sha256.Sum256([]byte(constEncKeyStr))
	d_des, err := decrypt(constEncKey[:], text[:len(text)-len(constEncKey)])
	if err != nil {
		return "", err
	}
	return string(d_des[:len(d_des)-len(constEncKey)-seedLen]), nil
}

//用密码对密文解密返回masterkeywithseed对应的byte数组
func decryptMasterkeyWithSeed(encryptMasterkeyStr string) (string, error) {
	text, err := hex.DecodeString(encryptMasterkeyStr)
	if err != nil {
		return "", err
	}
	constEncKey := sha256.Sum256([]byte(constEncKeyStr))
	d_des, err := decrypt(constEncKey[:], text[:len(text)-len(constEncKey)])
	if err != nil {
		return "", err
	}
	return string(d_des[:len(d_des)-len(constEncKey)]), nil
}

//用密码对密文解密返回masterkey对应的byte数组
func decryptSeed(encryptMasterkeyStr string) ([]byte, error) {
	text, err := hex.DecodeString(encryptMasterkeyStr)
	var seed []byte
	if err != nil {
		return seed, err
	}
	constEncKey := sha256.Sum256([]byte(constEncKeyStr))
	d_des, err := decrypt(constEncKey[:], text[:len(text)-len(constEncKey)])
	if err != nil {
		return seed, err
	}
	tempSeed := (d_des[len(d_des)-len(constEncKey)-seedLen : len(d_des)-len(constEncKey)])
	//fmt.Println("The  decryptSeed is #v", tempSeed)
	return tempSeed, nil
}

//用文本来验证密码是否正确
func CheckPwdIsCorrect(masterKeyStr, passwordStr string) (right bool) {
	if (len(masterKeyStr) == 0) || (len(passwordStr) == 0) {
		return false
	}
	dec, err := decryptMasterkeyWithSeed(masterKeyStr)
	if err != nil {
		return false
	}
	d_des := []byte(dec)
	key := []byte(passwordStr)
	//text := masterKeyStr
	text, err := hex.DecodeString(masterKeyStr)
	if err != nil {
		return false
	}
	hashKey := sha256.Sum256(key)
	hashstr := make([]byte, len(d_des)+len(hashKey))
	copy(hashstr[:len(d_des)], d_des[:])
	copy(hashstr[len(d_des):], hashKey[:])

	hash_des := sha256.Sum256(hashstr[:])
	return byteSliceEqual(text[len(text)-len(hashKey):], hash_des[:])
}

//Backup the Mnemonic
func BackupMnemonic(masterKeyStr string) (mnemonic string, err error) {
	if len(masterKeyStr) == 0 {
		return "", errors.New("some params is empty!!!")
	}
	seed, err := decryptSeed(masterKeyStr)
	if err != nil {
		return "", err
	}
	//Create Mnemonic
	mnemonic, err = generateMnemonic(seed)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

//func CheckAuthAndGetPrivateKey(masterKey string, coinType string, account, change, index int) ( privateKey string, err error)
func CheckAuthAndGetPrivateKey(masterKey string, password string, coinType string, account, change, index int) (privateKey string, err error) {
	if (len(masterKey) == 0) || (len(password) == 0) || (len(coinType) == 0) {
		return "", errors.New("some params is empty!!!")
	}

	//Check the Auth
	right := CheckPwdIsCorrect(masterKey, password)
	if !right {
		return "", errors.New("You don't have the Authority to get Pass word!!!")
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
	change = 0
	drivedChange, err := drivedAccount.Child((uint32)(change))
	if err != nil {
		return "", err
	}
	//create change Index
	index = 0
	_, privateKey, err = createChangeIndex(drivedChange, index, flag)
	return privateKey, err

}

//
func GetHexAddress(qtumAddress string) string {
	buf := base58.Decode(qtumAddress)
	hexstr := hex.EncodeToString(buf[1:21])
	return hexstr
}

func GetDecAddress(qtumAddress string) string {
	buf := base58.Encode([]byte(qtumAddress))
	return buf
}
