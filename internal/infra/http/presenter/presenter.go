package presenter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryeoman/tour/internal/infra/presenter"
)

// HTTPPresenter handle http presenter
type HTTPPresenter struct {
	writer http.ResponseWriter
}

// NewHTTPPresenter construct new http presenter
func NewHTTPPresenter(w http.ResponseWriter) *HTTPPresenter {
	return &HTTPPresenter{
		writer: w,
	}
}

// Present http response
func (p *HTTPPresenter) Present(output interface{}, t presenter.PresentType) {
	p.writer.Header().Set("Content-Type", "application/json")

	switch t {
	case presenter.BadRequestResult:
		p.writer.WriteHeader(http.StatusBadRequest)
		output = map[string]string{
			"error": "bad request",
		}
		break
	case presenter.InternalErrorResult:
		p.writer.WriteHeader(http.StatusInternalServerError)
		output = map[string]string{
			"error": "Internal server error",
		}
		break
	case presenter.SuccessResult:
		p.writer.WriteHeader(http.StatusOK)
		break
	default:
	}

	data, errMarshal := json.Marshal(map[string]interface{}{
		"data": output,
	})
	if errMarshal != nil {
		log.Println(errMarshal)
	}
	p.writer.Write(data)
}
