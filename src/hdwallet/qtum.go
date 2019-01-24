package hdwallet

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (hd *Hdwallet) qtumAddress(child *hdkeychain.ExtendedKey) (address string, err error) {
	address_str, err := child.Address(&qtumAddressNetParams)
	if err != nil {
		return
	}
	address = address_str.String()
	return
}

///Send QTUM RawTransaction
func SendQTUMTokenRawTxByPrivateKey(privateKey string, toAddress string, balance float64, amount string, txFee string, gasLimit int64, gasPrice int64, data string) (signedParam string, err error) {
	amount = strings.Replace(amount, " ", "", -1)
	txFee = strings.Replace(txFee, " ", "", -1)
	fmt.Println("==================接收数据打印=========")
	fmt.Println("privateKey:", privateKey)
	fmt.Println("toAddress:", toAddress)
	fmt.Println("balance:", balance)
	fmt.Println("amount:", amount)
	fmt.Println("txFee:", txFee)
	fmt.Println("gasLimit:", gasLimit)
	fmt.Println("gasPrice:", gasPrice)
	fmt.Println("data:", data)

	var signRawTx string
	if (len(privateKey) == 0) || (len(toAddress) == 0) {
		return signRawTx, errors.New("some params is empty!!!")
	}
	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return signRawTx, err
	}
	//private_key, _ := child.ECPrivKey()
	address_str, err := child.Address(&qtumAddressNetParams)
	if err != nil {
		return signRawTx, err
	}
	//fromAddress string
	fromAddress := address_str.String()
	fmt.Println("The QTUM send address is ", fromAddress)

	changeAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("changeAmount err", err.Error())
	}
	changeTxfee, err := strconv.ParseFloat(txFee, 64)
	if err != nil {
		fmt.Println("changeTxfee err", err.Error())
	}
	fmt.Println("The string to double changeAmount  is ", changeAmount)
	fmt.Println("The string to double changeTxfee  is ", changeTxfee)

	createRawTx, err := createQtumTokenRawTransactionNew(fromAddress, toAddress, balance, changeAmount, changeTxfee, gasLimit, gasPrice, data)
	if err != nil {
		return signRawTx, err
	}
	//Sign
	signRawTx, err = signQtumTokenRawTransactionNew(createRawTx.Tx, createRawTx.PrevScripts, privateKey)
	return signRawTx, err
}

