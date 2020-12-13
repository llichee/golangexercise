package main

import "fmt"

type Users struct {
	ID string
	name string
}

func main() {
	var id string
	fmt.Println("please input ID:")
	fmt.Scan(&id)
	users := []Users{
		{ID:"1", name:"user01"},
		{ID:"2", name:"user02"},
		{ID:"3", name:"user03"},
		{ID:"4", name:"user04"},
	}
	user := make([]Users, 0, len(users))
	fmt.Printf("%T, %#v\n", user, user)
	for _, v := range users {
		fmt.Println(v)
		if v.ID == id {
			continue
		}
		user = append(user, v)
		fmt.Println("====================")
		fmt.Println(user)
	}
	users = user
	fmt.Println(users)
}
