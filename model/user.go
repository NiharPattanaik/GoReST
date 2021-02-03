package model

import (
	"fmt"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

func (user *User) IsUserValid() (bool, string) {
	var sb strings.Builder
	var isPresent bool
	if len(strings.TrimSpace(user.FirstName)) == 0 {
		sb.WriteString("<First Name>")
		isPresent = true
	}

	if len(strings.TrimSpace(user.LastName)) == 0 {
		sb.WriteString(" <Last Name>")
		isPresent = true
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		sb.WriteString(" <email>")
		isPresent = true
	}

	if isPresent {
		return false, fmt.Sprintf("%s mandatory fields are not available in the request.", sb.String())
	}

	return true, ""

}
