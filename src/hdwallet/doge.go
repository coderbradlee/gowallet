package hdwallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"hdwallet/nuls"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetDOGEBalanceByAddr(address string) (balance string, err error) {
	// /api/v2/get_address_balance/{NETWORK}/{ADDRESS}[/{MINIMUM CONFIRMATIONS}]
	//https://chain.so/api/v2/get_address_balance/DOGE/%s
	var _url string
	if dogeAddressNetParams.Name == "mainnet" {
		_url = fmt.Sprintf("https://chain.so/api/v2/get_address_balance/DOGE/%s/", address)
	}
	client := &http.Client{
		Timeout: requestTimeout,
	}
	var rest ChainBalanceInfo
	resp, err := client.Get(_url)
	if err != nil {
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("The get Balance info is ", string(bs))
	err = json.Unmarshal(bs, &rest)
	if err != nil {
		//fmt.Println("There are some errors:", err)
		return getLtcExtendBalance(address)
	}
	if rest.Status != "success" {
		//err = errors.New("The Get Balance is wrong!!!!")
		return getLtcExtendBalance(address)
	}
	//fmt.Println("The balance is #v%", rest.Data.Confirmed_balance)
	//balance, _ := strconv.ParseFloat(rest.Data.Confirmed_balance, 64)
	return rest.Data.Confirmed_balance, err
}

//======Below is for DOGE Info============================================================//

//Send DOGE RawTransaction
func SendDOGERawTxByPrivateKey(privateKey string, toAddress string, amount float64, txFee float64) (signedParam string, err error) {
	var signRawTx string
	if (len(privateKey) == 0) || (len(toAddress) == 0) {
		return signRawTx, errors.New("some params is empty!!!")
	}
	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return signRawTx, err
	}
	//private_key, _ := child.ECPrivKey()
	address_str, err := child.Address(&dogeAddressNetParams)
	if err != nil {
		return signRawTx, err
	}
	//fromAddress string
	fromAddress := address_str.String()
	fmt.Println("The DOGE send address is ", fromAddress)

	createRawTx, err := createDogeRawTransactionNew(fromAddress, toAddress, amount, txFee)
	if err != nil {
		return signRawTx, err
	}
	//Sign
	signRawTx, err = signDogeRawTransactionNew(createRawTx.Tx, createRawTx.PrevScripts, privateKey)
	return signRawTx, err
}

