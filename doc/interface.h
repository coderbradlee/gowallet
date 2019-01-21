# go-HDWallet

# go-HDWallet

//func CreateWalletByteRandAndPwd(random, salt []byte) (masterKey, mnemonic string, err error)
//Return masterKey(encrypt masterkey) and mnemonic(remember the user to backup)
/*
  启动创建钱包，需要用户输入密码：password;随机数(random) 可以自己输入，也可以不输入，后续根据XRandom生成；
  
  输出参数：
      ＠主私钥字符串键值（masterKey），助记词字符串（mnemonic）；
	  通过gomobile 生成之后aar文件；主私钥与助记词字符串返回字符串使用ＪＳＯＮ格式（{"MasterKey":"xprv9s",,"Mnemonic":"wordslist"}）；
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
*/
func CreateWalletByteRandAndPwd(random string, password string) (masterKey, mnemonic string, err error)

//CreateWallet or ImportWallet by  mnemonic
//Return masterKey

/*
  进行导入钱包，需要用户输入助记词：mnemonic，密码：password；
  
  输出参数：
      主私钥字符串键值masterKey；
	  
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func CreateWalletByMnnicAndPwd(mnemonic string, password string ) (masterKey string, err error)

//Check Password is right 
/*
  校验钱包密码是否是正确，需要用户输入加密后的主私钥：masterKey，密码：password； 都为字符串
  
  输出参数：
      bool
	  
	  通过gomobile 生成之后aar文件；直接返回bool 值
*/
func CheckPwdIsCorrect(masterKey, password string) (right bool)

// Generate BIP44 account extended private key and extended public key by the path.
//Input masterKey string, coinType string(such as:BTC,ETC,ETH), account =0 , change=0, index =0  
//return address and the privateKey
/*
   在用户进行了输入密码，生成主私钥后，可以根据币种进行创建真正具体的钱包 
  输入参数：   
	  @主私钥：masterKey（字符串），
	  @币种类型： coinType（字符串：BTC,ETC,ETH）；
	  @ account, change, index 为 整数类型，都设置为：0
  
  输出参数：
        ＠具体钱包的地址字符串值（address），具体钱包的私钥字符串值（privateKey）；
	  通过gomobile 生成之后aar文件；地址与私钥字符串返回字符串使用"1JS7r"
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index int) (address string, err error)

//Send ETH RawTransaction
//Input privateKey, nonce uint64, toAddr string, amount, gasLimit, gasPrice *big.Int; data is nil
//Return the singed data 

/*
   在用户进行选择具体钱包之后，进行发送ETH,ETC 交易 
  输入参数：   
	@ 钱包对应的私钥：privateKey（字符串），
	@ ETH 随机数：nonce（字符串---为十进制整数类型如：123455，不能为：abcd ； 
	@ 转账到目的地地址：toAddr（字符串）；
	@ 转账金额：amount（字符串---为十进制整数类型（big.Int）如：123455，不能为：abcd）;
	@ GAS 上限：gasLimit（字符串---为十进制整数类型（big.Int）如：123455，不能为：abcd）;
	@ gas价格： gasPrice字符串---为十进制整数类型（big.Int）如：123455，不能为：abcd）;
	@ 额外数据：Data 不填（在发送代币时使用）；
  
  输出参数：
        ＠签名后的字符串值（signedParam）
	  通过gomobile 生成之后aar文件；直接以字符串返回
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func SendETHRawTxByPrivateKey(privateKey string, nonce, toAddr, amount, gasLimit, gasPrice string, data []byte) (signedParam string, err error)


/*
      在用户进行选择具体钱包之后，进行发送交易之前进行密码校验并获取私钥 
  输入参数：   
	  @主私钥：masterKey（字符串），
	  @密码：password（字符串），
	  @币种类型： coinType（字符串：BTC,ETC,ETH）；
	  @ account, change, index 为 整数类型，都设置为：0
  
  输出参数：
        ＠具体钱包的私钥字符串值（privateKey）；
	  通过gomobile 生成之后aar文件；地址与私钥字符串返回字符串如：xprv........
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func CheckAuthAndGetPrivateKey(masterKey string, password string, coinType string, account, change, index int) ( privateKey string, err error)

//BTC 相关接口说明 
/*
      输入BTC 的地址，然后进行余额查询 
  输入参数：   
	  @BTC 的地址：address（字符串），
  
  输出参数：
        ＠具体BTC 的地址的余额值balance （字符串）；

	    @ 错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func GetBTCBalanceByAddr(address string) (balance string, err error)


//Send BTC RawTransaction
/*
   在用户进行选择BTC钱包之后，进行发送BTC 交易 
  输入参数：   
	@ 钱包对应的私钥：privateKey（字符串），
	@ 转账到目的地地址：toAddress（字符串---为十六进制类型如：0x122222ab)； 
	@ 转账金额：amount(Double ,对应IOS以及Android 的Double 类型);
	@ 转账费用：txFee(Double ,对应IOS以及Android 的Double 类型,最低为：0.00001);
  
  输出参数：
        ＠签名后的字符串值（signedParam）
	     通过gomobile 生成之后aar文件；直接以字符串返回
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func SendBTCRawTxByPrivateKey(privateKey string, toAddress string, amount float64, txFee float64) (signedParam string, err error) 

//LTC 相关接口说明 
/*
      输入LTC 的地址，然后进行余额查询 
  输入参数：   
	  @LTC 的地址：address（字符串），
  
  输出参数：
        ＠具体LTC 的地址的余额值balance （字符串）；

	    @ 错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func GetLTCBalanceByAddr(address string) (balance string, err error)


//Send LTC RawTransaction
/*
   在用户进行选择LTC钱包之后，进行发送LTC 交易 
  输入参数：   
	@ 钱包对应的私钥：privateKey（字符串），
	@ 转账到目的地地址：toAddress（字符串---为十六进制类型如：0x122222ab)； 
	@ 转账金额：amount(Double ,对应IOS以及Android 的Double 类型);
	@ 转账费用：txFee(Double ,对应IOS以及Android 的Double 类型,最低为：0.00001);
  
  输出参数：
        ＠签名后的字符串值（signedParam）
	     通过gomobile 生成之后aar文件；直接以字符串返回
	   错误代码；（android需要使用ｔｒｙ－ｃａｔｃｈ；）；IOS 可以直接使用error
	  通过gomobile 生成之后aar文件；直接返回字符串，需要进行try-catch
*/
func SendLTCRawTxByPrivateKey(privateKey string, toAddress string, amount float64, txFee float64) (signedParam string, err error) 