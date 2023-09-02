package tourplan

import (
	"database/sql"

	"github.com/ryeoman/tour/internal/app/tour_plan/repository"
	"github.com/ryeoman/tour/internal/app/tour_plan/repository/sqlite"
	"github.com/ryeoman/tour/internal/app/tour_plan/usecase"
	"github.com/ryeoman/tour/internal/infra/http/handler"
)

// NewApp construct TourPlan apps.
func NewApp(conn *sql.DB) (*handler.TourPlan, error) {
	tourPlanRepo, err := sqlite.NewTourPlanSqlite(conn)
	if err != nil {
		return nil, err
	}
	tourPlanReder := repository.NewTourPlanReader(tourPlanRepo)
	tourPlanWriter := repository.NewTourPlanWriter(tourPlanRepo)

	tourPlanUsecaseFactory := usecase.NewFactory(tourPlanReder, tourPlanWriter)
	tourPlanHandler := handler.NewTourPlan(tourPlanUsecaseFactory)

	return tourPlanHandler, nil
}
