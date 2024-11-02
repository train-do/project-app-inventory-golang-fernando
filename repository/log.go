package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-inventory-golang-fernando/model"
)

type RepoLog struct {
	Log model.Log
}

func (r *RepoLog) FindAllLog(db *sql.DB, page int) ([]model.Log, error) {
	query := `
	with "total_items" as (
		select COUNT(*) as total_items
		from "Log" l join "Goods" g on l.good_id = g.id
		)
	select l.id, g.name, l.good_id, l.information, l.qty, l.created_at, ti.total_items from "Log" l join "Goods" g on l.good_id = g.id join "total_items" ti on true order by l.id limit 5 offset $1;`
	var allLog []model.Log = []model.Log{}
	if page == 1 {
		page = 0
	} else {
		page = (page - 1) * 5
	}
	rows, err := db.Query(query, page)
	if err != nil {
		fmt.Println("Read All Goods Error Rows: ", err)
	}
	for rows.Next() {
		rows.Scan(&r.Log.Id, &r.Log.Name, &r.Log.GoodId, &r.Log.Information, &r.Log.Qty, &r.Log.CreatedAt, &r.Log.TotalItems)
		if err != nil {
			fmt.Println("Read All Log Error Next: ", err)
		}
		allLog = append(allLog, r.Log)
	}
	return allLog, nil
}
func (r *RepoLog) Create(tx *sql.Tx) (int, error) {
	query := `insert into "Log" (good_id, information, qty, created_at) values ($1, $2, $3, $4) returning id;`
	// fmt.Println(r.Log.GoodId, "-----------", r.Log.GoodId, r.Log.Information, r.Log.Qty, time.Now())
	err := tx.QueryRow(query, r.Log.GoodId, r.Log.Information, r.Log.Qty, time.Now()).Scan(&r.Log.Id)
	if err != nil {
		fmt.Println("Create Log Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.Log.Id, nil
}
