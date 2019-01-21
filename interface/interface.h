# go-HDWallet

//func CreateWalletByteRandAndPwd(random, salt []byte) (masterKey, mnemonic string, err error)
//Return masterKey(encrypt masterkey) and mnemonic(remember the user to backup)
func CreateWalletByteRandAndPwd(random []byte, password string) (masterKey, mnemonic string, err error)

//CreateWallet or ImportWallet by  mnemonic
//Return masterKey
func CreateWalletByMnnicAndPwd(mnemonic string, password []byte) (masterKey string, err error)

//Check Password is right //TODO
func CheckPwdIsCorrect(masterKey, password string) (right bool)

// Generate BIP44 account extended private key and extended public key by the path.
//Input masterKey string, coinType string(such as:BTC,ETC,ETH), account =0 , change=0, index =0  
//return address and the privateKey
func GenerateBIP44AccountWallet(masterKey string, coinType string, account, change, index uint32) (address, privateKey string, err error)

//Send ETH RawTransaction
//Input privateKey, nonce uint64, toAddr string, amount, gasLimit, gasPrice *big.Int; data is nil
//Return the singed data 
func SendETHRawTxByPrivateKey(privateKey string, nonce uint64, toAddr string, amount, gasLimit, gasPrice *big.Int, data []byte) (signedParam string, err error)
