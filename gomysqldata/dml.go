package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	dsn := "root:123456@tcp(192.168.93.175:3306)/godata?charset=utf8mb4&loc=Local&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sql := `
	INSERT INTO user(name, password, birthday) VALUES('kk', 'xxx', '1999-01-11')
`
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		fmt.Println(result)
	}
//	sql := `
//	UPDATE user SET birthday=now()
//`

}
