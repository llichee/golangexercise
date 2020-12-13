package routers

import (
	"net/http"
	"html/template"
	"strconv"
	"gocode/user/service"
)

func GetUserscontroller(writer http.ResponseWriter, request *http.Request) {
	tpl := template.Must(template.ParseFiles("template/user.html"))
	tpl.ExecuteTemplate(writer, "user.html", service.GetUsers())
}

func AddUsercontroller(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/create.html"))
		tpl.ExecuteTemplate(writer, "create.html", nil)
	} else {
		service.AddUser(
			request.FormValue("name"),
			request.FormValue("gander") == "1",
			request.FormValue("addr"),
		)
		http.Redirect(writer,request,"/",302)
	}
}

func DeleteUsercontroller(writer http.ResponseWriter, request *http.Request) {
	if id, err := strconv.ParseInt(request.FormValue("id"), 10, 64); err == nil {
		service.Delete(id)
	}
	http.Redirect(writer, request, "/", 302)
}

