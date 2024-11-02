package model

type ResponseRead struct {
	StatusCode int
	Message    string
	Page       int
	Limit      int
	TotalItems int
	TotalPages int
	Data       interface{}
}
type ResponseSuccess struct {
	StatusCode int
	Message    string
	Data       interface{}
}
type ResponseError struct {
	StatusCode int
	Message    string
}
