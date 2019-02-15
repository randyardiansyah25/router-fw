package model

type Member struct {
	ID int64
	Name string
	Address string
	Email string
}

type Members struct {
	Data []Member
}
