package publicutils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	//"time"
)

//func MysqlConnect(connectinfo string) (db *sqlx.DB) {
//
//	db, err := sqlx.Open("mysql", connectinfo)
//	HandlerError("sqlx.Open", err)
//	return db
//
//}


func MysqlConnect(user,passwd,host,port,dbname string) (db *sqlx.DB) {

	connectinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		passwd,
		host,
		port,
		dbname,
	)

	db, err := sqlx.Open("mysql", connectinfo)
/*
	// 设置最大连接数
	db.SetMaxOpenConns(300)

	// 设置最大空闲连接数
	db.SetMaxIdleConns(100)

	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(time.Second * 10)
*/
	HandlerError("sqlx.Open", err)
	return db

}
