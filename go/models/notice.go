package models

type Notice struct {
	Id     uint
	UserID int
	ChatID int
	Type   string // "REQUEST", "MESSAGE"
}
