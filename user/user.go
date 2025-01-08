package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate) // same with fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) { // utility function that create struct

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required\n")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
