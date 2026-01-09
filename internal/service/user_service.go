package service

import (
	"errors"
	"ticket-system/internal/model"
	"ticket-system/internal/pkg/errcode"
	"ticket-system/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(username, password string) error {
	if username == "" || password == "" {
		return errcode.ErrInvalidParam
	}

	return s.userRepository.Create(&model.User{Username: username, Password: password})
}

func (s *UserService) GetByUsername(username string) (*model.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *UserService) ValidatePassword(username, password string) error {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if user.Password != password {
		return errcode.ErrUnathorized
	}

	return nil
}

func (s *UserService) Register(username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errcode.ErrInvalidParam
	}

	_, err := s.userRepository.GetByUsername(username)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		Username: username,
		Password: password,
		Role:     model.RoleUser,
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
