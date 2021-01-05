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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}
