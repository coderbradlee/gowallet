package main

import (
	//"runtime"

	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/ethdb"

	"github.com/ethereum/go-ethereum/common"
	//"os"
	"testing"

	"github.com/ethereum/go-ethereum/trie"
)

func readDir(dir string) (uint64, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return 0, nil
}
func Test(t *testing.T) {
	x, err := readDir(".")
	fmt.Println(x, ":", err)

	//s := []int{0, 1, 2, 3, 4, 5, 6}
	//s = s[:4]
	//i := 5
	//_ = s[:i] // line 22: bounds check
	//_ = s[i:] // line 23: bounds check, not smart enough?

	//data, _ := hex.DecodeString("7261756c6c656e6368616978")
	//fmt.Println(len(data))
	//fmt.Println(data)
	//
	//xx := hex.EncodeToString([]byte("DRoute.Capital"))
	//fmt.Println(len(xx))
	//fmt.Println(xx)
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
