package hdwallet

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	// "github.com/btcsuite/btcd/btcec"
	"math/big"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func (hd *Hdwallet) ethAddress(child *hdkeychain.ExtendedKey) (address string, err error) {
	// private_key, err := child.ECPrivKey()
	// fmt.Printf("%x", private_key.D.Bytes())
	// if err != nil {
	// 	return
	// }
	// privateKeyBytes := private_key.Serialize()
	// private_str = hex.EncodeToString(privateKeyBytes) //wif格式的私钥
	ethaddress_key, err := addressforEth(child)
	if err != nil {
		return
	}
	address = hex.EncodeToString(ethaddress_key)
	return
}

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

func GetBalance(addr string) (balance string, err error) {
	// Request
	//	curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xc94770007dda54cF92009BFF0dE90c06F603a09f", "latest"],"id":1}'
	//
	//// Result
	//{
	//"id":1,
	//"jsonrpc": "2.0",
	//"result": "0x0234c8a3397aab58" // 158972490234375000
	//}
	url := "https://mainnet.infura.io/v3/33d5efb02b384213b386b08b324cdaaa"
	res, err := doPost(url, "eth_getBalance", []string{addr, "latest"})
	if err != nil {
		return
	}
	//b, err := res.Result.MarshalJSON()
	//if err != nil {
	//	return
	//}
	//fmt.Println("b:", b)
	err = json.Unmarshal(res.Result[:], &balance)
	return
}

func doPost(url string, method string, params interface{}) (*JSONRpcResp, error) {
	client := &http.Client{
		Timeout: requestTimeout,
	}
	jsonReq := map[string]interface{}{"jsonrpc": "2.0", "method": method, "params": params, "id": 0}
	data, _ := json.Marshal(jsonReq)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Length", (string)(len(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rpcResp *JSONRpcResp
	err = json.NewDecoder(resp.Body).Decode(&rpcResp)
	if err != nil {
		return nil, err
	}
	if rpcResp.Error != nil {
		return nil, errors.New(rpcResp.Error["message"].(string))
	}
	return rpcResp, err
}

type JSONRpcResp struct {
	Id      *json.RawMessage       `json:"id,omitempty"`
	Result  *json.RawMessage       `json:"result,omitempty"`
	Error   map[string]interface{} `json:"error,omitempty"`
	Jsonrpc string                 `json:"jsonrpc,omitempty"`
}
