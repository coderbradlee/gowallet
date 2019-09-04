package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//CREATE TABLE `test`  (
//`create_at` smallint(255) NULL DEFAULT NULL,
//`id` int(11) NOT NULL,
//PRIMARY KEY (`id`) USING BTREE
//)
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
	tx2, err := db.Begin()
	if err != nil {
		log.Fatal("Begin() - ", err)
	}
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			if _, err := tx.Exec("insert into `test` (`id`,`create_at`) values (?,?)", i, i); err != nil {
				log.Println("Exec() - ", err)
			}
		}

		//if i == 10 {
		//	if _, err := tx.Exec("CREATE TABLE IF NOT EXISTS test3 (epoch_number DECIMAL(65, 0) NOT NULL, voted_token DECIMAL(65,0) NOT NULL)"); err != nil {
		//		log.Println("Exec() - ", err)
		//	}
		//}
		//log.Println("Press ENTER")
		//fmt.Scanln()
	}
	//for i := 0; i < 10; i++ {
	//	if _, err := tx.Exec("insert into `test2` (`id`,`create_at`) values (?,?)", i, i); err != nil {
	//		log.Println("Exec() - ", err)
	//	}
	//}
	//for i := 8; i < 20; i++ {
	//	if _, err := tx.Exec("INSERT IGNORE into `test2` (`id`,`create_at`) values (?,?)", i, i); err != nil {
	//		log.Println("Exec() - ", err)
	//	}
	//	if i == 10 {
	//		if _, err := tx.Exec("CREATE TABLE IF NOT EXISTS test3 (epoch_number DECIMAL(65, 0) NOT NULL, voted_token DECIMAL(65,0) NOT NULL)"); err != nil {
	//			log.Println("Exec() - ", err)
	//		}
	//	}
	//}
	//if err := tx.Commit(); err != nil {
	//	log.Fatal("Commit() - ", err)
	//}
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			if _, err := tx2.Exec("insert into `test2` (`id`,`create_at`) values (?,?)", i, i); err != nil {
				log.Println("Exec() - ", err)
			}
		}

		//if i == 10 {
		//	if _, err := tx.Exec("CREATE TABLE IF NOT EXISTS test3 (epoch_number DECIMAL(65, 0) NOT NULL, voted_token DECIMAL(65,0) NOT NULL)"); err != nil {
		//		log.Println("Exec() - ", err)
		//	}
		//}
		//log.Println("Press ENTER")
		//fmt.Scanln()
	}
	if err := tx.Commit(); err != nil {
		log.Fatal("tx.Commit()", err)
	}
	if err := tx2.Commit(); err != nil {
		log.Fatal("tx2.Commit()", err)
	}
	//stmt.Close()
	db.Close()
}
