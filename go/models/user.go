package models

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

type LoginUser struct {
	Email    string
	Password string
}
