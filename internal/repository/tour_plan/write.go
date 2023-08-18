package repository

import (
	entity "github.com/ryeoman/tour/internal/entity/tour_plan"
)

// TourPlanWrite serve as writer to TourPlan.
type TourPlanWrite struct {
	writer TourPlanWriter
}

// TourPlanWriter writer
type TourPlanWriter interface {
	Insert(tourPlan *entity.TourPlan) error
}

// NewTourPlanWriter consturct TourPlanWriter
func NewTourPlanWriter(w TourPlanWriter) *TourPlanWrite {
	return &TourPlanWrite{
		writer: w,
	}
}

// Insert TourPlan
func (tp *TourPlanWrite) Insert(tourPlan *entity.TourPlan) error {
	return tp.writer.Insert(tourPlan)
}
