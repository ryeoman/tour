package sqlite

import (
	"database/sql"

	entity "github.com/ryeoman/tour/internal/entity/tour_plan"
)

// CREATE TABLE TourPlan (
//     id INTEGER PRIMARY KEY,
//     customer_name TEXT,
//     from_location TEXT,
//     tour_package TEXT,
//     visit_plan TEXT,
//     number_of_participants INTEGER,
//     timestamp_unix INTEGER DEFAULT (strftime('%s', 'now'))
// );

// TourPlanSqlite serve as interface to TourPlan Table in sqlite.
type TourPlanSqlite struct {
	db            *sql.DB
	preparedQuery preparedQuery
}

type preparedQuery struct {
	retrieve, insert *sql.Stmt
}

// NewTourPlanSqlite construct TourPlanSqlite.
func NewTourPlanSqlite(db *sql.DB) (*TourPlanSqlite, error) {
	insertQuery := `
	INSERT INTO 
		TourPlan (customer_name, from_location, tour_package, visit_plan, number_of_participants)
	VALUES 
		(?, ?, ?, ?, ?)`
	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return nil, err
	}

	retrieveQuery := `
	SELECT 
		id, customer_name, from_location, tour_package, visit_plan, number_of_participants, timestamp_unix
	FROM 
		TourPlan
	WHERE
		id = ?
	`
	retrieveStmt, err := db.Prepare(retrieveQuery)
	if err != nil {
		return nil, err
	}

	return &TourPlanSqlite{
		db: db,
		preparedQuery: preparedQuery{
			insert:   insertStmt,
			retrieve: retrieveStmt,
		},
	}, nil
}

// GetByID get TourPlan by given id.
func (tp *TourPlanSqlite) GetByID(id int) (*entity.TourPlan, error) {
	row := tp.preparedQuery.retrieve.QueryRow(id)

	var tourPlan entity.TourPlan
	if err := row.Scan(&tourPlan.ID, &tourPlan.CustomerName, &tourPlan.FromLocation, &tourPlan.TourPackage, &tourPlan.VisitPlan, &tourPlan.NumberOfParticipants, &tourPlan.TimestampUnix); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &tourPlan, nil
}

// Insert insert TourPlan to TourPlan table.
func (tp *TourPlanSqlite) Insert(tourPlan *entity.TourPlan) error {
	_, err := tp.preparedQuery.insert.Exec(tourPlan.CustomerName, tourPlan.FromLocation, tourPlan.TourPackage, tourPlan.VisitPlan, tourPlan.NumberOfParticipants)
	return err
}
