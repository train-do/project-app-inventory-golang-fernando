package handler

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/train-do/project-app-inventory-golang-fernando/model"
)

func Login(username string, password string) bool {
	if username == "admin" && password == "1234567" {
		return true
	}
	response := model.ResponseError{
		StatusCode: 401,
		Message:    "No Authentication",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
	return false
}
func SuccessCreateUpdate(data interface{}) {
	response := make(map[string]interface{})
	response["statusCode"] = 201
	response["message"] = "Data Create/Update Succesfully"
	response["data"] = data
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func SuccessGetAllGoods(page int, data []model.Goods) {
	if len(data) != 0 {
		response := make(map[string]interface{})
		response["statusCode"] = 200
		response["message"] = "Data Retrieved Succesfully"
		response["page"] = page
		response["limit"] = 5
		response["totalItems"] = data[0].TotalItems
		response["totalPages"] = math.Ceil(float64(data[0].TotalItems) / float64(5))
		response["data"] = data
		jsonData, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			fmt.Println("JSON Marshal Error")
		}
		fmt.Println(string(jsonData))
	}
}
func SuccessGetAllLog(page int, data []model.Log) {
	fmt.Printf("%+v", data)
	if len(data) != 0 {
		response := make(map[string]interface{})
		response["statusCode"] = 200
		response["message"] = "Data Retrieved Succesfully"
		response["page"] = page
		response["limit"] = 5
		response["totalItems"] = data[0].TotalItems
		response["totalPages"] = math.Ceil(float64(data[0].TotalItems) / float64(5))
		response["data"] = data
		jsonData, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			fmt.Println("JSON Marshal Error")
		}
		fmt.Println(string(jsonData))
	}
}
func NotFound(response model.ResponseError) {
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func BadRequest(msg string) {
	response := model.ResponseError{
		StatusCode: 400,
		Message:    msg,
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func InternalServerError() {
	response := model.ResponseError{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
