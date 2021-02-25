package repository

import (
	"regexp"
	"strings"
)

//ValidatePhoneNum ..
func ValidatePhoneNum(phone string) bool {
	if !strings.Contains(phone, ")") {
		return false
	}
	code, _, _ := ParsePhone(phone)
	isValid, _ := regexp.MatchString(countries[code][1], phone)
	return isValid
}

//ParsePhone , Extract country name , code , and phone number.
func ParsePhone(phone string) (string, string, string) {

	//Extracting what's between ")" and "(" , which can be 1, 2 or 3 depending on the
	// country
	idx := strings.Index(phone, ")")
	code := phone[1:idx]
	number := strings.TrimSpace(phone[idx+1:])
	country := countries[code][0]
	return code, number, country
}

var countries = map[string][2]string{"237": {"Cameroon", `\(237\)\ ?[2368]\d{7,8}$`}, "251": {"Ethiopia", `\(251\)\ ?[1-59]\d{8}$`}, "212": {"Morocco", `\(212\)\ ?[5-9]\d{8}$`}, "258": {"Mozambique", `\(258\)\ ?[28]\d{7,8}$`}, "256": {"Uganda", `\(256\)\ ?\d{9}$`}}
