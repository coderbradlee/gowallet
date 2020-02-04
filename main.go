package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lzxm160/gowallet/src/hdwallet"
	_ "github.com/mattn/go-sqlite3"
)

const (
	table_name = "table_name"
	creation   = "CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT)"
	selectSql  = "select max(id) from %s"
	cointable  = "CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, private TEXT UNIQUE, cointype TEXT,address TEXT,balance TEXT, time TIMESTAMP)"
	insert     = "INSERT OR REPLACE INTO %s (private, cointype, address,time) VALUES (?, ?, ?,?)"
)

func generateAddress() error {
	sqlDB, err := sql.Open("sqlite3", "./coin.db")
	if err != nil {
		return err
	}
	// create table to record coin'table count
	if _, err = sqlDB.Exec(fmt.Sprintf(creation, table_name)); err != nil {
		return err
	}
	stmt, err := sqlDB.Prepare(fmt.Sprintf(selectSql, table_name))
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	var name int
	err = stmt.QueryRow("1").Scan(&name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(name)
	coint := fmt.Sprintf("%s", name)
	if _, err = sqlDB.Exec(fmt.Sprintf(creation, coint)); err != nil {
		return err
	}
	stmt, err = sqlDB.Prepare(fmt.Sprintf(insert, coint))
	if err != nil {
		return err
	}
	for {
		hd := hdwallet.Hdwallet{}
		addr, pri, err := hd.GenerateAddress(0, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("btc:", addr)

		if _, err = stmt.Exec(pri, "btc", addr, time.Now().Unix()); err != nil {
			fmt.Println(err)
			continue
		}
		/////////
		addr, pri, err = hd.GenerateAddress(2, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("ltc:", addr)
		if _, err = stmt.Exec(pri, "ltc", addr, time.Now().Unix()); err != nil {
			fmt.Println(err)
			continue
		}
		///////////////////
		addr, pri, err = hd.GenerateAddress(60, 0, 0, 0)
		if err != nil {
			continue
		}
		fmt.Println("eth:", addr)
		if _, err = stmt.Exec(pri, "eth", addr, time.Now().Unix()); err != nil {
			fmt.Println(err)
			continue
		}
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	err := generateAddress()
	fmt.Println(err)
	//f, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//for {
	//	hd := hdwallet.Hdwallet{}
	//	addr, pri, err := hd.GenerateAddress(0, 0, 0, 0)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println("btc:", addr)
	//	balance, err := hdwallet.GetBTCBalanceByAddr(addr)
	//	if err != nil {
	//		fmt.Println("btc:", err)
	//		continue
	//	}
	//	if balance != "0.00000000" {
	//		f.WriteString(addr + ":" + pri + ":" + balance + "\n")
	//	}
	//
	//	/////////
	//	addr, pri, err = hd.GenerateAddress(2, 0, 0, 0)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println("ltc:", addr)
	//	balance, err = hdwallet.GetLTCBalanceByAddr(addr)
	//	//balance, err = hdwallet.GetLTCBalanceByAddr("LZ43jcFdxNVpJWJ6o3neYEsnqEGxQTsP9M")
	//	if err != nil {
	//		fmt.Println("ltc:", err)
	//		continue
	//	}
	//	if balance != "0" && balance != "" {
	//		fmt.Println("ltc balance:", balance)
	//		f.WriteString(addr + ":" + pri + ":" + balance + "\n")
	//	}
	//	///////////////////
	//	addr, pri, err = hd.GenerateAddress(60, 0, 0, 0)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println("eth:", addr)
	//	balance, err = hdwallet.GetBalance(addr)
	//	if err != nil {
	//		fmt.Println("eth:", err)
	//		continue
	//	}
	//	if balance != "0x0" {
	//		f.WriteString(addr + ":" + pri + ":" + balance + "\n")
	//	}
	//	time.Sleep(time.Second)
	//}

}
