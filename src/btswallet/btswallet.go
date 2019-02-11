package btswallet

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	//eoseos "github.com/eoscanada/eos-go"
	eosbtcec "github.com/eoscanada/eos-go/btcsuite/btcd/btcec"
	eosbtcutil "github.com/eoscanada/eos-go/btcsuite/btcutil"
	eosecc "github.com/eoscanada/eos-go/ecc"
)

var testvar = "bts"

//转账交易结构
type Transfer struct {
	Blocknum      uint16
	Blockprefix   uint32
	Expiration    uint32
	opsize        byte
	optype        byte
	free          uint64
	freeasserid   byte
	from          []byte
	to            []byte
	amount        uint64
	amountasserid byte
}

//获取bts公私钥对
func GetBtsKey(User string, Password string) (string, error) {
	if User == "" || len(User) == 0 {
		err := errors.New("user is null")
		return "", err
	}
	if Password == "" || len(Password) == 0 {
		err := errors.New("password is null")
		return "", err
	}
	var rebuf string
	//active权限
	{
		//种子
		randKey := User + "active" + Password
		h := sha256.New()
		h.Write([]byte(randKey))
		priKey, _ := eosbtcec.PrivKeyFromBytes(elliptic.P256(), h.Sum(nil))
		wif, _ := eosbtcutil.NewWIF(priKey, '\x80', false)
		//转换出公钥
		privKeyWif, _ := eosecc.NewPrivateKey(wif.String())
		pubKey := privKeyWif.PublicKey()
		pubKeyStr := pubKey.String()

		rebuf = privKeyWif.String() + "|" + strings.Replace(pubKeyStr, "EOS", "BTS", -1)
	}
	//memo权限(暂时跟active权限一致)
	{
		//种子
		randKey := User + "active" + Password
		h := sha256.New()
		h.Write([]byte(randKey))
		priKey, _ := eosbtcec.PrivKeyFromBytes(elliptic.P256(), h.Sum(nil))
		wif, _ := eosbtcutil.NewWIF(priKey, '\x80', false)
		//转换出公钥
		privKeyWif, _ := eosecc.NewPrivateKey(wif.String())
		pubKey := privKeyWif.PublicKey()
		pubKeyStr := pubKey.String()
		strings.Replace(pubKeyStr, "EOS", "BTS", 0)
		rebuf += "|" + privKeyWif.String() + "|" + strings.Replace(pubKeyStr, "EOS", "BTS", -1)
	}
	//owner权限
	{
		//种子
		randKey := User + "owner" + Password
		h := sha256.New()
		h.Write([]byte(randKey))
		priKey, _ := eosbtcec.PrivKeyFromBytes(elliptic.P256(), h.Sum(nil))
		wif, _ := eosbtcutil.NewWIF(priKey, '\x80', false)
		//转换出公钥
		privKeyWif, _ := eosecc.NewPrivateKey(wif.String())
		pubKey := privKeyWif.PublicKey()
		pubKeyStr := pubKey.String()
		strings.Replace(pubKeyStr, "EOS", "BTS", -1)
		rebuf += "|" + privKeyWif.String() + "|" + strings.Replace(pubKeyStr, "EOS", "BTS", -1)
	}
	return rebuf, nil
}

//获取对应权限的私钥
func GetBtsPriKey(User string, Role string, Password string) (string, error) {
	if User == "" || len(User) == 0 {
		err := errors.New("user is null")
		return "", err
	}
	if Role == "" || len(Role) == 0 {
		err := errors.New("role is null")
		return "", err
	}
	if Password == "" || len(Password) == 0 {
		err := errors.New("password is null")
		return "", err
	}
	//种子
	randKey := User + Role + Password
	h := sha256.New()
	h.Write([]byte(randKey))
	priKey, _ := eosbtcec.PrivKeyFromBytes(elliptic.P256(), h.Sum(nil))
	wif, _ := eosbtcutil.NewWIF(priKey, '\x80', false)
	return wif.String(), nil
}

//int转换为byte数组
func intToByte(lValue uint64) []byte {
	buf := new(bytes.Buffer)
	for true {
		if lValue <= 0 {
			break
		}
		b := byte(lValue & 0x7f)
		lValue >>= 7
		if lValue > 0 {
			b |= (1 << 7)
		}
		buf.Write([]byte{b}[0:1])
	}
	return buf.Bytes()
}

//字符串转换为byte
func stringToInt(buf string) int32 {
	//string转换为byte[]
	head_block_id, err := hex.DecodeString(buf)
	//构造
	if err != nil {
		fmt.Println("HashIndexToInt error:" + err.Error())
		return 0
	}
	//构造buffer对象
	bin_buf := bytes.NewBuffer(head_block_id)
	var x int32
	binary.Read(bin_buf, binary.BigEndian, &x)
	return x
}

