package service

import (
	"context"
	"gostrengthen/week04/api/internel/biz"
	"week04/app/book/internal/data/ent"
)

var _ biz.GRPCBookRepo = (*GRPCBookService)(nil)

type GRPCBookService struct {
	Client *ent.Client
}

func (g GRPCBookService) CreateBook(ctx context.Context, req CreateBookReq) error {
	item := req.Book
	_, err := g.Client.Book.
		Create().
		SetName(item.GetName()).
		SetAuthor(item.GetAuthor()).
		SetNumber(int(item.GetNumber())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil
}
