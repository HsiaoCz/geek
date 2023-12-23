package av1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/HsiaoCz/geek/what/api/define"
	"github.com/HsiaoCz/geek/what/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetBookByID(c *fiber.Ctx) error {
	id := c.Query("book-id")
	book_id, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "book id is not currect",
			"Code":    10002,
		})
	}
	conn, err := grpc.Dial(define.BookServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewWhatClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()
	book, err := client.GetBook(ctx, &pb.GetBookRequest{Identity: int64(book_id)})
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Get book successed",
		"Date":    book,
	})
}
