package utils

import (
	"github.com/FianGumilar/microservices/user-service/infrastructures/grpc/rpc/pb"
	"github.com/FianGumilar/microservices/user-service/models/dto"
)

func SetPagination(req *pb.PaginationRequest) *dto.Pagination {
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 1
	}

	offset := (req.Page - 1) * req.Limit

	return &dto.Pagination{
		Page:   req.Page,
		Limit:  req.Limit,
		Offset: offset,
	}
}
