package repository

import (
	entity "github.com/ryeoman/tour/internal/entity/tour_plan"
)

// TourPlanRead serve as reader to read TourPlan.
type TourPlanRead struct {
	reader TourPlanReader
}

// NewTourPlanReader construct TourPlanRead.
func NewTourPlanReader(r TourPlanReader) *TourPlanRead {
	return &TourPlanRead{
		reader: r,
	}
}

// TourPlanReader reader
type TourPlanReader interface {
	GetByID(id int) (*entity.TourPlan, error)
}

// GetByID get TourPlan by given id.
func (tp *TourPlanRead) GetByID(id int) (*entity.TourPlan, error) {
	return tp.reader.GetByID(id)
}
