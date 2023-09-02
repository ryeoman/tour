package repository

import domain "github.com/ryeoman/tour/internal/app/tour_plan"

// Note: For now we will share TourPlanReader abstraction.

// TourPlanReader reader
type TourPlanReader interface {
	GetByID(id int) (*domain.TourPlan, error)
	Get(limit, offset int) ([]domain.TourPlan, error)
}

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

// GetByID get TourPlan by given id.
func (tp *TourPlanRead) GetByID(id int) (*domain.TourPlan, error) {
	return tp.reader.GetByID(id)
}

// Get will fetch TourPlan by given offset and limit by 10
func (tp *TourPlanRead) Get(limit, offset int) ([]domain.TourPlan, error) {
	return tp.reader.Get(limit, offset)
}