//The DOGE Create Raw Transaction
func createDogeRawTransactionNew(fromAddress string, toAddress string, amount float64, minTxFee float64) (returnauthTx AuthoredTx, err error) {

	//Some Variant
	authTx := AuthoredTx{}
	params := dogeAddressNetParams
	//金额校验
	if amount <= 0 {
		return authTx, &btcjson.RPCError{
			Code:    btcjson.ErrRPCType,
			Message: "Invalid amount",
		}
	}
	//Analyse the unspent from the Wallet
	//Analyse the unspent from the blockchain
	chainUnspent, err := loadDogeUnspentByAddress(fromAddress)
	if err != nil {
		return authTx, err
	}
	//init the unspent
	unspentlen := len(chainUnspent.Data.Txs)
	if unspentlen < 0 {
		return authTx, errors.New("The list unSpent is null!!")
	}
	var listunspents = make([]btcjson.ListUnspentResult, unspentlen)
	for index := 0; index < unspentlen; index++ {
		unspentTx := chainUnspent.Data.Txs[index]
		var txUnspent btcjson.ListUnspentResult
		txUnspent.TxID = unspentTx.Txid
		txUnspent.Vout = (uint32)(unspentTx.Vout)
		txUnspent.Amount, _ = strconv.ParseFloat(unspentTx.Value, 10)
		txUnspent.ScriptPubKey = unspentTx.Script_hex
		txUnspent.Confirmations = (int64)(unspentTx.Confirmations)
		txUnspent.RedeemScript = unspentTx.Script_asm
		txUnspent.Address = chainUnspent.Data.Address
		txUnspent.Spendable = true
		listunspents[index] = txUnspent
	}
	//Get the send info
	var array_transaction_in = make([]btcjson.TransactionInput, len(listunspents))
	var array_prevPkScripts = make([]string, len(listunspents))

	var sum_amount float64 = 0.0
	var inputsNum = 0
	for uu := 0; uu < len(listunspents); uu++ {
		unspent_record := listunspents[uu]
		if (unspent_record.Amount > 0) && (unspent_record.Confirmations > 0) {
			sum_amount += (unspent_record.Amount) //* 100000000
			var txInput btcjson.TransactionInput
			txInput.Txid = unspent_record.TxID
			txInput.Vout = unspent_record.Vout
			array_transaction_in[uu] = txInput
			addr, _ := btcutil.DecodeAddress(unspent_record.Address, &params)
			scriptAdd, _ := txscript.PayToAddrScript(addr)
			array_prevPkScripts[uu] = string(scriptAdd)
			inputsNum++

			if sum_amount >= (amount + minTxFee) { //*100000000
				break
			}
		}
	}
	fmt.Println("The sum amount is ", sum_amount)

	var inputs = make([]btcjson.TransactionInput, inputsNum)
	var prevPkScripts = make([]string, inputsNum)
	copy(inputs, array_transaction_in[:inputsNum])
	copy(prevPkScripts, array_prevPkScripts[:inputsNum])

	//确保新交易的输入金额满足最小交易条件
	if sum_amount < (amount + minTxFee) {
		return authTx, errors.New("Invalid unspent amount")
	}
	fmt.Println("Transaction_in:", inputs)
	var tempAmount float64
	if toAddress == fromAddress {
		tempAmount = sum_amount - minTxFee
	} else {
		tempAmount = sum_amount - amount - minTxFee
	}
	changeAmountstr := fmt.Sprintf("%.8f", tempAmount)
	fmt.Println("The change Amount is ", changeAmountstr)
	fmt.Println("The toAddress is ", toAddress)
	//changeAmountstr: = strconv.FormatFloat(changeAmount,'f',8,64)
	changeAmount, err := strconv.ParseFloat(changeAmountstr, 64)
	if err != nil {
		return authTx, errors.New("the Float can not save 8 point number")
	}
	//生成测试新交易的输出数据块，此处示例是给指定目标测试钱包地址转账一小笔测试比特币
	//注意：输入总金额与给目标转账加找零金额间的差额即MIN_TRANSACTION_FEE，就是支付给比特币矿工的交易成本费用
	addAmoutsMap := map[string]float64{
		toAddress:   amount,       //目标转账地址和金额
		fromAddress: changeAmount, //(sum_amount - amount - minTxFee),找零地址和金额，默认用发送者地址
	}
	// Add all transaction inputs to a new transaction after performing some validity checks.
	var lockTime int64
	lockTime = 0
	mtx := wire.NewMsgTx(1) //wire.TxVersion
	for _, input := range inputs {
		txHash, err := chainhash.NewHashFromStr(input.Txid)
		if err != nil {
			return authTx, err //rpcDecodeHexError(input.Txid)
		}

		prevOut := wire.NewOutPoint(txHash, input.Vout)
		txIn := wire.NewTxIn(prevOut, []byte{}, nil)
		if lockTime != 0 {
			txIn.Sequence = wire.MaxTxInSequenceNum - 1
		}
		mtx.AddTxIn(txIn)
	}
	// Add all transaction outputs to the transaction after performing
	// some validity checks.
	for encodedAddr, amount := range addAmoutsMap {
		// Ensure amount is in the valid range for monetary amounts.
		if amount < 0 || amount > btcutil.MaxSatoshi {
			return authTx, &btcjson.RPCError{
				Code:    btcjson.ErrRPCType,
				Message: "Invalid amount",
			}
		}

		// Decode the provided address.
		//addr, err := btcutil.DecodeAddress(encodedAddr, nil)
		addr, err := btcutil.DecodeAddress(encodedAddr, &params)
		if err != nil {
			return authTx, &btcjson.RPCError{
				Code:    btcjson.ErrRPCInvalidAddressOrKey,
				Message: "Invalid address or key: " + err.Error(),
			}
		}

		// Ensure the address is one of the supported types and that
		// the network encoded with the address matches the network the
		// server is currently on.
		switch addr.(type) {
		case *btcutil.AddressPubKeyHash:
		case *btcutil.AddressScriptHash:
		default:
			return authTx, &btcjson.RPCError{
				Code:    btcjson.ErrRPCInvalidAddressOrKey,
				Message: "Invalid address or key",
			}
		}
		if !addr.IsForNet(&params) {
			return authTx, &btcjson.RPCError{
				Code: btcjson.ErrRPCInvalidAddressOrKey,
				Message: "Invalid address: " + encodedAddr +
					" is for the wrong network",
			}
		}
		// Create a new script which pays to the provided address.
		pkScript, err := txscript.PayToAddrScript(addr)
		if err != nil {
			context := "Failed to generate pay-to-address script"
			return authTx, errors.New(context) //internalRPCError(err.Error(), context)
		}
		// Convert the amount to satoshi.
		satoshi, err := btcutil.NewAmount(amount)
		if err != nil {
			context := "Failed to convert amount"
			return authTx, errors.New(context) //internalRPCError(err.Error(), context)
		}

		txOut := wire.NewTxOut(int64(satoshi), pkScript)
		mtx.AddTxOut(txOut)
	}

	// Set the Locktime, if given.
	if lockTime != 0 {
		mtx.LockTime = uint32(lockTime)
	}
	authTx = AuthoredTx{
		Tx:          mtx,
		PrevScripts: prevPkScripts,
	}

	// Return the serialized and hex-encoded transaction.
	mtxHex, err := messageToHex(mtx)
	if err != nil {
		return authTx, err
	}
	fmt.Println("The New Create DOGE raw Transaction is", mtxHex)
	return authTx, nil
}

