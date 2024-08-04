package models

type Person struct {
	Id       uint
	UserID   int
	Name     string
	Gender   string
	Brith    int
	Stations []int
	Areas    []int
	Gyms     []int
	Times    []string
	Bp       int
	Sq       int
	Dl       int
}

type PersonFilter struct {
	Stations []int
	Areas    []int
	Gyms     []int
}
