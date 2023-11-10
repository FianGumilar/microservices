package interfaces

import "github.com/FianGumilar/microservices/user-service/models/dto"

type UserRepository interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, error)
	FindUserByEmail(email string) (*dto.FindUserDTO, error)
	FindUserByID(id string) (*dto.FindUserDTO, error)
	CountAllUsers() (int32, error)
	InsertUser(user *dto.InsertUserDTO) error
	UpdateUser(id string, user *dto.UpdateUserDTO) error
	DeleteUser(id string) error
}

type UserService interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, *dto.Pagination, error)
	FindUserByEmail(email string) (*dto.FindUserDTO, error)
	FindUserByID(id string) (*dto.FindUserDTO, error)
	InsertUser(user *dto.InsertUserDTO) error
	UpdateUser(id string, user *dto.UpdateUserDTO) error
	DeleteUser(id string) error
}
