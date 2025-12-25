package database

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
)

func ConnectTo(url, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AddUser(db *sql.DB, user User) error {
	cleanedFname := cleanData(user.FirstName)
	cleanedLname := cleanData(user.LastName)
	cleanedPhone := cleanData(user.PhoneNumber)

	ok := validatePhoneNumber(cleanedPhone)
	if !ok {
		return errors.New("Phone number provided is invalid.")
	}

	err := execQuery(db, AddUserFull, cleanedFname, cleanedLname, cleanedPhone)
	if err != nil {
		return err
	}
	return nil
}

func validatePhoneNumber(phone string) bool {
	// Not super confident in this expression since it is just stolen from
	// GeeksForGeeks
	r, err := regexp.Compile(`^[+]{1}(?:[0-9\-\(\)\/\.]\s?){6, 15}[0-9]{1}$`)
	if err != nil {
		return false
	}

	return r.MatchString(phone)
}

func cleanData(data string) string {
	return strings.TrimSpace(data)
}

func execQuery(db *sql.DB, query Query, values ...any) error {
	_, err := db.Exec(string(query), values...)

	if err != nil {
		return err
	}

	return nil
}
