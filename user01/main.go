package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID int64
	Name string
	Gander bool
	Addr string
}

var db *sql.DB

func NewUser(id int64, name string, gander bool, addr string) *Users {
	return &Users{id, name, gander, addr}
}
func GetUsers() []*Users {
	row, err := db.Query("select id, name, gander, addr from user;")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	users := make([]*Users,0,10)
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
		users = append(users, NewUser(id, name, gander, addr))
	}
	return users
}

func AddUser(name string, gander bool, addr string) {
	result, err := db.Exec("insert into user(name, gander, addr, created_at, updated_at) values(?, ?, ?, now(), now())", name, gander, addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func Delete(id int64) {
	result, err := db.Exec("delete from user where id=?", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func main(){
	addr := ":9999"
	dsn := "root:123456@tcp(192.168.93.175:3306)/users?charset=utf8mb4&loc=Local&parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("template/user.html"))
		tpl.ExecuteTemplate(writer, "user.html", GetUsers())
	})

	http.HandleFunc("/create/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			tpl := template.Must(template.ParseFiles("template/create.html"))
			tpl.ExecuteTemplate(writer, "create.html", nil)
		} else {
			AddUser(
				request.FormValue("name"),
				request.FormValue("gander") == "1",
				request.FormValue("addr"),
			)
			http.Redirect(writer,request,"/",302)
		}
	})
	http.HandleFunc("/delete/", func(writer http.ResponseWriter, request *http.Request) {
		if id, err := strconv.ParseInt(request.FormValue("id"), 10, 64); err == nil {
			Delete(id)
		}
		http.Redirect(writer, request, "/", 302)
	})
	http.ListenAndServe(addr, nil)
}
