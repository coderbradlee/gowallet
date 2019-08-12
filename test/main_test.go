package test

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/boltdb/bolt"
)

//func readDir(dir, model string) (uint64, error) {
//	files, err := ioutil.ReadDir(dir)
//	if err != nil {
//		return 0, err
//	}
//	maxN := uint64(0)
//	for _, file := range files {
//		name := file.Name()
//		lens := len(name)
//		if lens < 11 || !strings.Contains(name, model) {
//			continue
//		}
//		num := name[lens-11 : lens-3]
//		n, err := strconv.Atoi(num)
//		if err != nil {
//			continue
//		}
//		if uint64(n) > maxN {
//			maxN = uint64(n)
//		}
//	}
//	return maxN + 1, nil
//}
func TestXx(t *testing.T) {
	//x, err := readDir(".", "chain")
	//fmt.Println(x, ":", err)
	//name := "chain" + fmt.Sprintf("-%08d", x) + ".db"
	//fmt.Println(name)
	//d := time.Duration(3) * time.Second
	//tt := time.NewTicker(d)
	//defer tt.Stop()
	//for {
	//	select {
	//	case <-tt.C:
	//		fmt.Println("start")
	//		time.Sleep(time.Second * 10)
	//		fmt.Println("end")
	//	}
	//}
	//testbolt()
	testhex()
}
func testhex() {
	data := "000000000000000000000000391694e7e0b0cce554cb130d723a9d27458f9298" + "0000000000000000000000000000000000000000000000000000000000000001"
	hb, _ := hex.DecodeString(data)
	out := crypto.Keccak256(hb)
	fmt.Println(hex.EncodeToString(out))
}
func testbolt() error {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, 10)
		key := append([]byte("1111"), bytes...)
		err = b.Put(key, []byte("10"))

		binary.BigEndian.PutUint64(bytes, 100)
		key = append([]byte("1111"), bytes...)
		err = b.Put(key, []byte("100"))
		binary.BigEndian.PutUint64(bytes, 1000)
		key = append([]byte("1111"), bytes...)
		err = b.Put(key, []byte("1000"))
		binary.BigEndian.PutUint64(bytes, 10000)
		key = append([]byte("1111"), bytes...)
		err = b.Put(key, []byte("10000"))
		binary.BigEndian.PutUint64(bytes, 100000)
		key = append([]byte("1111"), bytes...)
		err = b.Put(key, []byte("100000"))

		//if err = tx.Commit(); err != nil {
		//	return err
		//}

		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("MyBucket")).Cursor()
		bytess := make([]byte, 8)
		binary.BigEndian.PutUint64(bytess, 100000)
		max := append([]byte("1111"), bytess...)
		for k, v := c.Seek(max); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Prev() {
			key := binary.BigEndian.Uint64(k[4:])
			fmt.Printf("%d: %s\n", key, v)
		}
		return nil
	})
	return err
}

//func BenchmarkOsStat(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		f, _ := os.Stat("chain.db")
//		f.Size()
//	}
//}
