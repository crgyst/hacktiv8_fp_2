package service

import (
	"context"
	"errors"
	"hacktiv8_fp_2/dto"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"
	"strconv"

	"github.com/mashingan/smapping"
)

type UserService interface {
	CreateUser(ctx context.Context, userDTO dto.UserRegisterDTO) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) (entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (s *userService) CreateUser(ctx context.Context, userDTO dto.UserRegisterDTO) (entity.User, error) {
	createdUser := entity.User{}
	err := smapping.FillStruct(&createdUser, smapping.MapFields(&userDTO))
	if err != nil {
		return createdUser, err
	}

	res, err := s.userRepository.CreateUser(ctx, createdUser)
	if err != nil {
		return createdUser, err
	}
	return res, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) (entity.User, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		return user, err
	}

	res, err := s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return user, err
	}
	return res, nil
}

func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return errors.New("incorrect id format")
	}

	err = s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
