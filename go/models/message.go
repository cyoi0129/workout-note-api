package models

type Message struct {
	Id       uint
	ChatID   int
	Sender   int
	Receiver int
	Content  string
	Date     string
}
