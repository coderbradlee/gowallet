package main

import (
	//"runtime"

	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/ethdb"

	"github.com/ethereum/go-ethereum/common"
	//"os"
	"testing"

	"github.com/ethereum/go-ethereum/trie"
)

func Test(t *testing.T) {
	data, _ := hex.DecodeString("608060405234801561001057600080fd5b50610370806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e67bed8811461005b5780638965356f146100d1578063d0e30db01461012a575b600080fd5b6040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff16935061013492505050565b60408051918252519081900360200190f35b3480156100dd57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf94369492936024939284019190819084018382808284375094975061018e9650505050505050565b6101326102a4565b005b6000806101408461018e565b60405190915073ffffffffffffffffffffffffffffffffffffffff84169082156108fc029083906000818181858888f19350505050158015610186573d6000803e3d6000fd5b509392505050565b6000600561029060024485426040516020018084815260200183805190602001908083835b602083106101d25780518252601f1990920191602091820191016101b3565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382019081905283519395509350839290850191508083835b602083106102335780518252601f199092019160209182019101610214565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af1158015610274573d6000803e3d6000fd5b5050506040513d602081101561028957600080fd5b50516102dc565b81151561029957fe5b066001019050919050565b60408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a2565b600080805b602081101561033d5780600101602060ff160360080260020a848260208110151561030857fe5b7f010000000000000000000000000000000000000000000000000000000000000091901a8102040291909101906001016102e1565b50929150505600a165627a7a72305820d432ab12368e59501d1e616eeecda22f3f3b461a5c5a6b919a1d2331a2e105bc0029")
	fmt.Println(len(data))
	//err := error(ErrorTyped{errors.New("an error occurred")})
	//err = errors.Wrap(err, "wrapped")
	//
	//fmt.Println("wrapped error: ", err)
	//
	//// 处理错误类型
	//switch errors.Cause(err).(type) {
	//case ErrorTyped:
	//	fmt.Println("a typed error occurred: ", err)
	//default:
	//	fmt.Println("an unknown error occurred")
	//}
}
func TestAll(t *testing.T) {
	dir, err := ioutil.TempDir("", "trie-bench")
	if err != nil {
		panic(fmt.Sprintf("can't create temporary directory: %v", err))
	}
	diskdb, err := ethdb.NewLDBDatabase(dir, 0, 0)
	if err != nil {
		panic(fmt.Sprintf("can't create temporary database: %v", err))
	}
	db := trie.NewDatabase(diskdb)
	tr, err := trie.New(common.Hash{}, db)
	updateString(tr, "120000", "qwerqwerqwerqwerqwerqwerqwerqwer")
	updateString(tr, "123456", "asdfasdfasdfasdfasdfasdfasdfasdf")
	root, _ := tr.Commit(nil)
	tr, _ = trie.New(root, db)
	v, err := tr.TryGet([]byte("120000"))
	fmt.Println(string(v))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func updateString(tr *trie.Trie, k, v string) {
	tr.Update([]byte(k), []byte(v))
}
