package models

import (
	// "database/sql"
	"fmt"
	// "gopkg.in/gorp.v1"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"../setting"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		// log.Fatalln(msg, err)
		log.Println(msg, err)
		return false
	}
	return true
}


func logger(obj interface{}) {
	log.Println(obj)
}

func MySQLConnect(host string, port int, user string, pass string, dbname string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Japan", user, pass, host, strconv.Itoa(port), dbname))
	checkErr(err, "sqlx.Open failed")
	return db
}

func DBConnect() *sqlx.DB {
	switch setting.GetInstance().RunMode {
	case setting.Development: 
		log.Println("dubug")
		return MySQLConnect("db", 3306, "developer", "password", "development")
	case setting.Production:
		return MySQLConnect("db", 3306, "developer", "password", "development")
	default: 
		return MySQLConnect("db", 3306, "developer", "password", "development")
	}
}

