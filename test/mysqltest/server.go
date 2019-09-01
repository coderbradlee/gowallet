package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//sql.Open(s.driverName, s.connectStr+s.dbName+"?autocommit=false")
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.146.140:3306)/analytics"+"?autocommit=false")
	if err != nil {
		log.Fatal("Open() - ", err)
	}

	//stmt, err := db.Prepare("insert into `test` (`create_at`) values (?)")
	//if err != nil {
	//	log.Fatal("Prepare() - ", err)
	//}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Begin() - ", err)
	}
	for i := 0; i < 20; i++ {
		if _, err := tx.Exec("insert into `test` (`id`,`create_at`) values (?,?)", i, i); err != nil {
			log.Println("Exec() - ", err)
		}
		//log.Println("Press ENTER")
		//fmt.Scanln()
	}
	for i := 0; i < 10; i++ {
		if _, err := tx.Exec("insert into `test2` (`id`,`create_at`) values (?,?)", i, i); err != nil {
			log.Println("Exec() - ", err)
		}
	}
	for i := 8; i < 20; i++ {
		if _, err := tx.Exec("insert into `test2` (`id`,`create_at`) values (?,?)", i, i); err != nil {
			log.Println("Exec() - ", err)
		}
	}
	//if err := tx.Commit(); err != nil {
	//	log.Fatal("Commit() - ", err)
	//}
	if err := tx.Rollback(); err != nil {
		log.Fatal("Rollback() - ", err)
	}
	//stmt.Close()
	db.Close()
}
