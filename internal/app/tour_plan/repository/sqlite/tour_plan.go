package sqlite

import (
	"database/sql"
	// Domain entity should not be available here.
	// It may cause domain to be changed over the time when requirement changed.
	// May need to define contract between specific database and fetcher/pusher.
	// And map it later with domain of the domain.
	// For now let's keep it like this, until the issue arise.
	domain "github.com/ryeoman/tour/internal/app/tour_plan"
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
	retrieve, retrieveList, insert *sql.Stmt
}

// NewTourPlanSqlite construct TourPlanSqlite.
func NewTourPlanSqlite(db *sql.DB) (*TourPlanSqlite, error) {
	insertQuery := `
	INSERT INTO 
		TourPlan (customer_name, from_location, tour_package, visit_plan, number_of_participants, timestamp_unix)
	VALUES 
		(?, ?, ?, ?, ?, ?)`
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

	retrieveListQuery := `
	SELECT 
		id, customer_name, from_location, tour_package, visit_plan, number_of_participants, timestamp_unix
	FROM 
		TourPlan
	WHERE
		id > ?
	LIMIT ?
	`
	retrieveListStmt, err := db.Prepare(retrieveListQuery)
	if err != nil {
		return nil, err
	}

	return &TourPlanSqlite{
		db: db,
		preparedQuery: preparedQuery{
			insert:       insertStmt,
			retrieve:     retrieveStmt,
			retrieveList: retrieveListStmt,
		},
	}, nil
}

// GetByID get TourPlan by given id.
func (tp *TourPlanSqlite) GetByID(id int) (*domain.TourPlan, error) {
	row := tp.preparedQuery.retrieve.QueryRow(id)

	var tourPlan domain.TourPlan
	if err := row.Scan(
		&tourPlan.ID,
		&tourPlan.CustomerName,
		&tourPlan.FromLocation,
		&tourPlan.TourPackage,
		&tourPlan.VisitPlan,
		&tourPlan.NumberOfParticipants,
		&tourPlan.TimestampUnix,
	); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &tourPlan, nil
}

// Get get TourPlan by given offset id.
func (tp *TourPlanSqlite) Get(limit, offset int) ([]domain.TourPlan, error) {
	rows, err := tp.preparedQuery.retrieveList.Query(offset, limit+1)
	if err != nil {
		return nil, err
	}

	results := []domain.TourPlan{}
	for rows.Next() {
		var tourPlan domain.TourPlan
		if err := rows.Scan(
			&tourPlan.ID,
			&tourPlan.CustomerName,
			&tourPlan.FromLocation,
			&tourPlan.TourPackage,
			&tourPlan.VisitPlan,
			&tourPlan.NumberOfParticipants,
			&tourPlan.TimestampUnix,
		); err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		results = append(results, tourPlan)
	}

	return results, nil
}

// Insert insert TourPlan to TourPlan table.
func (tp *TourPlanSqlite) Insert(tourPlan *domain.TourPlan) error {
	_, err := tp.preparedQuery.insert.Exec(
		tourPlan.CustomerName,
		tourPlan.FromLocation,
		tourPlan.TourPackage,
		tourPlan.VisitPlan,
		tourPlan.NumberOfParticipants,
		tourPlan.TimestampUnix,
	)
	return err
}
