package model

type User struct {
	Username string
	Password string
}
type Search struct {
	Key   string
	Value string
}
type FormGoods struct {
	Id         int
	Name       string
	Stock      int
	CategoryId int
	LocationId int
}
type FormCategory struct {
	Name string
}
type FormLocation struct {
	Warehouse string
	Rack      string
}
type Body struct {
	Endpoint     string
	Page         int
	Search       Search
	User         User
	FormGoods    FormGoods
	FormCategory FormCategory
	FormLocation FormLocation
}

var Request Body
