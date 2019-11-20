package db

import (
	"httpToJson/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

var db *sqlx.DB

func InitDB(conf config.Conf) (err error) {
	mysqlInfo:=fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Username,
	)

	db, err = sqlx.Connect("mysql", mysqlInfo)
	if err != nil{
		panic("链接数据库失败")
		return err
	}

	return nil

}

func QueryInfo(sql string) (queryinfo []config.QueryInfo) {
	var info []config.QueryInfo
	e := db.Select(&info, sql)

	if e != nil{
		fmt.Println("db.Select error : " , e)
	}
	return info
}
