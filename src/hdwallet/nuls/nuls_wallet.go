package nuls

import (
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"encoding/hex"
	"sort"
	"bytes"
)

const (
	httpReqTimeOut = 60 * time.Second
	nulsPrecision  = 100000000 //nuls精度
)

var (
	nulsNetChain          = []byte{4, 35}
	//nulsTestNetChain          = []byte{1, 0} //测试网络地址前缀
	nulsNetAddressType    = byte(1)
	nulsRpcNodes          = []string{"https://qtum-api.tokenxx.com"} //查询余额节点
	nulsRpcUtxoNodes      = []string{"https://qtum-api.tokenxx.com"} //查询utxo节点
	nulsRpcBroadcastNodes = []string{"https://qtum-api.tokenxx.com"} //nuls交易广播的交易节点
)

type resultModel struct {
	Success bool   `json:"success"`
	Faild   bool   `json:"faild"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
}

type balanceModel struct {
	resultModel
	Data balanceModelData `json:"data"`
}

type balanceModelData struct {
	Id          string `json:"id"`
	Address     string `json:"address"`
	Locked      uint64 `json:"locked"`
	Usable      uint64 `json:"usable"`
	BlockHeight uint64 `json:"blockHeight"`
	AssetsCode  uint64 `json:"assetsCode"`
}

type unspendModel struct {
	resultModel
	Data UnspendMap `json:"data"`
}
type UnspendMap struct {
	UtxoDtoList []UnspendModelData `json:"utxoDtoList"`
}

type UnspendModelData struct {
	TxHash   string `json:"txHash"`
	TxIndex  int    `json:"txIndex"`
	Amount   uint64 `json:"value"`
	LockTime uint64 `json:"lockTime"`
}

func (s *unspendModel) Len() int {
	return len(s.Data.UtxoDtoList)
}

func (s *unspendModel) Swap(i, j int) {
	s.Data.UtxoDtoList[i], s.Data.UtxoDtoList[j] = s.Data.UtxoDtoList[j], s.Data.UtxoDtoList[i]
}

func (s *unspendModel) Less(i, j int) bool {
	return s.Data.UtxoDtoList[i].Amount < s.Data.UtxoDtoList[j].Amount
}

type broadcastData struct {
	TxHex    string `json:"txHex"`
	Address  string `json:"address"`
	PriKey   string `json:"priKey"`
	Password string `json:"password"`
}

type broadcastResult struct {
	resultModel
	Data broadcastMSg `json:"data"`
}
type broadcastMSg struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Value string `json:"value"`
}

func Address(eckey *hdkeychain.ExtendedKey) string {
	eCPrivKey, _ := eckey.ECPrivKey()
	pkHash := btcutil.Hash160(eCPrivKey.PubKey().SerializeCompressed())
	b := make([]byte, 0, len(pkHash)+4)
	b = append(b, nulsNetChain...)
	b = append(b, nulsNetAddressType)
	b = append(b, pkHash[:]...)
	bodyB := b[0:23]

	orX := byte(0x00)
	for _, v := range bodyB {
		orX ^= v
	}
	b = append(b, byte(orX))
	nulsAdd := base58.Encode(b)
	return nulsAdd
}

/*
//从区块浏览器查询余额
func GetBalance(address string) (amount []string, err error) {
	client := &http.Client{
		Timeout: httpReqTimeOut,
	}
	var url string
	var tempErr error
	for i, _ := range nulsRpcNodes {
		url = fmt.Sprintf("%s/nuls/balance/get/%s", nulsRpcNodes[i], address)
		respJson, err := client.Get(url)
		if err != nil {
			tempErr = err
			continue
		}
		bs, err := ioutil.ReadAll(respJson.Body)
		respJson.Body.Close()
		var resData balanceModel
		err = json.Unmarshal(bs, &resData)
		if err != nil || !resData.Success {
			tempErr = errors.New(resData.Msg)
			continue
		}
		lockedAmount := strconv.FormatFloat(float64(resData.Data.Locked)/nulsPrecision, 'f', -1, 64)

		usableAmount := strconv.FormatFloat(float64(resData.Data.Usable)/nulsPrecision, 'f', -1, 64)

		amount = []string{usableAmount, lockedAmount}
		return amount, nil
	}
	return amount, tempErr
}*/

func GetBalance(address string) (amount []string, err error) {
	uxto, err := GetUtxoUnspent(address)
	if err != nil {
		return nil, err
	}
	total := uint64(0)
	for _, v := range uxto {
		total += v.Amount
	}

	usableAmount := strconv.FormatFloat(float64(total)/nulsPrecision, 'f', -1, 64)
	amount = []string{usableAmount, "0"}
	return amount, nil
}

func GetUtxoUnspent(address string) (unspent []UnspendModelData, err error) {
	client := &http.Client{
		Timeout: httpReqTimeOut,
	}
	var url string
	var tempErr error
	for i, _ := range nulsRpcUtxoNodes {
		url = fmt.Sprintf("%s/api/utxo/limit/%s/100000000", nulsRpcUtxoNodes[i], address)
		respJson, err := client.Get(url)
		if err != nil {
			tempErr = err
			continue
		}
		bs, err := ioutil.ReadAll(respJson.Body)
		respJson.Body.Close()
		var resData unspendModel
		err = json.Unmarshal(bs, &resData)
		if err != nil || !resData.Success {
			tempErr = errors.New(resData.Msg)
			if len(resData.Msg) == 0 {
				tempErr = errors.New("query fail")
			}
			continue
		}
		sort.Sort(&resData)
		return resData.Data.UtxoDtoList, nil
	}
	return nil, tempErr
}

func TransferFee(from, to, amount, price, remark string) (string, error) {
	return createTransfer(from, to, amount, price, remark, nil, true)
}

func Transfer(from, to, amount, price, remark string, ecKey *hdkeychain.ExtendedKey) (string, error) {
	return createTransfer(from, to, amount, price, remark, ecKey, false)
}

func createTransfer(from, to, amount, price, remark string, ecKey *hdkeychain.ExtendedKey, calcFee bool) (string, error) {
	from = strings.TrimSpace(from)
	to = strings.TrimSpace(to)
	amount = strings.TrimSpace(amount)
	remark = strings.TrimSpace(remark)

	feePrice := MIN_PRECE_PRE_1024_BYTES
	if len(price) > 0 {
		price = strings.TrimSpace(price)
		temp, err := strconv.ParseUint(price, 10, 64)
		if err != nil {
			return "", errors.New("Incorrect transfer price format")
		}
		feePrice = int(temp)
	}

	//去掉最后一位的校验码
	fromB, err := decodeAddress(from)
	if err != nil {
		return "", err
	}
	toB, err := decodeAddress(to)
	if err != nil {
		return "", err
	}

	//转账金额格式转换
	transAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", errors.New("Incorrect transfer amount format")
	}
	//查询地址可用金额
	fromAmount, err := GetBalance(from)
	if err != nil {
		return "", errors.New("Incorrect get address amount")
	}
	usableAmount, _ := strconv.ParseFloat(fromAmount[0], 64)
	if usableAmount < transAmount {
		return "", errors.New("balance is not enough")
	}

	//组装交易
	time := time.Now().UnixNano() / 1000000
	t := Transaction{
		Types:  TX_TYPE_TRANSFER,
		time:   uint64(time),
		remark: []byte(remark),
	}
	transAmountNa := uint64(transAmount * nulsPrecision) //转换为最小单位na
	coinData := CoinData{}
	coinData.AddTo(Coin{owner: toB, na: Na(transAmountNa)})

	//查询未花费的交易
	fromUnspent, err := GetUtxoUnspent(from)
	if err != nil {
		return "", errors.New("Incorrect not get utxo")
	}
	if len(fromUnspent) <= 0 {
		return "", errors.New("balance is not enough")
	}
	txsize := t.Size() + DEFAULT_SERIALIZE_LENGTH //必须包含签名脚本

	var values uint64 = 0
	fee := Na(0)
	for i, v := range fromUnspent {
		if v.Amount == 0 {
			continue
		}
		hash := make([]byte, len(v.TxHash)/2)
		hex.Decode(hash, []byte(v.TxHash))
		index := VarInt(v.TxIndex);
		hash = append(hash, index.encode()...)
		input := Coin{owner: hash, na: Na(v.Amount),}
		coinData.AddFrom(input)
		txsize += input.Size()
		if i == 127 {
			txsize += 1
		}

		fee, err = getFee(txsize, Na(feePrice))
		values += v.Amount
		if err != nil {
			return "", err
		}
		//余额足够后，需要判断是否找零，如果有找零，则需要重新计算手续费
		if values >= transAmountNa+uint64(fee) {
			change := values - transAmountNa - uint64(fee)
			if change > 0 {
				changeOut := Coin{owner: fromB, na: Na(change),}
				fee, _ = getFee(txsize+changeOut.Size(), Na(feePrice))
				if values < transAmountNa+uint64(fee) {
					continue
				}
				coinData.AddTo(changeOut)
			}
			break
		}

	}
	if calcFee {
		return strconv.FormatFloat(float64(fee)/float64(nulsPrecision), 'f', -1, 64), nil
	}
	if transAmountNa+uint64(fee) > uint64(usableAmount*nulsPrecision) {
		return "", errors.New("balance is not enough pay fee")
	}
	t.coinData = coinData
	t.setHash(*calcDigestData(t.SerializeForHash(), 0))
	ecPubKey, _ := ecKey.ECPubKey()
	enPriKey, _ := ecKey.ECPrivKey()

	sig, err := enPriKey.Sign(t.getHash().digestBytes)
	script := P2PKHScriptSig{publicKey: ecPubKey.SerializeCompressed(),
		signData: NulsSignData{signAlgType: SIGN_ALG_ECC, signBytes: sig.Serialize()},
	}
	t.ScriptSig, _ = script.SerializeToByte()

	tbytes, err := t.SerializeToByte()
	if err != nil {
		return "", err
	}
	json, err := json.Marshal(t)
	fmt.Println(string(json)) //打印交易信息
	thp := t.getHash()
	h, _ := thp.SerializeToByte()
	fmt.Printf("%x", h) //打印交易信息
	return hex.EncodeToString(tbytes), nil
}

func Broadcast(txHex string) (string, error) {
	if len(txHex) == 0 {
		return "-1", errors.New("tx data null")
	}
	jsonStr, err := json.Marshal(broadcastData{TxHex: txHex})
	if err != nil {
		return "-1", err
	}
	client := &http.Client{
		Timeout: httpReqTimeOut,
	}
	var url string
	var tempErr error
	for i, _ := range nulsRpcBroadcastNodes {
		url = fmt.Sprintf("%s/api/accountledger/transaction/broadcast", nulsRpcBroadcastNodes[i])
		bf := bytes.NewBuffer(jsonStr)
		respJson, err := client.Post(url, "application/json", bf)
		if err != nil {
			tempErr = err
			continue
		}
		bs, err := ioutil.ReadAll(respJson.Body)
		respJson.Body.Close()
		var resData broadcastResult
		err = json.Unmarshal(bs, &resData)
		if err != nil || !resData.Success {
			tempErr = errors.New(fmt.Sprintf("%s(%s)", resData.Data.Code, resData.Data.Msg))
			continue
		}
		return resData.Data.Value, nil
	}
	return "", tempErr
}

func decodeAddress(address string) (b []byte, err error) {
	if len(address) != 32 {
		return nil, errors.New("Invalid address")
	}
	defer func() {
		if recover() != nil {
			err = errors.New("Invalid address")
		}
	}()
	ad := base58.Decode(address)
	if len(ad)!=24{
		return nil,errors.New("Invalid address")
	}
	return ad[0:23], nil
}
