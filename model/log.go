package model

import "time"

type Log struct {
	Id          int
	Name        string
	GoodId      int
	Information string
	Qty         int
	CreatedAt   time.Time
	TotalItems  int
}
