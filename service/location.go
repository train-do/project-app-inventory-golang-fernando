package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-inventory-golang-fernando/handler"
	"github.com/train-do/project-app-inventory-golang-fernando/model"
	"github.com/train-do/project-app-inventory-golang-fernando/repository"
)

func insertLocation(db *sql.DB) {
	if !handler.Login(model.Request.User.Username, model.Request.User.Password) {
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error Tx Begin")
	}
	repoLocation := repository.RepoLocation{
		Location: model.Location{
			Warehouse: model.Request.FormLocation.Warehouse,
			Rack:      model.Request.FormLocation.Rack,
		},
	}
	_, err = repository.CreateRepo(&repoLocation, tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	fmt.Println(repoLocation.Location.Id)
	response, err := repoLocation.FindById(tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	tx.Commit()
	// fmt.Println(response)
	handler.SuccessCreateUpdate(response)
}
