package nuls

import (
	"testing"
	"fmt"
	"encoding/json"
	"encoding/binary"
)

func TestGetBalance(t *testing.T) {
	amount, err := GetBalance("6HgUzobcB8pa5coQjYhWMexRjDJdtM5e")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("可用：%s \n冻结：%s \n", amount[0], amount[1])
	}
}

func TestGetUtxoUnspent(t *testing.T) {
	tx, err := GetUtxoUnspent("6HgUzobcB8pa5coQjYhWMexRjDJdtM5e")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json, _ := json.Marshal(tx)
		fmt.Printf("未发费的交易:\n %s\n", json)
	}
}
func TestTransferFee(t *testing.T) {
	fee,err:=TransferFee("6HgUzobcB8pa5coQjYhWMexRjDJdtM5e","HgjScviL1LWubo7U79v2bb4sE3mQ9UFs","1065","", "")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json, _ := json.Marshal(fee)
		fmt.Printf("未发费的交易:\n %s\n", json)
	}
}


func TestType(t *testing.T) {
	a:=make([]byte, 8)
	b:=a[0:6]
	fmt.Printf("%x",b)

	demo := map[string]bool{
		"a": false,
	}

	//错误，a存在，但是返回false
	fmt.Println(demo["av"])

	//正确判断方法
	v, ok := demo["av"]
	fmt.Println(v,ok)

	tt:=make([]byte,2)
	binary.LittleEndian.PutUint16(tt,uint16(8964))
	fmt.Println(tt)
}
