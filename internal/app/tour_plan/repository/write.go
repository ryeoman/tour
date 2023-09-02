package repository

import domain "github.com/ryeoman/tour/internal/app/tour_plan"

// Note: For now we will share TourPlanWriter abstraction.

// TourPlanWriter writer
type TourPlanWriter interface {
	Insert(tourPlan *domain.TourPlan) error
}

// TourPlanWrite serve as writer to TourPlan.
type TourPlanWrite struct {
	writer TourPlanWriter
}

// NewTourPlanWriter consturct TourPlanWriter
func NewTourPlanWriter(w TourPlanWriter) *TourPlanWrite {
	return &TourPlanWrite{
		writer: w,
	}
}

// Insert TourPlan
func (tp *TourPlanWrite) Insert(tourPlan *domain.TourPlan) error {
	return tp.writer.Insert(tourPlan)
}
