package user

import (
	"github.com/pgrau/bookstore-user-api/lib/error"
	"strings"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	CreatedAt 	string `json:"created_at"`
}

func (user *User) Validate() *error.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return error.BadRequest("Inavalid email address")
	}

	return nil
}