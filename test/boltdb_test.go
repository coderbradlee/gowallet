package test

import (
	"fmt"
	"testing"

	bolt "go.etcd.io/bbolt"
)

func TestYY(t *testing.T) {
	testopenfile()
}

func testopenfile() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(db)
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket11111"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b, err = tx.CreateBucket([]byte("MyBucket22222"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b, err = tx.CreateBucket([]byte("Xxxxxxxxxxx"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b, err = tx.CreateBucket([]byte("YYYYYYYYYYYYYYYYYYY"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b, err = tx.CreateBucket([]byte("MyBucket333333333"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		fmt.Println(b)
		//bytes := make([]byte, 8)
		//binary.BigEndian.PutUint64(bytes, 1)
		//err = b.Put(bytes, []byte("1"))
		//if err != nil {
		//	return err
		//}
		//binary.BigEndian.PutUint64(bytes, 10)
		//err = b.Put(bytes, []byte("10"))
		//if err != nil {
		//	return err
		//}
		//binary.BigEndian.PutUint64(bytes, 100)
		////err = b.Put([]byte{100}, []byte("zhangsan2"))
		////if err != nil {
		////	return err
		////}
		//err = b.Put(bytes, []byte("100"))
		return err
	})
	err = db.View(func(tx *bolt.Tx) error {
		if err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Println(string(name))
			return nil
		}); err != nil {
			fmt.Println(err)
		}
		return nil
	})
	//
	//fmt.Println(err)
	//err = db.Update(func(tx *bolt.Tx) error {
	//	b, err := tx.CreateBucket([]byte("MyBucket"))
	//
	//	if err != nil {
	//		return fmt.Errorf("create bucket: %s", err)
	//	}
	//
	//	err = b.Put([]byte("history1"), []byte("1"))
	//	if err != nil {
	//		return err
	//	}
	//	err = b.Put([]byte("history2"), []byte("2"))
	//	if err != nil {
	//		return err
	//	}
	//	err = b.Put([]byte("history3"), []byte("3"))
	//	if err != nil {
	//		return err
	//	}
	//	err = b.Put([]byte("history4"), []byte("4"))
	//	if err != nil {
	//		return err
	//	}
	//	err = b.Put([]byte("history5"), []byte("5"))
	//	return err
	//})
	//var allKey [][]byte
	//err = db.View(func(tx *bolt.Tx) error {
	//	c := tx.Bucket([]byte("MyBucket")).Cursor()
	//	prefix := []byte("history")
	//	for k, _ := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, _ = c.Next() {
	//		allKey = append(allKey, k)
	//	}
	//	return nil
	//})
	//for _, k := range allKey {
	//	fmt.Println(string(k))
	//}
}
