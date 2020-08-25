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

func FindByStatus(status string) ([]user.User, *error.RestErr)  {
	dao := &user.User{}
	return dao.FindByStatus(status)
}

func CreateUser(user user.User) (*user.User, *error.RestErr) {
	user.DefaultValues()
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func  UpdateUser(isPartial bool, user user.User) (*user.User, *error.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err;
	}

	return current, nil;
}

func DeleteUser(userId int64) *error.RestErr  {
	user := &user.User{Id: userId}

	return user.Delete()
}