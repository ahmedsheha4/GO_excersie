package repository

import (
	"database/sql"
	"strings"

	businesslogic "github.com/ahmedsheha4/GO_excersie/businessLogic"
)

//Repository struct
type Repository struct {
	db *sql.DB
}

// SetupStorage , sqlite ..
func SetupStorage() (*Repository, error) {
	database, err := sql.Open("sqlite3", "./repository/data/sample.db")
	if err != nil {
		return &Repository{}, err
	}
	return &Repository{db: database}, nil
}

//GetPhonesFromRepository , extract related info and validation
func (r *Repository) GetPhonesFromRepository(countryFilter string, filterByState string) ([]businesslogic.PhoneNumber, error) {
	var categorizedPhoneNumbers []businesslogic.PhoneNumber
	var phone string
	rows, _ := r.db.Query("SELECT phone FROM customer")
	for rows.Next() {
		rows.Scan(&phone)
		code, number, country := ParsePhone(phone)
		isValidNum := ValidatePhoneNum(phone)
		if len(countryFilter) > 0 && !strings.Contains(strings.ToLower(country), strings.ToLower(countryFilter)) {
			continue
		}
		if len(filterByState) > 0 && ((filterByState == "1" && isValidNum == false) ||
			(filterByState == "0" && isValidNum == true)) {
			continue
		}

		categorizedPhoneNumbers = append(categorizedPhoneNumbers, businesslogic.PhoneNumber{Country: country, IsValid: isValidNum, Code: "+" + code, Number: number})
	}
	return categorizedPhoneNumbers, nil
}
