package user

import (
	"fmt"
	"github.com/pgrau/bookstore-user-api/lib/date"
	"github.com/pgrau/bookstore-user-api/lib/error"
)

var(
	userDB = make(map[int64]*User)
)

func (user *User) Get() *error.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return error.NotFound(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *error.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return error.Conflict(fmt.Sprintf("email %s aleardy registered", user.Email))
		}

		return error.Conflict(fmt.Sprintf("user %d aleardy exists", user.Id))
	}

	user.CreatedAt = date.GetNowString()

	userDB[user.Id] = user

	return nil
}
