package models

type Users struct {
	ID int64
	Name string
	Gander bool
	Addr string
}

func NewUser(id int64, name string, gander bool, addr string) *Users {
	return &Users{id, name, gander, addr}
}
