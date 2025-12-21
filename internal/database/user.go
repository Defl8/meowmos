package database

type User struct {
	ID          uint 
	FirstName   string
	LastName    string
	PhoneNumber string
}


func NewUser(firstName, lastName, phoneNumber string) *User {
	return &User{
		FirstName: firstName,
		LastName: lastName,
		PhoneNumber: phoneNumber,
	}
}
