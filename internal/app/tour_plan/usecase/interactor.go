package usecase

import (
	"time"

	domain "github.com/ryeoman/tour/internal/app/tour_plan"
	"github.com/ryeoman/tour/internal/app/tour_plan/repository"
	"github.com/ryeoman/tour/internal/infra/presenter"
)

// Interactor TourPlan usecase.
type Interactor struct {
	reader    repository.TourPlanReader
	writer    repository.TourPlanWriter
	presenter presenter.Presenter
}

// NewTourPlan construct TourPlan usecase.
func newInteractor(
	r repository.TourPlanReader,
	w repository.TourPlanWriter,
	p presenter.Presenter,
) *Interactor {
	return &Interactor{
		reader:    r,
		writer:    w,
		presenter: p,
	}
}

// GetByID tour plan by id.
func (i *Interactor) GetByID(id int) {
	tourPlan, err := i.reader.GetByID(id)
	if err != nil {
		i.presenter.Present(nil, presenter.InternalErrorResult)
		return
	}

	i.presenter.Present(newTourPlan(tourPlan), presenter.SuccessResult)
}

// Get tour plan.
func (i *Interactor) Get(takes int) {
	if takes <= 0 {
		takes = 10
	}

	hasNext := true

	// configure maximum data fetched 10000000
	if takes > 10000000 {
		takes = 10000000
	}

	result := make([]domain.TourPlan, 0, takes)
	maxPerFetch := 10000
	for hasNext {
		offset := 0
		limit := maxPerFetch
		if takes > maxPerFetch {
			takes -= limit
		} else {
			limit = takes
		}

		tourPlans, err := i.reader.Get(limit, offset)
		if err != nil {
			i.presenter.Present(nil, presenter.InternalErrorResult)
			return
		}
		result = append(result, tourPlans...)

		if len(tourPlans) <= limit {
			hasNext = false
		}
	}

	i.presenter.Present(newTourPlans(result), presenter.SuccessResult)
}

// CreateTourPlanReq create tourPlant.
type CreateTourPlanReq struct {
	CustomerName         string
	FromLocation         string
	TourPackage          string
	VisitPlan            string
	NumberOfParticipants int
}

// Create tour plan.
func (i *Interactor) Create(req CreateTourPlanReq) {
	// Validation

	// construct domain entity
	entity := domain.TourPlan{
		CustomerName:         req.CustomerName,
		FromLocation:         req.FromLocation,
		TourPackage:          req.TourPackage,
		VisitPlan:            req.VisitPlan,
		NumberOfParticipants: req.NumberOfParticipants,
		TimestampUnix:        time.Now().UnixMilli(),
	}

	err := i.writer.Insert(&entity)
	if err != nil {
		i.presenter.Present(err, presenter.BadRequestResult)
		return
	}
	i.presenter.Present("success", presenter.SuccessResult)
}
