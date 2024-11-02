package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-inventory-golang-fernando/model"
)

type RepoLocation struct {
	Location model.Location
}

func (r *RepoLocation) Create(tx *sql.Tx) (int, error) {
	query := `insert into "Location" (warehouse, rack, created_at) values ($1, $2, $3) returning id;`
	err := tx.QueryRow(query, r.Location.Warehouse, r.Location.Rack, time.Now()).Scan(&r.Location.Id)
	if err != nil {
		fmt.Println("Create Location Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.Location.Id, nil
}
func (r *RepoLocation) FindById(tx *sql.Tx) (model.Location, error) {
	query := `select * from "Location" where id=$1;`
	err := tx.QueryRow(query, r.Location.Id).Scan(&r.Location.Id, &r.Location.Warehouse, &r.Location.Rack, &r.Location.CreatedAt)
	if err != nil {
		fmt.Println("FindById Location Error: ", err)
		tx.Rollback()
		return r.Location, err
	}
	return r.Location, nil
}
