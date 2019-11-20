package main

import (
	"httpToJson/config"
	"fmt"
	"httpToJson/db"
	"github.com/BurntSushi/toml"
	"net/http"
	"encoding/json"
)

var (
	cfg config.Conf
)

func init() {
	_, cfg_err := toml.DecodeFile("./conf.toml", &cfg)

	if cfg_err != nil {
		fmt.Println("toml.DecodeFile error : ", cfg_err)
		return
	}
}


func httpMysqlToJson(w http.ResponseWriter, r *http.Request)  {

	e := db.InitDB(cfg)
	if e != nil{
		fmt.Println("db.InitDB " , e)
	}

	sql := "select * from userinfo"
	queryinfo := db.QueryInfo(sql)

	bytes, err := json.Marshal(queryinfo)
	if err != nil{
		fmt.Println("json.Marshal error " ,err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}


func main() {

	http.HandleFunc("/", httpMysqlToJson)
	http.ListenAndServe(":8080", nil)

}
