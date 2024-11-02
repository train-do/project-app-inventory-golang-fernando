package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-inventory-golang-fernando/model"
)

type RepoGoods struct {
	Goods model.Goods
}

func (r *RepoGoods) FindAllGoods(db *sql.DB, page int, search ...model.Search) ([]model.Goods, error) {
	query := `
	select 
		g.id,
		g.name,
		g.stock,
		g.created_at,
		c.id as category_id,
		c.name as category_name,
		c.created_at as category_created_at,
		l.id as location_id,
		l.warehouse,
		l.rack,
		l.created_at as location_created_at,
		total_items
	from 
		"Goods" g
	join 
		"Conjuction" cj on g.id = cj.good_id
	join 
		"Category" c on cj.category_id = c.id
	join 
		"Location" l on cj.location_id = l.id
	join 
		"total_items" ti on true`
	var allGoods []model.Goods = []model.Goods{}
	if page == 1 {
		page = 0
	} else {
		page = (page - 1) * 5
	}
	var rows *sql.Rows
	var err error
	if len(search) != 0 && search[0].Key == "name" {
		query = `
		with "total_items" as (
		select COUNT(*) as total_items
		from 
			"Goods" g
		join 
			"Conjuction" cj on g.id = cj.good_id
		join 
			"Category" c on cj.category_id = c.id
		join 
			"Location" l on cj.location_id = l.id
		where g.name ilike $1
		)` + query
		query += `
		where g.name ilike $2
		order by g.id 
		limit 5 offset $3;
		`
		// fmt.Println(query)
		rows, err = db.Query(query, "%"+search[0].Value+"%", "%"+search[0].Value+"%", page)
	} else if len(search) != 0 && search[0].Key == "category" {
		query = `
		with "total_items" as (
		select COUNT(*) as total_items
		from 
			"Goods" g
		join 
			"Conjuction" cj on g.id = cj.good_id
		join 
			"Category" c on cj.category_id = c.id
		join 
			"Location" l on cj.location_id = l.id
		where c.name ilike $1
		)` + query
		query += `
		where c.name ilike $2
		order by g.id 
		limit 5 offset $3;
		`
		rows, err = db.Query(query, search[0].Value, search[0].Value, page)
	} else if len(search) != 0 && search[0].Key == "code" {
		query = `
		with "total_items" as (
		select COUNT(*) as total_items
		from 
			"Goods" g
		join 
			"Conjuction" cj on g.id = cj.good_id
		join 
			"Category" c on cj.category_id = c.id
		join 
			"Location" l on cj.location_id = l.id
		where g.id = $1
		)` + query
		query += `
		where g.id = $2
		order by g.id 
		limit 5 offset $3;
		`
		rows, err = db.Query(query, search[0].Value, search[0].Value, page)
	} else if len(search) != 0 && search[0].Key == "stock" {
		query = `
		with "total_items" as (
		select COUNT(*) as total_items
		from 
			"Goods" g
		join 
			"Conjuction" cj on g.id = cj.good_id
		join 
			"Category" c on cj.category_id = c.id
		join 
			"Location" l on cj.location_id = l.id
		where g.stock < 10
		)` + query
		query += `
		where g.stock < 10
		order by g.id 
		limit 5 offset $1;
		`
		rows, err = db.Query(query, page)
	} else {
		query = `
		with "total_items" as (
		select COUNT(*) as total_items
		from 
			"Goods" g
		join 
			"Conjuction" cj on g.id = cj.good_id
		join 
			"Category" c on cj.category_id = c.id
		join 
			"Location" l on cj.location_id = l.id
		)` + query
		query += `
		order by g.id limit 5 offset $1;`
		rows, err = db.Query(query, page)
	}
	if err != nil {
		fmt.Println("Read All Goods Error Rows: ", err)
		return []model.Goods{}, err
	}
	for rows.Next() {
		var goods model.Goods
		rows.Scan(&goods.Id, &goods.Name, &goods.Stock, &goods.CreatedAt, &goods.Category.Id, &goods.Category.Name, &goods.Category.CreatedAt, &goods.Location.Id, &goods.Location.Warehouse, &goods.Location.Rack, &goods.Location.CreatedAt, &goods.TotalItems)
		if err != nil {
			fmt.Println("Read All Goods Error Next: ", err)
		}
		allGoods = append(allGoods, goods)
	}
	return allGoods, nil
}
func (r *RepoGoods) Create(tx *sql.Tx) (int, error) {
	query := `insert into "Goods" (name, stock, created_at) values ($1, $2, $3) returning id;`
	fmt.Println(r.Goods.Name)
	err := tx.QueryRow(query, r.Goods.Name, r.Goods.Stock, time.Now()).Scan(&r.Goods.Id)
	if err != nil {
		fmt.Println("Create Goods Error: ", err)
		tx.Rollback()
		return -1, err
	}
	return r.Goods.Id, nil
}
func (r *RepoGoods) Update(tx *sql.Tx, updateStock int) error {
	query := `update "Goods" set stock=$1 where id=$2 returning id;`
	fmt.Println(r.Goods.Name)
	err := tx.QueryRow(query, updateStock, r.Goods.Id).Scan(&r.Goods.Id)
	if err != nil {
		fmt.Println("Update Goods Error: ", err)
		tx.Rollback()
		return err
	}
	return nil
}
func (r *RepoGoods) FindById(tx *sql.Tx) (model.Goods, error) {
	query := `
	select 
		g.id,
		g.name,
		g.stock,
		g.created_at,
		c.id as category_id,
		c.name as category_name,
		c.created_at as category_created_at,
		l.id as location_id,
		l.warehouse,
		l.rack,
		l.created_at as location_created_at
	from 
		"Goods" g
	join 
		"Conjuction" cj on g.id = cj.good_id
	join 
		"Category" c on cj.category_id = c.id
	join 
		"Location" l on cj.location_id = l.id
	where g.id = $1
	order by g.id;`
	err := tx.QueryRow(query, r.Goods.Id).Scan(&r.Goods.Id, &r.Goods.Name, &r.Goods.Stock, &r.Goods.CreatedAt, &r.Goods.Category.Id, &r.Goods.Category.Name, &r.Goods.Category.CreatedAt, &r.Goods.Location.Id, &r.Goods.Location.Warehouse, &r.Goods.Location.Rack, &r.Goods.Location.CreatedAt)
	if err != nil {
		fmt.Println("FindById Goods Error: ", err)
		tx.Rollback()
		return r.Goods, err
	}
	return r.Goods, nil
}
