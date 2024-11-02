package service

import (
	"database/sql"

	"github.com/train-do/project-app-inventory-golang-fernando/handler"
	"github.com/train-do/project-app-inventory-golang-fernando/model"
	"github.com/train-do/project-app-inventory-golang-fernando/utils"
)

func init() {
	utils.DecodeFromJSON(&model.Request)
}
func RunningApp(db *sql.DB) {
	switch model.Request.Endpoint {
	case "allGoods":
		getAllGoods(db)
	case "addGoods":
		insertGoods(db)
	case "updateGoods":
		updateGoods(db)
	case "createCategory":
		insertCategory(db)
	case "createLocation":
		insertLocation(db)
	case "allLog":
		getAllLog(db)
	default:
		response := model.ResponseError{
			StatusCode: 404,
			Message:    "404 Not Found",
		}
		handler.NotFound(response)
	}
}
