package controller

import (
	"gocode/user/routers"
	"net/http"
)
func Register() {
	http.HandleFunc("/", routers.GetUserscontroller)
	http.HandleFunc("/create/", routers.AddUsercontroller)
	http.HandleFunc("/delete/", routers.DeleteUsercontroller)
}