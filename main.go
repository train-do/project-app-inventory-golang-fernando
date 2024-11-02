package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/train-do/project-app-inventory-golang-fernando/database"
	"github.com/train-do/project-app-inventory-golang-fernando/service"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service.RunningApp(db)
}
func init() {
	// utils.ClearScreen()
}
