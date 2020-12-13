package service

import (
	"fmt"
	"gocode/user/config"
	"gocode/user/models"
)

func AddUser(name string, gander bool, addr string) {
	result, err := config.Db.Exec("insert into user(name, gander, addr, created_at, updated_at) values(?, ?, ?, now(), now())", name, gander, addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func Delete(id int64) {
	result, err := config.Db.Exec("delete from user where id=?", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func GetUsers() []*models.Users {
	row, err := config.Db.Query("select id, name, gander, addr from user;")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	users := make([]*models.Users,0,10)
	defer row.Close()
	for row.Next() {
		var (
			id int64
			name string
			gander bool
			addr string
		)
		if err := row.Scan(&id, &name, &gander, &addr); err != nil {
			fmt.Println(err)
			break
		}
		users = append(users, models.NewUser(id, name, gander, addr))
	}
	return users
}
