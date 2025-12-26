package database

type Query string

const (
	AddUserFull Query = "INSERT INTO users (firstname, lastname, phonenumber) VALUES (?, ?, ?)"
	DeleteUser Query = "DELETE FROM users WHERE ID = ?"
	GetUsers Query = "SELECT * FROM users"
)
