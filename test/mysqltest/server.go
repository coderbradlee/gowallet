package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//sql.Open(s.driverName, s.connectStr+s.dbName+"?autocommit=false")
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.146.140:3306)/analytics"+"?autocommit=false")
	if err != nil {
		log.Fatal("Open() - ", err)
	}

	stmt, err := db.Prepare("insert into `test` (`create_at`) values (?)")
	if err != nil {
		log.Fatal("Prepare() - ", err)
	}

	for i := 0; i < 10; i++ {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal("Begin() - ", err)
		}

		if _, err := stmt.Exec(time.Now().Unix()); err != nil {
			log.Fatal("Exec() - ", err)
		}

		if err := tx.Commit(); err != nil {
			log.Fatal("Commit() - ", err)
		}

		log.Println("Press ENTER")
		fmt.Scanln()
	}

	stmt.Close()
	db.Close()
}
