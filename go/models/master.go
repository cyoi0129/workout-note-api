package models

type Master struct {
	Id      uint
	UserID  int
	Name    string
	Image   string
	Type    int
	Target  int
	Muscles []int
}
