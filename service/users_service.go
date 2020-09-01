package service

import (
	"github.com/pgrau/bookstore-user-api/domain/user"
	"github.com/pgrau/bookstore-user-api/lib/crypto"
	"github.com/pgrau/bookstore-user-api/lib/error"
)

var(
	User userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	Get(int64) (*user.User, *error.RestErr)
	FindByStatus(string) (user.Users, *error.RestErr)
	Create(user.User) (*user.User, *error.RestErr)
	Update(bool, user.User) (*user.User, *error.RestErr)
	Delete(int64) *error.RestErr
	Login(user.LoginRequest) (*user.User, *error.RestErr)
}

func (s *userService) Get(userId int64) (*user.User, *error.RestErr)  {
	result := &user.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) FindByStatus(status string) (user.Users, *error.RestErr)  {
	dao := &user.User{}
	return dao.FindByStatus(status)
}

func (s *userService) Create(user user.User) (*user.User, *error.RestErr) {
	user.DefaultValues()
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Update(isPartial bool, user user.User) (*user.User, *error.RestErr) {
	current, err := s.Get(user.Id)
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

func (s *userService) Delete(userId int64) *error.RestErr  {
	user := &user.User{Id: userId}

	return user.Delete()
}

func (s *userService) Login(request user.LoginRequest) (*user.User, *error.RestErr) {
	dao := &user.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}