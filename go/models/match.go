package models

type Match struct {
	Id        uint
	Requester int
	Approver  int
	Status    string // "REQUEST", "APPROVAL", "REJECT"
}

type MatchData struct {
	Id        uint
	Requester int
	Approver  int
	Status    string // "REQUEST", "APPROVAL", "REJECT"
	Name      string
	Gender    string
	Brith     int
	Bp        int
	Sq        int
	Dl        int
}
