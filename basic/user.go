package basic

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func NewUser(firstname, lastname string, age int) *User {
	return &User{firstname, lastname, age}
}

func (u User) Getfullname() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
