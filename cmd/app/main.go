package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	repository "github.com/ryeoman/tour/internal/repository/tour_plan"
	"github.com/ryeoman/tour/internal/repository/tour_plan/sqlite"
)

const file string = "database/sqlite/tour.db"

func main() {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	tourPlanSqliteReader, err := sqlite.NewTourPlanSqlite(db)
	if err != nil {
		log.Fatalln(err)
	}
	tourPlanRepository := repository.NewTourPlanReader(tourPlanSqliteReader)

	tourPlan, err := tourPlanRepository.GetByID(1)
	if err != nil {
		log.Println(tourPlan)
	}
	log.Println(tourPlan)
}
