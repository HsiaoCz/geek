package service

import (
	"context"

	"github.com/HsiaoCz/geek/what/book/data"
	"github.com/HsiaoCz/geek/what/book/pb"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
}

func (b *BookService) GetBook(ctx context.Context, in *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book := &data.Book{
		Identity:   in.GetIdentity(),
		Name:       "逆天邪神",
		Title:      "玄幻",
		Auther:     "火星引力",
		CreateDate: "2014/9/8",
		TypeBanker: "纵横中文网",
		Price:      "2000",
	}
	return &pb.GetBookResponse{BookInfo: &pb.Book{Identity: book.Identity, Name: book.Name, Title: book.Title, Auther: book.Auther, CreateDate: book.CreateDate, TypeBanker: book.TypeBanker, Price: book.Price}}, nil
}

func NewBookService() *BookService {
	return &BookService{}
}
