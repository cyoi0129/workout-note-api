package models

type User struct {
	Id       uint
	Email    string
	Password string
}

type UserResponse struct {
	Info  Person
	Token string
}

type LoginUser struct {
	Email    string
	Password string
}
