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

	cleanedPhone, ok := validatePhoneNumber(user.PhoneNumber)
	if !ok {
		return errors.New("Phone number provided is invalid.")
	}

	err := execQuery(db, AddUserFull, cleanedFname, cleanedLname, cleanedPhone)
	if err != nil {
		return err
	}
	return nil
}

func validatePhoneNumber(phone string) (string, bool) {
	cleanedPhone := cleanPhone(phone)

	// Not super confident in this expression since it is just stolen from
	// GeeksForGeeks
	r, err := regexp.Compile(`^\+[1-9]\d{1,14}$`)
	if err != nil {
		return "", false
	}

	// If the first regex match fails, try again with country code added.
	if !r.MatchString(cleanedPhone) {
		cleanedPhone = "+1" + cleanedPhone
	}

	return cleanedPhone, r.MatchString(cleanedPhone)

}

func cleanData(data string) string {
	return strings.TrimSpace(data)
}

func cleanPhone(phone string) string {
	newPhone := strings.ReplaceAll(phone, "-", "")
	newPhone = strings.ReplaceAll(newPhone, ")", "")
	newPhone = strings.ReplaceAll(newPhone, "(", "")
	newPhone = strings.ReplaceAll(newPhone, " ", "")
	return newPhone
}

func execQuery(db *sql.DB, query Query, values ...any) error {
	_, err := db.Exec(string(query), values...)

	if err != nil {
		return err
	}

	return nil
}