//转账序列化
func transferSerial(chain_id string, transaction Transfer) []byte {

	bf := new(bytes.Buffer)
	chinid, err := hex.DecodeString(chain_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	bf.Write(chinid[0:32])

	temp1 := make([]byte, 2)
	binary.LittleEndian.PutUint16(temp1, transaction.Blocknum)
	bf.Write(temp1[0:2])

	temp2 := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp2, transaction.Blockprefix)
	bf.Write(temp2[0:4])

	temp3 := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp3, transaction.Expiration)
	bf.Write(temp3[0:4])

	temp4 := make([]byte, 1)
	temp4[0] = transaction.opsize
	bf.Write(temp4[0:1])

	temp5 := make([]byte, 1)
	temp5[0] = transaction.optype
	bf.Write(temp5[0:1])

	temp6 := make([]byte, 8)
	binary.LittleEndian.PutUint64(temp6, transaction.free)
	bf.Write(temp6[0:8])

	temp7 := make([]byte, 1)
	temp7[0] = transaction.freeasserid
	bf.Write(temp7[0:1])

	temp8 := transaction.from
	bf.Write(temp8[:])

	temp9 := transaction.to
	bf.Write(temp9[:])

	temp10 := make([]byte, 8)
	binary.LittleEndian.PutUint64(temp10, transaction.amount)
	bf.Write(temp10[0:8])

	temp11 := make([]byte, 1)
	temp11[0] = transaction.amountasserid
	bf.Write(temp11[0:1])

	//memo
	temp12 := make([]byte, 1)
	temp12[0] = 0
	bf.Write(temp12[0:1])
	//extensions
	temp13 := make([]byte, 1)
	temp13[0] = 0
	bf.Write(temp13[0:1])
	//extensions
	temp14 := make([]byte, 1)
	temp14[0] = 0
	bf.Write(temp14[0:1])

	return bf.Bytes()
}

//组织转账发送报文
func BtsSignTransfer(trade_id string, func_id string, block_id_time string, head_block_id string, chain_id string, private_key string, from string, to string, amount string, fee string, asser_id string) (string, error) {
	//产生私钥结构
	privKeyWif, _ := eosecc.NewPrivateKey(private_key)
	fmt.Println(privKeyWif.PublicKey().String())
	//构造交易结构
	var transaction Transfer
	transaction.Blocknum = uint16(stringToInt(head_block_id[0:8]))
	transaction.Blockprefix = uint32(stringToInt(head_block_id[8:16]) & -1)
	from_64, err := strconv.ParseUint(strings.Split(from, ".")[2], 10, 64)
	to_64, err := strconv.ParseUint(strings.Split(to, ".")[2], 10, 64)
	transaction.from = intToByte(from_64)
	transaction.to = intToByte(to_64)
	if strings.Split(asser_id, ".")[2][0] == '0' {
		transaction.amountasserid = 0
		transaction.freeasserid = 0
	} else {
		transaction.amountasserid = byte(strings.Split(asser_id, ".")[2][0])
		transaction.freeasserid = byte(strings.Split(asser_id, ".")[2][0])
	}
	transaction.amount, _ = strconv.ParseUint(amount, 10, 64)
	transaction.free, _ = strconv.ParseUint(fee, 10, 64)
	//转换为time_t类型
	block_id_time_tm, _ := time.Parse("2006-01-02 15:04:05", block_id_time[0:10]+" "+block_id_time[12:19])
	//默认设置有效期为2分钟之后
	transaction.Expiration = uint32(block_id_time_tm.UTC().Unix()) + 120
	transaction.optype = 0
	transaction.opsize = 1
	//序列化交易数据
	txdata := transferSerial(chain_id, transaction)
	//转换为sha256进行签名
	h := sha256.New()
	h.Write(txdata)
	fmt.Println("交易信息hash值:", hex.EncodeToString(h.Sum(nil)))
	sig, err := privKeyWif.Sign(h.Sum(nil))
	if err != nil {
		return "", err
	}
	//组织发送jaso数据
	tm := time.Unix(int64(transaction.Expiration), 0)
	//组织json数据
	date := tm.UTC().Format("2006-01-02T15:04:05")
	signature := hex.EncodeToString(sig.Content)
	block_num := strconv.Itoa(int(transaction.Blocknum))
	block_prefix := strconv.Itoa(int(transaction.Blockprefix))
	sendbuf := "{\"id\":" + trade_id + ",\"method\":\"call\",\"params\":[" + func_id + ",\"broadcast_transaction\",[{\"signatures\":[\"" + signature + "\"],\"expiration\":\"" + date + "\",\"ref_block_num\":" + block_num + ",\"ref_block_prefix\":" + block_prefix + ",\"operations\":[[0,{\"amount\":{\"amount\":" + amount + ",\"asset_id\":\"" + asser_id + "\"},\"extensions\":[],\"fee\":{\"amount\":" + fee + ",\"asset_id\":\"" + asser_id + "\"},\"from\":\"" + from + "\",\"to\":\"" + to + "\"}]],\"extensions\":[]}]]}"
	fmt.Println(sendbuf)
	return sendbuf, nil
}
