package user

import (
	"github.com/pgrau/bookstore-user-api/lib/crypto"
	"github.com/pgrau/bookstore-user-api/lib/date"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"regexp"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	Status 		string `json:"status"`
	Password 	string `json:"password"`
	CreatedAt 	string `json:"created_at"`
}

type Users []User

func (user *User) Validate() *error.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(user.Email) {
		return error.BadRequest("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return error.BadRequest("Invalid password")
	}

	return nil
}

func (user *User) DefaultValues() *error.RestErr {
	user.Status = StatusActive
	user.CreatedAt = date.GetNowString();
	user.Password = crypto.GetMd5(user.Password)

	return nil
}