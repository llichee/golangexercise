package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	// dsn := "user:password@protocol(host:port)/dbname?charset=utf8mb4&loc=Local&parseTime=true"
	dsn := "root:123456@tcp(192.168.93.175:3306)/godata?charset=utf8mb4&loc=Local&parseTime=true"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("show tables")
	if err != nil {
		fmt.Println(err)
		return
	}
}
