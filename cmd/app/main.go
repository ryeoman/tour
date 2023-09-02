package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	tourplan "github.com/ryeoman/tour/cmd/app/tour_plan"
	"github.com/ryeoman/tour/internal/infra/database/sqlite"
	"github.com/ryeoman/tour/internal/infra/http/server"
)

func main() {
	db, err := sqlite.NewDatabase("database/sqlite/tour.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	tourPlanApp, err := tourplan.NewApp(db.DB)

	server := server.NewHTTPServer(tourPlanApp)
	server.Start(":8000")
}
