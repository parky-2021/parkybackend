package utils

import "regexp"

func ValidatePhone(phone string) bool {
	Re := regexp.MustCompile(`^[0-9]{10}$`)
	return Re.MatchString(phone)

}
