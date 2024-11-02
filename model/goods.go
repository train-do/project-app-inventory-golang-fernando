package model

import "time"

type Goods struct {
	Id         int
	Name       string
	Stock      int
	Category   Category
	Location   Location
	CreatedAt  time.Time
	TotalItems int
}
