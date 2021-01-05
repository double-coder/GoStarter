package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	//ptr are efficient
	users  []*User
	nextID = 1
)
