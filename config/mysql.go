package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectorStruct struct {
	mysql *sql.DB
}

func (c ConnectorStruct) MySQL() *sql.DB {
	return c.mysql
}

type ConnectorIface interface {
	MySQL() *sql.DB
}

func InitDB() ConnectorIface {
	inst := new(ConnectorStruct)
	db, err := sql.Open("mysql", os.Getenv("STRING_CONNECTION"))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	inst.mysql = db
	log.Print("InitDB() OK")
	return inst

}
