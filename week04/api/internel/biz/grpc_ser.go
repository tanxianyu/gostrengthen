package biz

import (
	"context"
)

type GRPCBookRepo interface {
	CreateBook(context.Context, *CreateBookReq) error
}

type GRPCBookBusiness struct {
	repo GRPCBookRepo
}
