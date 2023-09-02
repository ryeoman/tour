package usecase

import (
	"time"

	domain "github.com/ryeoman/tour/internal/app/tour_plan"
)

func newTourPlan(tp *domain.TourPlan) TourPlan {
	return TourPlan{
		ID:                   tp.ID,
		CustomerName:         tp.CustomerName,
		FromLocation:         tp.FromLocation,
		TourPackage:          tp.TourPackage,
		VisitPlan:            tp.VisitPlan,
		NumberOfParticipants: tp.NumberOfParticipants,
		TimestampUnix:        time.UnixMilli(tp.TimestampUnix).Format("2006-01-02 15:04:05 MST"),
	}
}

func newTourPlans(tps []domain.TourPlan) []TourPlan {
	tourPlans := make([]TourPlan, len(tps))
	for i, tourPlan := range tps {
		tourPlans[i] = newTourPlan(&tourPlan)
	}
	return tourPlans
}

type TourPlan struct {
	ID                   int    `json:"id"`
	CustomerName         string `json:"customer_name"`
	FromLocation         string `json:"from_location"`
	TourPackage          string `json:"tour_package"`
	VisitPlan            string `json:"visit_plan"`
	NumberOfParticipants int    `json:"number_of_participant"`
	TimestampUnix        string `json:"createrd_at"` // We may want to format timestampt to more readable format like `yyyy/MM/dd hh:mm`
}