//GetUnspent list info
func loadDogeUnspentByAddress(address string) (chainUnspent ChainUnspentInfo, err error) {
	//https://chain.so/api/v2/get_tx_unspent/BTC/n4GKiozs2zqokewPEcPoy7wXfcYap8q1Ai
	//https://chain.so/api/v2/get_tx_unspent/BTCTEST/n4GKiozs2zqokewPEcPoy7wXfcYap8q1Ai
	var rest ChainUnspentInfo
	var _url string
	if dogeAddressNetParams.Name == "mainnet" {
		_url = fmt.Sprintf("https://chain.so/api/v2/get_tx_unspent/DOGE/%s", address)
	}

	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Get(_url)
	if err != nil {
		return rest, err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rest, err
	}
	fmt.Println("The get Block chain info is ", string(bs))
	err = json.Unmarshal(bs, &rest)
	if err != nil {
		fmt.Println("There are some errors:", err)
		return rest, err
	}
	if rest.Status == "success" {
		fmt.Println("The Struct is #v%", rest)
	} else {
		rest = ChainUnspentInfo{}
		err = errors.New("The Get Block Chain is wrong!!!!")
	}
	//fmt.Println("The Struct is #v%", rest)
	return rest, err
}

//Sign Raw Transaction with Privatekey
func signDogeRawTransactionNew(tx *wire.MsgTx, prevPkScripts []string, privateKey string) (r string, err error) {
	child, err := hdkeychain.NewKeyFromString(privateKey)
	inputs := tx.TxIn
	chainParams := dogeAddressNetParams
	private_key, _ := child.ECPrivKey()
	//privateKeyBytes := private_key.Serialize()
	address, err := child.Address(&dogeAddressNetParams)
	//pkScript, err := txscript.PayToAddrScript(address)
	//fmt.Println("The create pkScript is", pkScript)
	if len(inputs) != len(prevPkScripts) {
		return "", errors.New("tx.TxIn and prevPkScripts slices must " +
			"have equal length")
	}
	for i := range inputs {
		pkScript := prevPkScripts[i]
		//sigScript := inputs[i].SignatureScript
		secrets1 := mkGetKey(map[string]addressToKey{
			address.EncodeAddress(): {private_key, true},
		})
		secrets2 := mkGetScript(nil)
		sigScript, err := txscript.SignTxOutput(&chainParams, tx, i,
			[]byte(pkScript), txscript.SigHashAll, secrets1, secrets2,
			nil)

		if err != nil {
			return "", err
		}
		inputs[i].SignatureScript = sigScript
	}
	reSignRawTx, err := messageToHex(tx)
	return reSignRawTx, err
}
