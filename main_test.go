package main

import (
	//"runtime"

	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	//"os"
	"testing"

	"github.com/ethereum/go-ethereum/trie"
)

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
