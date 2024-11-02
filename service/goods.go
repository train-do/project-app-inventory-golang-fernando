package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-inventory-golang-fernando/handler"
	"github.com/train-do/project-app-inventory-golang-fernando/model"
	"github.com/train-do/project-app-inventory-golang-fernando/repository"
)

func getAllGoods(db *sql.DB) {
	repoGoods := repository.RepoGoods{}
	switch model.Request.Search.Key {
	case "name", "category", "code", "stock":
		response, err := repoGoods.FindAllGoods(db, model.Request.Page, model.Request.Search)
		if err != nil {
			handler.InternalServerError()
			return
		}
		handler.SuccessGetAllGoods(model.Request.Page, response)
	default:
		response, err := repoGoods.FindAllGoods(db, model.Request.Page)
		if err != nil {
			handler.InternalServerError()
			return
		}
		handler.SuccessGetAllGoods(model.Request.Page, response)
	}
}
func insertGoods(db *sql.DB) {
	if !handler.Login(model.Request.User.Username, model.Request.User.Password) {
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error Tx Begin")
	}
	repoGoods := repository.RepoGoods{
		Goods: model.Goods{
			Name:  model.Request.FormGoods.Name,
			Stock: model.Request.FormGoods.Stock,
		},
	}
	goods, err := repository.CreateRepo(&repoGoods, tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	repoConjuntion := repository.RepoConjuction{
		GoodsId:    goods,
		CategoryId: model.Request.FormGoods.CategoryId,
		LocationId: model.Request.FormGoods.LocationId,
	}
	_, err = repository.CreateRepo(&repoConjuntion, tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	repoLog := repository.RepoLog{
		Log: model.Log{
			GoodId:      goods,
			Information: "in",
			Qty:         model.Request.FormGoods.Stock,
		},
	}
	repoGoods.Goods.Id = goods
	_, err = repository.CreateRepo(&repoLog, tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	response, err := repoGoods.FindById(tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	tx.Commit()
	handler.SuccessCreateUpdate(response)
}
func updateGoods(db *sql.DB) {
	if !handler.Login(model.Request.User.Username, model.Request.User.Password) {
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error Tx Begin")
	}
	updateStock := model.Request.FormGoods.Stock
	repoGoods := repository.RepoGoods{
		Goods: model.Goods{
			Id: model.Request.FormGoods.Id,
		},
	}
	goods, err := repoGoods.FindById(tx)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	err = repoGoods.Update(tx, updateStock)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	var information string
	var qty int
	if updateStock < 0 {
		handler.BadRequest("Invalid Qty Stock")
		return
	} else if updateStock > goods.Stock {
		information = "in"
		qty = updateStock - goods.Stock
	} else if updateStock < goods.Stock {
		information = "out"
		qty = goods.Stock - updateStock
	} else if updateStock == goods.Stock {
		handler.BadRequest("Failed Update Because qty same as old stock")
		return
	}
	// fmt.Println(goods.Id, information, qty, "+++++++")
	repoLog := repository.RepoLog{
		Log: model.Log{
			GoodId:      goods.Id,
			Information: information,
			Qty:         qty,
		},
	}
	_, err = repository.CreateRepo(&repoLog, tx)
	// fmt.Println(logId)
	if err != nil {
		handler.BadRequest(err.Error())
		return
	}
	tx.Commit()
	handler.SuccessCreateUpdate(goods)
}