func GetQTUMBalanceByAddr(address string) (balance string, err error) {
	var _url string
	if qtumAddressNetParams.Name == "main" {
		_url = fmt.Sprintf("https://explorer.qtum.org/insight-api/addr/%s/balance", address)
	}
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Get(_url)
	if err != nil {
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("The get Balance info is ", string(bs))
	if resp.StatusCode != 200 {
		return "0.00", errors.New(string(bs))
	}
	balanceValue, err := strconv.ParseFloat(string(bs), 10)
	qtum := balanceValue / (100 * 1000 * 1000)
	retBalance := fmt.Sprintf("%.8f", qtum)
	if err != nil {
		return "", err
	} else {
		return retBalance, nil
	}
}

//Send QTUM RawTransaction
func SendQTUMRawTxByPrivateKey(privateKey string, toAddress string, amount string, txFee string) (signedParam string, err error) {
	fmt.Println("==================(SendQTUMRawTxByPrivateKey)接收数据打印=========")
	fmt.Println("privateKey:", privateKey)
	fmt.Println("toAddress:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("txFee:", txFee)

	var signRawTx string
	if (len(privateKey) == 0) || (len(toAddress) == 0) {
		return signRawTx, errors.New("some params is empty!!!")
	}
	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return signRawTx, err
	}
	//private_key, _ := child.ECPrivKey()
	address_str, err := child.Address(&qtumAddressNetParams)
	if err != nil {
		return signRawTx, err
	}
	//fromAddress string
	fromAddress := address_str.String()
	fmt.Println("The QTUM send address is ", fromAddress)

	changeAmount, err := strconv.ParseFloat(amount, 64)
	changeTxfee, err := strconv.ParseFloat(txFee, 64)

	createRawTx, err := createQtumRawTransactionNew(fromAddress, toAddress, changeAmount, changeTxfee)
	if err != nil {
		return signRawTx, err
	}
	//Sign
	signRawTx, err = signQtumRawTransactionNew(createRawTx.Tx, createRawTx.PrevScripts, privateKey)
	fmt.Println("The QTUM SIGN is ", signRawTx)
	return signRawTx, err
}

//Send QTUM RawTransaction size
func SendQTUMRawTxSizeByPrivateKey(privateKey string, toAddress string, amount string, txFee string) (size string, err error) {
	fmt.Println("==================接收数据打印(SendQTUMRawTxSizeByPrivateKey)=========")
	fmt.Println("privateKey:", privateKey)
	fmt.Println("toAddress:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("txFee:", txFee)

	var signRawTxSize string = "0"
	if (len(privateKey) == 0) || (len(toAddress) == 0) {
		return signRawTxSize, errors.New("some params is empty!!!")
	}
	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return signRawTxSize, err
	}
	//private_key, _ := child.ECPrivKey()
	address_str, err := child.Address(&qtumAddressNetParams)
	if err != nil {
		return signRawTxSize, err
	}
	//fromAddress string
	fromAddress := address_str.String()
	fmt.Println("The QTUMSendQTUMRawTxSizeByPrivateKey send address is ", fromAddress)

	changeAmount, err := strconv.ParseFloat(amount, 64)
	changeTxfee, err := strconv.ParseFloat(txFee, 64)

	createRawTx, err := createQtumRawTransactionNew(fromAddress, toAddress, changeAmount, changeTxfee)
	if err != nil {
		return signRawTxSize, err
	}
	//额外加10,size可大不可小,小交易会报错
	signRawTxSize = strconv.Itoa((148*len(createRawTx.Tx.TxIn) + 34*len(createRawTx.Tx.TxOut) + 10) - 1 + 10)
	//Sign(
	return signRawTxSize, err
}

//The QTUM Create Raw Transaction
func createQtumRawTransactionNew(fromAddress string, toAddress string, amount float64, minTxFee float64) (returnauthTx AuthoredTx, err error) {
	authTx := AuthoredTx{}
	if amount < 0.002184 {
		return authTx, &btcjson.RPCError{
			Code:    btcjson.ErrRPCType,
			Message: "The payment must be greater than 0.002184",
		}
	}
	//Some Variant
	params := qtumAddressNetParams
	//金额校验
	if amount <= 0 {
		return authTx, &btcjson.RPCError{
			Code:    btcjson.ErrRPCType,
			Message: "Invalid amount",
		}
	}
	//Analyse the unspent from the Wallet
	//Analyse the unspent from the blockchain
	qtumUnspent, err := loadQtumUnspentByAddress(fromAddress)
	if err != nil {
		return authTx, err
	}
	//init the unspent
	unspentlen := len(qtumUnspent)
	if unspentlen < 0 {
		return authTx, errors.New("The list unSpent is null!!")
	}
	var isStake int = 0
	for index := 0; index < unspentlen; index++ {
		if qtumUnspent[index].IsStake == true && qtumUnspent[index].Confirmations < 501 {
			isStake++
		}
	}
	var listIndex int = 0
	var listunspents = make([]btcjson.ListUnspentResult, unspentlen-isStake)

	for index := 0; index < unspentlen; index++ {
		if qtumUnspent[index].IsStake == true && qtumUnspent[index].Confirmations < 501 {
			continue
		}
		unspentTx := qtumUnspent[index]
		var txUnspent btcjson.ListUnspentResult
		txUnspent.TxID = unspentTx.Txid
		txUnspent.Vout = (uint32)(unspentTx.Vout)
		txUnspent.Amount = unspentTx.Amount
		txUnspent.ScriptPubKey = unspentTx.ScriptPubKey
		txUnspent.Confirmations = unspentTx.Confirmations
		//txUnspent.RedeemScript = unspentTx.Script_asm
		txUnspent.Address = unspentTx.Address
		txUnspent.Spendable = true
		listunspents[listIndex] = txUnspent
		listIndex++
	}
	//Get the send info
	var reallistunspentslen int = 0
	for uu := 0; uu < len(listunspents); uu++ {
		unspent_record := listunspents[uu]
		if (unspent_record.Amount > 0) && (unspent_record.Confirmations > 0) {
			reallistunspentslen++
		}
	}
	var array_transaction_in = make([]btcjson.TransactionInput, reallistunspentslen)
	var array_prevPkScripts = make([]string, reallistunspentslen)

	var sum_amount float64 = 0.0
	var inputsNum = 0
	//统计所有未花费的交易金额之和，减少因为量子链的最小交易限制造成的损失
	for uu := 0; uu < len(listunspents); uu++ {
		unspent_record := listunspents[uu]
		if (unspent_record.Amount > 0) && (unspent_record.Confirmations > 0) {
			sum_amount += (unspent_record.Amount) //* 100000000
			var txInput btcjson.TransactionInput
			txInput.Txid = unspent_record.TxID
			txInput.Vout = unspent_record.Vout
			array_transaction_in[inputsNum] = txInput
			addr, _ := btcutil.DecodeAddress(unspent_record.Address, &params)
			scriptAdd, _ := txscript.PayToAddrScript(addr)
			array_prevPkScripts[inputsNum] = string(scriptAdd)
			inputsNum++
		}
	}
	fmt.Printf("The sum amount is %.8f", sum_amount)

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
	mtx := wire.NewMsgTx(2) //wire.TxVersion
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

		if amount < 0.002184 {
			continue
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
	fmt.Println("The New Create QTUM raw Transaction is", mtxHex)
	return authTx, nil
}

//The QTUM Create Raw Transaction
func createQtumTokenRawTransactionNew(fromAddress string, toAddress string, balance float64, amount float64, minTxFee float64, gasLimit int64, gasPrice int64, data string) (returnauthTx AuthoredTx, err error) {
	authTx := AuthoredTx{}
	if amount < 0.002184 {
		return authTx, &btcjson.RPCError{
			Code:    btcjson.ErrRPCType,
			Message: "The payment must be greater than 0.001873",
		}
	}
	//Some Variant
	params := qtumAddressNetParams
	//金额校验
	if amount <= 0 {
		return authTx, &btcjson.RPCError{
			Code:    btcjson.ErrRPCType,
			Message: "Invalid amount",
		}
	}
	//Analyse the unspent from the Wallet
	//Analyse the unspent from the blockchain
	qtumUnspent, err := loadQtumUnspentByAddress(fromAddress)
	if err != nil {
		return authTx, err
	}
	//init the unspent
	unspentlen := len(qtumUnspent)
	if unspentlen < 0 {
		return authTx, errors.New("The list unSpent is null!!")
	}

	var isStake int = 0
	for index := 0; index < unspentlen; index++ {
		if qtumUnspent[index].IsStake == true && qtumUnspent[index].Confirmations < 501 {
			isStake++
		}
	}
	var listIndex int = 0
	var listunspents = make([]btcjson.ListUnspentResult, unspentlen-isStake)

	for index := 0; index < unspentlen; index++ {
		if qtumUnspent[index].IsStake == true && qtumUnspent[index].Confirmations < 501 {
			continue
		}
		unspentTx := qtumUnspent[index]
		var txUnspent btcjson.ListUnspentResult
		txUnspent.TxID = unspentTx.Txid
		txUnspent.Vout = (uint32)(unspentTx.Vout)
		txUnspent.Amount = unspentTx.Amount
		txUnspent.ScriptPubKey = unspentTx.ScriptPubKey
		txUnspent.Confirmations = unspentTx.Confirmations
		//txUnspent.RedeemScript = unspentTx.Script_asm
		txUnspent.Address = unspentTx.Address
		txUnspent.Spendable = true
		listunspents[listIndex] = txUnspent
		listIndex++
	}
	//Get the send info
	var array_transaction_in = make([]btcjson.TransactionInput, len(listunspents))
	var array_prevPkScripts = make([]string, len(listunspents))

	var sum_amount float64 = 0.0
	var inputsNum = 0
	//统计所有未花费的交易金额之和，减少因为量子链的最小交易限制造成的损失
	for uu := 0; uu < len(listunspents); uu++ {
		unspent_record := listunspents[uu]
		if (unspent_record.Amount > 0) && (unspent_record.Confirmations > 0) {
			var txInput btcjson.TransactionInput
			sum_amount += (unspent_record.Amount) //* 100000000
			txInput.Txid = unspent_record.TxID
			txInput.Vout = unspent_record.Vout
			array_transaction_in[uu] = txInput
			addr, _ := btcutil.DecodeAddress(unspent_record.Address, &params)
			scriptAdd, _ := txscript.PayToAddrScript(addr)
			array_prevPkScripts[uu] = string(scriptAdd)
			inputsNum++
		}
	}
	fmt.Printf("The sum amount is %.8f", sum_amount)

	var inputs = make([]btcjson.TransactionInput, inputsNum)
	var prevPkScripts = make([]string, inputsNum)
	copy(inputs, array_transaction_in[:inputsNum])
	copy(prevPkScripts, array_prevPkScripts[:inputsNum])

	//确保新交易的输入金额满足最小交易条件
	if sum_amount < minTxFee {
		fmt.Println("Invalid unspent amount")
		return authTx, errors.New("Invalid unspent amount")
	}
	fmt.Println("Transaction_in:", inputs)

	var tempAmount float64
	tempAmount = sum_amount - minTxFee
	changeAmountstr := fmt.Sprintf("%.8f", tempAmount)
	fmt.Println("The change Amount is ", changeAmountstr)
	//changeAmountstr: = strconv.FormatFloat(changeAmount,'f',8,64)
	changeAmount, err := strconv.ParseFloat(changeAmountstr, 64)
	if err != nil {
		return authTx, errors.New("the Float can not save 8 point number")
	}
	//生成测试新交易的输出数据块，此处示例是给指定目标测试钱包地址转账一小笔测试比特币
	//注意：输入总金额与给目标转账加找零金额间的差额即MIN_TRANSACTION_FEE，就是支付给比特币矿工的交易成本费用
	addAmoutsMap := map[string]float64{
		fromAddress: changeAmount, //(sum_amount - amount - minTxFee),找零地址和金额，默认用发送者地址
	}
	// Add all transaction inputs to a new transaction after performing some validity checks.
	var lockTime int64
	lockTime = 0
	mtx := wire.NewMsgTx(2) //wire.TxVersion
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
	{
		//设置gas
		gas_buf := bytes.NewBuffer([]byte{})
		binary.Write(gas_buf, binary.BigEndian, gasLimit)
		gas := flashBackInt(gas_buf.Bytes())
		//设置gaslimit
		price_buf := bytes.NewBuffer([]byte{})
		binary.Write(price_buf, binary.BigEndian, gasPrice)
		price := flashBackInt(price_buf.Bytes())
		//设置合约数据
		hexdata := make([]byte, len(data)/2)
		hex.Decode(hexdata, []byte(data))
		//设置合约地址
		hextoaddress := make([]byte, len(toAddress)/2)
		hex.Decode(hextoaddress, []byte(toAddress))
		//创建pkscrip
		pkScript, err := txscript.NewScriptBuilder().AddOp(0x54).AddData(gas).AddData(price).AddData(hexdata).AddData(hextoaddress).AddOp(0xc2).Script()

		// Convert the amount to satoshi.
		satoshi, err := btcutil.NewAmount(0)
		if err != nil {
			context := "Failed to convert amount"
			return authTx, errors.New(context) //internalRPCError(err.Error(), context)
		}
		txOut := wire.NewTxOut(int64(satoshi), pkScript)
		mtx.AddTxOut(txOut)
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
		//量子链有限制小于这个金额不能交易
		if amount < 0.002184 {
			continue
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
	fmt.Println("The New Create QTUM raw Transaction is", mtxHex)
	return authTx, nil
}

//GetUnspent list info
func loadQtumUnspentByAddress(address string) (qtumUnspent []QutmDataInfo, err error) {
	//https://chain.so/api/v2/get_tx_unspent/BTC/n4GKiozs2zqokewPEcPoy7wXfcYap8q1Ai
	//https://chain.so/api/v2/get_tx_unspent/BTCTEST/n4GKiozs2zqokewPEcPoy7wXfcYap8q1Ai
	var rest []QutmDataInfo
	var _url string
	if qtumAddressNetParams.Name == "main" {
		_url = fmt.Sprintf("https://explorer.qtum.org/insight-api/addr/%s/utxo", address)
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
	if resp.StatusCode == 200 {
		fmt.Println("The Struct is #v%", rest)
	} else {
		rest[0] = QutmDataInfo{}
		err = errors.New("The Get Block Chain is wrong!!!!")
	}
	//fmt.Println("The Struct is #v%", rest)
	return rest, err
}

//Sign Raw Transaction with Privatekey
func signQtumRawTransactionNew(tx *wire.MsgTx, prevPkScripts []string, privateKey string) (r string, err error) {
	child, err := hdkeychain.NewKeyFromString(privateKey)
	inputs := tx.TxIn
	chainParams := qtumAddressNetParams
	private_key, _ := child.ECPrivKey()
	//privateKeyBytes := private_key.Serialize()
	address, err := child.Address(&qtumAddressNetParams)
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

//Sign Raw Transaction with Privatekey
func signQtumTokenRawTransactionNew(tx *wire.MsgTx, prevPkScripts []string, privateKey string) (r string, err error) {
	child, err := hdkeychain.NewKeyFromString(privateKey)
	inputs := tx.TxIn
	chainParams := qtumAddressNetParams
	private_key, _ := child.ECPrivKey()
	//privateKeyBytes := private_key.Serialize()
	address, err := child.Address(&qtumAddressNetParams)
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

//Sign Raw Transaction with Privatekey
func signQtumRawTransactionSizeNew(tx *wire.MsgTx, prevPkScripts []string, privateKey string) (r int, err error) {
	var reSignRawTx int = 0
	child, err := hdkeychain.NewKeyFromString(privateKey)
	inputs := tx.TxIn
	chainParams := qtumAddressNetParams
	private_key, _ := child.ECPrivKey()
	//privateKeyBytes := private_key.Serialize()
	address, err := child.Address(&qtumAddressNetParams)
	//pkScript, err := txscript.PayToAddrScript(address)
	//fmt.Println("The create pkScript is", pkScript)
	if len(inputs) != len(prevPkScripts) {
		return reSignRawTx, errors.New("tx.TxIn and prevPkScripts slices must " +
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
			return reSignRawTx, err
		}
		inputs[i].SignatureScript = sigScript
	}
	//fmt.Println(sizestruct.SizeStruct(tx))
	reSignRawTx, err = messageToHexLen(tx)
	return reSignRawTx, err
}
