package models

type Master struct {
	Areas    []Area
	Gyms     []Gym
	Lines    []LineStation
	Stations []Station
	Menus    []Menu
	Muscles  []Muscle
}

type Area struct {
	Id   uint
	Name string
}

type Gym struct {
	Id   uint
	Name string
}

type Station struct {
	Id     uint
	Name   string
	LineID int
}

type Line struct {
	Id   uint
	Name string
}

type LineStation struct {
	Id       uint
	Name     string
	Stations []Station
}

type Menu struct {
	Id      uint
	Name    string
	Image   string
	Type    string
	Target  int
	Muscles []int
}

type Muscle struct {
	Id   uint
	Part string
	Name string
}
