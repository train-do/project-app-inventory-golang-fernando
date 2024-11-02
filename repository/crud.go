package repository

import (
	"database/sql"
)

type CRUD interface {
	// Read(db *sql.DB, page int, search ...model.Search) (interface{}, error)
	Create(tx *sql.Tx) (int, error)
}

//	func ReadRepo(c CRUD, db *sql.DB, page int, search ...model.Search) (interface{}, error) {
//		return c.Read(db, page, search...)
//	}
func CreateRepo(c CRUD, tx *sql.Tx) (int, error) {
	return c.Create(tx)
}
