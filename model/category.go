package model

import "time"

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
}
