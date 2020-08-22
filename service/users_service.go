package service

import (
	"github.com/pgrau/bookstore-user-api/domain/user"
	"github.com/pgrau/bookstore-user-api/lib/error"
)

func GetUser(userId int64) (*user.User, *error.RestErr)  {
	result := &user.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user user.User) (*user.User, *error.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
