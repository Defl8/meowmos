package database

import "strconv"

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	PhoneNumber string
}

func NewUser(firstName, lastName, phoneNumber string) *User {
	return &User{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
}

// Interface compliance for bubbles list
func (u User) FilterValue() string {
	return strconv.Itoa(int(u.ID))
}

func (u User) Title() string {
	return u.LastName + ", " + u.FirstName
}

func (u User) Description() string {
	return u.PhoneNumber
}
