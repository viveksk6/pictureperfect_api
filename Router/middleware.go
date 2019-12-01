package Router

import (
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (handler *Handlers) CheckURLParams(h http.HandlerFunc) http.HandlerFunc {
	checkparams := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query().Get("pageNo")) == 0 || len(r.URL.Query().Get("pageSize")) == 0{
			http.Error(w, "Missing params", http.StatusBadRequest)
			return // don't call original handler
		}
		h(w, r)
	})
	return handler.Logger(checkparams)

}

func (handler *Handlers) Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handler.logger.Printf("request processed")
		h(w, r)
	}

}
