package stat

import (
	"app/adv-http/configs"
	"app/adv-http/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("GET /stat", handler.GetStat())
}

func (h *StatHandler) GetStat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		from, err := time.Parse("2006-01-02", req.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", req.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		by := req.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid filter", http.StatusBadRequest)
			return
		}
		stats := h.StatRepository.GetStat(from, to, by)
		response.JsonResponse(w, stats, http.StatusOK)
	}
}
