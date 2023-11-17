package handler

import (
	"context"

	pb "github.com/FianGumilar/microservices/user-service/infrastructures/grpc/rpc/pb"
	"github.com/FianGumilar/microservices/user-service/interfaces"
	"github.com/FianGumilar/microservices/user-service/models/dto"
	"github.com/FianGumilar/microservices/user-service/utils"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserServices interfaces.UserService
}

func NewUserHandlerUserServices(UserServices interfaces.UserService) *UserHandler {
	return &UserHandler{
		UserServices: UserServices,
	}
}

func (h *UserHandler) FindAllUsers(ctx context.Context, req *pb.PaginationRequest) (*pb.FindAllUsersResponse, error) {
	setPage := utils.SetPagination(req)

	users, page, err := h.UserServices.FindAllUsers(setPage)
	if err != nil {
		return nil, err
	}

	paginationRes := &pb.Pagination{
		Limit:     page.Limit,
		Page:      page.Page,
		TotalData: page.TotalData,
		TotalPage: page.TotalPage,
	}

	var usersRes []*pb.FindAllUsersDTO
	for _, v := range users {
		userRes := &pb.FindAllUsersDTO{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		}
		usersRes = append(usersRes, userRes)
	}

	res := pb.FindAllUsersResponse{
		Code:       200,
		Message:    "success find all users",
		Pagination: paginationRes,
		Result:     usersRes,
	}

	return &res, nil
}

func (h *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserDTO) (res *pb.UserResponse, err error) {
	// Request
	userRequest := &dto.InsertUserDTO{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	// Insert User
	err = h.UserServices.InsertUser(userRequest)
	if err != nil {
		return nil, err
	}

	res = &pb.UserResponse{
		Code:    200,
		Message: "success insert insert",
	}

	return res, nil
}

func (h *UserHandler) FindUserByEmail(ctx context.Context, req *pb.FindUsersByEmailRequest) (res *pb.FindUsersResponse, err error) {
	user, err := h.UserServices.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	userRes := &pb.FindUsersDTO{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}

	res = &pb.FindUsersResponse{
		Code:    200,
		Message: "find user by email",
		Result:  userRes,
	}

	return res, nil
}

func (h *UserHandler) FindUserById(ctx context.Context, req *pb.FindUsersByIDRequest) (res *pb.FindUsersResponse, err error) {
	user, err := h.UserServices.FindUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	userRes := &pb.FindUsersDTO{
		Id:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}

	res = &pb.FindUsersResponse{
		Code:    200,
		Message: "find user by ID",
		Result:  userRes,
	}

	return res, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	userReq := &dto.UpdateUserDTO{
		Name:      req.Name,
		Email:     req.Email,
		UpdatedAt: req.UpdatedAt,
	}

	err := h.UserServices.UpdateUser(req.Id, userReq)
	if err != nil {
		return nil, err
	}

	res := &pb.UserResponse{
		Code:    200,
		Message: "succss update user",
	}

	return res, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.UpdateUserRequest) (res *pb.UserResponse, err error) {
	err = h.UserServices.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	res = &pb.UserResponse{
		Code:    200,
		Message: "success delete user",
	}

	return res, nil
}
