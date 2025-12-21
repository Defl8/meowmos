package database


type Query string
const (
	AddUserFull Query = "INSERT INTO users (firstname, lastname, phonenumber) VALUES (?, ?, ?)"
)
