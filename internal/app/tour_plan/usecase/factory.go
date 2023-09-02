package usecase

import (
	"github.com/ryeoman/tour/internal/app/tour_plan/repository"
	"github.com/ryeoman/tour/internal/infra/presenter"
)

type FactoryItf interface {
	ConstructUsecase(p presenter.Presenter) Input
}

type Factory struct {
	reader repository.TourPlanReader
	writer repository.TourPlanWriter
}

// NewFactory create factory for usecase.
func NewFactory(
	r repository.TourPlanReader,
	w repository.TourPlanWriter,
) *Factory {
	return &Factory{
		reader: r,
		writer: w,
	}
}

func (f *Factory) ConstructUsecase(p presenter.Presenter) Input {
	return newInteractor(f.reader, f.writer, p)
}
