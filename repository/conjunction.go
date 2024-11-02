package repository

import (
	"database/sql"
	"fmt"
)

type RepoConjuction struct {
	Id         int
	GoodsId    int
	CategoryId int
	LocationId int
}

func (r *RepoConjuction) Create(tx *sql.Tx) (int, error) {
	query := `insert into "Conjuction" (good_id, category_id, location_id) values ($1, $2, $3) returning id;`
	err := tx.QueryRow(query, r.GoodsId, r.CategoryId, r.LocationId).Scan(&r.Id)
	if err != nil {
		fmt.Println("Create Conjuction Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.Id, nil
}
