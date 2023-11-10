package services

import (
	"math"

	"github.com/FianGumilar/microservices/user-service/interfaces"
	"github.com/FianGumilar/microservices/user-service/models/dto"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &UserService{userRepo: userRepo}
}

// DeleteUser implements interfaces.UserService.
func (s *UserService) DeleteUser(id string) error {
	_, err := s.userRepo.FindUserByID(id)
	if err != nil {
		return err
	}

	err = s.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return err
}

// FindAllUsers implements interfaces.UserService.
func (s *UserService) FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, *dto.Pagination, error) {
	count, err := s.userRepo.CountAllUsers()
	if err != nil {
		return nil, nil, err
	}

	pagination.TotalData = count
	totalPage := math.Ceil(float64(count) / float64(pagination.Limit))
	pagination.TotalPage = int32(totalPage)

	users, err := s.userRepo.FindAllUsers(pagination)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, err
}

// FindUserByEmail implements interfaces.UserService.
func (s *UserService) FindUserByEmail(email string) (*dto.FindUserDTO, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByID implements interfaces.UserService.
func (s *UserService) FindUserByID(id string) (*dto.FindUserDTO, error) {
	user, err := s.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// InsertUser implements interfaces.UserService.
func (s *UserService) InsertUser(user *dto.InsertUserDTO) error {
	user.ID = uuid.New().String()
	err := s.userRepo.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser implements interfaces.UserService.
func (s *UserService) UpdateUser(id string, user *dto.UpdateUserDTO) error {
	_, err := s.userRepo.FindUserByID(id)
	if err != nil {
		return err
	}

	err = s.userRepo.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}
