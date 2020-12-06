package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
func main(){
	driverName := "mysql"
	// dsn := "user:password@protocol(host:port)/dbname?charset=utf8mb4&loc=Local&parseTime=true"
	dsn := "root:123456@tcp(192.168.93.175:3306)/godata?charset=utf8mb4&loc=Local&parseTime=true"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer db.Close()
	err = db.Ping()
	fmt.Println(err)

	rows, err := db.Query("select id,name,password,sex,birthday,tel from user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	//  判断是否是最后一行数据
	for rows.Next(){
		var (
			id int64
			name string
			password string
			sex bool
			birthday *time.Time
			//addr string
			tel string
		)
		err := rows.Scan(&id, &name, &password, &sex, &birthday, &tel)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(id, name, password, sex, birthday, tel)
		}
	}
	var id int64
	err = db.QueryRow("select id,name,password,sex,birthday,tel from user").Scan(&id)
	fmt.Println(id, err)

}
