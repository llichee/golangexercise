package main

import (
	"gocode/user/controller"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"gocode/user/config"
)

func main(){
	//file, err := os.Open("config.json")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()

	addr := ":9999"
	dsn := "root:123456@tcp(192.168.93.175:3306)/users?charset=utf8mb4&loc=Local&parseTime=true"
	driverName := "mysql"
	config.InitDb(driverName, dsn)
	defer config.CloseDB()
	controller.Register()
	http.ListenAndServe(addr, nil)
}
