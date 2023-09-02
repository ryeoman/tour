package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryeoman/tour/internal/app/tour_plan/usecase"
	http_presenter "github.com/ryeoman/tour/internal/infra/http/presenter"
	"github.com/ryeoman/tour/internal/infra/presenter"
)

// TourPlan handler for TourPlan.
type TourPlan struct {
	factory usecase.FactoryItf
}

// NewTourPlan construct handler for TourPlan.
func NewTourPlan(f usecase.FactoryItf) *TourPlan {
	return &TourPlan{
		factory: f,
	}
}

// GetByID handle get TourPlan by id.
func (h *TourPlan) GetByID(c *gin.Context) {
	// construct usecase
	httpPresenter := http_presenter.NewHTTPPresenter(c.Writer)
	tourPlanUsecase := h.factory.ConstructUsecase(httpPresenter)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpPresenter.Present(err, presenter.InternalErrorResult)
		return
	}
	tourPlanUsecase.GetByID(id)
}

// Get handle get TourPlans.
func (h *TourPlan) Get(c *gin.Context) {
	// construct usecase
	httpPresenter := http_presenter.NewHTTPPresenter(c.Writer)
	tourPlanUsecase := h.factory.ConstructUsecase(httpPresenter)

	takes, err := strconv.Atoi(c.Query("takes"))
	if err != nil {
		httpPresenter.Present(nil, presenter.BadRequestResult)
		return
	}
	tourPlanUsecase.Get(takes)
}

// Create handle create TourPlans.
func (h *TourPlan) Create(c *gin.Context) {
	// construct usecase
	httpPresenter := http_presenter.NewHTTPPresenter(c.Writer)
	tourPlanUsecase := h.factory.ConstructUsecase(httpPresenter)

	type Request struct {
		CustomerName         string `json:"customer_name"`
		FromLocation         string `json:"from_location"`
		TourPackage          string `json:"tour_package"`
		VisitPlan            string `json:"visit_plan"`
		NumberOfParticipants int    `json:"number_of_participant"`
	}

	req := Request{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		httpPresenter.Present(nil, presenter.BadRequestResult)
		return
	}

	tourPlanUsecase.Create(usecase.CreateTourPlanReq(req))
}
