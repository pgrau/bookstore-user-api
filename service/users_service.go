package service

import "github.com/pgrau/bookstore-user-api/domain/user"
import "github.com/pgrau/bookstore-user-api/lib/error"

func GetUser()  {

}

func CreateUser(user user.User) (*user.User, *error.RestErr) {
	return &user, nil
}

func FindUser()  {

}
