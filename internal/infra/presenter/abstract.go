package presenter

// Presenter abstraction for presenter
type Presenter interface {
	Present(output interface{}, t PresentType)
}

type PresentType int

const (
	SuccessResult PresentType = iota
	InternalErrorResult
	BadRequestResult
)
