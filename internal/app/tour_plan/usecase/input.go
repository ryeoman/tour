package usecase

type Input interface {
	GetByID(id int)
	Get(takes int)
	Create(req CreateTourPlanReq)
}
