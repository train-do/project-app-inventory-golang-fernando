package service

import (
	"database/sql"

	"github.com/train-do/project-app-inventory-golang-fernando/handler"
	"github.com/train-do/project-app-inventory-golang-fernando/model"
	"github.com/train-do/project-app-inventory-golang-fernando/repository"
)

func getAllLog(db *sql.DB) {
	if !handler.Login(model.Request.User.Username, model.Request.User.Password) {
		return
	}
	repoGoods := repository.RepoLog{}
	response, err := repoGoods.FindAllLog(db, model.Request.Page)
	if err != nil {
		handler.InternalServerError()
		return
	}
	handler.SuccessGetAllLog(model.Request.Page, response)
}
