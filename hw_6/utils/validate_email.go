package utils

import "regexp"

const emailReq = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRe = regexp.MustCompile(emailReq)

func ValidateEmail(email string) bool {
	return emailRe.MatchString(email)
}
