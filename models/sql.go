package models

import (
	"database/sql"
	"log"
	"sync"

	"github.com/astaxie/beego"
)

var sqldb *sql.DB

/*SQLConnect SQL Database connect*/
func SQLConnect() error {
	mutex := &sync.Mutex{}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
	}()
	if IsSQLConnected() {
		return nil
	}
	url := beego.AppConfig.String("sqluser") + ":" + beego.AppConfig.String("sqlpwd") +
		"@tcp(" + beego.AppConfig.String("sqlendpoint") + ":3306)/" + beego.AppConfig.String("polypandadb") + "?charset=utf8"
	db, err := sql.Open("mysql", url)
	if err != nil {
		sqldb = nil
		return err
	}
	sqldb = db
	return nil
}

/*SQLDisconnect SQL Database disconnect*/
func SQLDisconnect() {
	if GetSQLDB() != nil {
		log.Println("sqldb is going to closed")
		sqldb.Close()
	}
}

/*GetSQLDB get polypanda database*/
func GetSQLDB() *sql.DB {
	return sqldb
}

/*IsSQLConnected is sql connection still connected*/
func IsSQLConnected() bool {
	if sqldb != nil && sqldb.Ping() == nil { //sqldb is alive
		//	log.Println("SQL is connected")
		return true
	}
	return false
}
