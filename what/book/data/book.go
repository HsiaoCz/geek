package data

import "time"

type Book struct {
	Identity   int64
	Name       string
	Title      string
	Auther     string
	CreateDate time.Time
	TypeBanker string
	Price      string
}
