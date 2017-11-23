package db

import (
	"database/sql"
	"log"
	_ "github.com/nakagami/firebirdsql"
	_ "github.com/lib/pq"
	//"../utils"
	"os"
	"fmt"
)

//var connection = utils.LoadConfiguration("config.json")

func db(database string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if database == "firebirdsl" {
		fmt.Println("is Method firebirdsql query db")
		db, err = sql.Open("firebirdsql", "sysdba:masterkey@192.168.1.1:3050/D:/Arena/DB/ARENA.GDB")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		return db, err
	}
	if database == "postgres" {
		fmt.Println("is Method postgresql query db")
		db, err = sql.Open("postgres", "postgres://pguser:pg137784!@192.168.1.222:5432/mis")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		return db, err
	}

	return db, err
}

func Select(database string, Query string, Args ...interface{}) (*sql.Rows, error) {
	db, err := db(database)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(Query, Args...)
	fmt.Println(Args)
	if(err != nil){
		log.Fatal(err, Args)
		db.Close()
	}else {
		db.Close()
		return rows, err

	}
	db.Close()
	return rows, err
}

//func Update(Query string){
//	var err error
//	db, _ := db()
//	tx, _ := db.Begin()
//	_, err = tx.Exec(Query)
//	if err != nil {
//		fmt.Println(err)
//	} else{
//		tx.Commit()
//		db.Close()
//	}
//
//}
//func Insert(Query string) {
//	db, _ := db()
//	tx, _ := db.Begin()
//	_, _ = tx.Exec(Query)
//	tx.Commit()
//	db.Close()
//}


func Exec(Query string, Args ...interface{}) {
	db, errdb := db("postgres")

	//fmt.Println(Args...)
	//fmt.Println(Query)
	if errdb !=nil {
		fmt.Println(Args...)
		fmt.Println(Query)
		log.Fatal(errdb)
		os.Exit(1)
	}
	tx, txerr := db.Begin()
	if txerr == nil {
		_ , reserr := tx.Exec(Query, Args...)
		if reserr == nil {
			tx.Commit()
		} else {
			// Обрабатываем ошибку команды reserr
			fmt.Println(Args...)
			fmt.Println(Query)
			tx.Rollback()
			fmt.Println(reserr)
			os.Exit(1)

		}
	} else {
		fmt.Println(txerr)
		// Тут обрабатываем ошибку создания транзакции txerr
		os.Exit(1)
	}
	db.Close()
}
//func ExecReturValue(Query string, Args ...interface{}) ( id int) {
//	db, errdb := db()
//	//fmt.Println(Args...)
//	//fmt.Println(Query)
//	if errdb !=nil {
//		log.Fatal(errdb)
//		os.Exit(1)
//	}
//	tx, _ := db.Begin()
//		lastInsertId := 0
//		err := tx.QueryRow(Query, Args...).Scan(&lastInsertId)
//		if err != nil {
//			log.Fatal(err)
//			fmt.Println(Args...)
//			fmt.Println(Query)
//			os.Exit(1)
//		}
//		tx.Commit()
//		db.Close()
//		return lastInsertId
//}
//
//
//
//func Delete(Query string) {
//	db, _ := db()
//	tx, _ := db.Begin()
//	_, _ = tx.Exec(Query)
//	tx.Commit()
//	db.Close()
//}

