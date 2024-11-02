package model

import "time"

type Location struct {
	Id        int
	Warehouse string
	Rack      string
	CreatedAt time.Time
}
