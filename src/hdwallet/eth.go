package hdwallet

import (
	"encoding/hex"
	"errors"
	"fmt"
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

//Send ETH/ETC RawTransaction by the privateKey
//Send ETH/ETC RawTransaction   amount, gasLimit, gasPrice *big.Int
func SendETHRawTxByPrivateKey(privateKey, nonce, toAddr, amount, gasLimit, gasPrice string, data []byte) (signedParam string, err error) {
	if (len(privateKey) == 0) || (len(nonce) == 0) || (len(toAddr) == 0) {
		return "", errors.New("some params length is 0")
	}

	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return "", err
	}
	address := make([]byte, 20)
	//toAddr need to check the prefix 0x
	find := (strings.HasPrefix(toAddr, "0x")) || (strings.HasPrefix(toAddr, "0X"))
	var tempAddr string
	if find {
		tempAddr = toAddr[2:len(toAddr)]
	}
	address, err = hex.DecodeString(tempAddr)
	if err != nil {
		return "", err
	}
	var toAddress common.Address
	for i := 0; i < len(address); i++ {
		toAddress[i] = address[i]
	}
	//Convert to big int
	nonceUint64, err := strconv.ParseUint(nonce, 10, 64)
	if err != nil {
		return "", err
	}
	amountInt := new(big.Int)
	amountInt, ok := amountInt.SetString(amount, 10)
	if !ok {
		//fmt.Println("Value set to Big Int Wrong!!", amountInt)
		return "", errors.New("Value set to Big Int Wrong!!")
	}

	/*gasLimitInt := new(big.Int)
	gasLimitInt, ok = gasLimitInt.SetString(gasLimit, 10)
	if !ok {
		//fmt.Println("Value set to Big Int Wrong!!", amountInt)
		return "", errors.New("Gas Limit set to Big Int Wrong!!")
	}*/
	gasLimitInt, _ := strconv.ParseUint(gasLimit, 10, 64)

	gasPriceInt := new(big.Int)
	gasPriceInt, ok = gasPriceInt.SetString(gasPrice, 10)
	if !ok {
		//fmt.Println("Value set to Big Int Wrong!!", amountInt)
		return "", errors.New("Gas Limit set to Big Int Wrong!!")
	}

	tx := types.NewTransaction(nonceUint64, toAddress, amountInt, gasLimitInt, gasPriceInt, data)
	private_key, _ := child.ECPrivKey()
	signed, err := types.SignTx(tx, types.HomesteadSigner{}, (*ecdsa.PrivateKey)(private_key))
	if err != nil {
		return "", err
	}

	encodeData, err := rlp.EncodeToBytes(signed)
	if err != nil {
		return "", err
	}
	hexString := common.ToHex(encodeData)

	fmt.Println("The real ETH sig is ", signed)
	fmt.Println("The real signed hex string is ", hexString)
	return hexString, err
}

func addressforEth(k *hdkeychain.ExtendedKey) ([]byte, error) {
	publickey, _ := k.ECPubKey()
	var p *ecdsa.PublicKey
	p = (*ecdsa.PublicKey)(publickey)
	pubBytes := crypto.FromECDSAPub(p)
	pkPrv := common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
	pkHash := pkPrv[:]
	return pkHash, nil
}
