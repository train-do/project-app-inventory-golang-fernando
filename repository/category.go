package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-inventory-golang-fernando/model"
)

type RepoCategory struct {
	Category model.Category
}

func (r *RepoCategory) Create(tx *sql.Tx) (int, error) {
	query := `insert into "Category" (name, created_at) values ($1, $2) returning id;`
	err := tx.QueryRow(query, r.Category.Name, time.Now()).Scan(&r.Category.Id)
	if err != nil {
		fmt.Println("Create Category Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.Category.Id, nil
}
func (r *RepoCategory) FindById(tx *sql.Tx) (model.Category, error) {
	query := `select * from "Category" where id=$1;`
	err := tx.QueryRow(query, r.Category.Id).Scan(&r.Category.Id, &r.Category.Name, &r.Category.CreatedAt)
	if err != nil {
		fmt.Println("FindById Category Error: ", err)
		tx.Rollback()
		return r.Category, err
	}
	return r.Category, nil
}
