package database

import (
	"database/sql"
)

func ConnectTo(url, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AddUser(db *sql.DB, user User) error {
	err := execQuery(db, AddUserFull, user.FirstName, user.LastName, user.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}


func execQuery(db *sql.DB, query Query, values ...any) error {
	_, err := db.Exec(string(query), values...)

	if err != nil {
		return err
	}

	return nil
}
