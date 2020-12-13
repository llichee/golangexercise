package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Users struct {
	ID int64
	Name string
	Gander bool
	Addr string
}

func main(){
	addr := ":9999"
	users := []*Users{
		{1,"xzw", true, "18"},
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("template/user.html"))
		tpl.ExecuteTemplate(writer, "user.html", users)
	})

	http.HandleFunc("/create/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			tpl := template.Must(template.ParseFiles("template/create.html"))
			tpl.ExecuteTemplate(writer, "create.html", nil)
		} else {
			users = append(users, &Users{
				time.Now().Unix(),
				request.FormValue("name"),
				request.FormValue("gander") == "1",
				request.FormValue("addr"),
			})
			http.Redirect(writer,request,"/",302)
		}
	})
	http.HandleFunc("/delete/", func(writer http.ResponseWriter, request *http.Request) {
		if id, err := strconv.ParseInt(request.FormValue("id"), 10, 64); err == nil {
			nUsers := make([]*Users, 0, len(users))
			for _, user := range users {
				if user.ID == id {
					continue
				}
				nUsers = append(nUsers, user)
			}
			users = nUsers
		}
		http.Redirect(writer, request, "/", 302)
	})
	http.ListenAndServe(addr, nil)
}
