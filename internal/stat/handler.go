package stat

import (
	"go/links-shorter/configs"
	"go/links-shorter/pkg/middleware"
	"go/links-shorter/pkg/resp"
	"net/http"
	"time"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	StatService    *StatService
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	StatService    *StatService
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		StatService:    deps.StatService,
		Config:         deps.Config,
	}
	router.Handle("GET /stat", middleware.Auth(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parse query params: from (yyyy-mm-dd), to (yyyy-mm-dd), by (day, month, year)
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		by := r.URL.Query().Get("by")

		if by == "" {
			by = "month"
		}

		if by != "day" && by != "month" && by != "year" {
			resp.Json(w, http.StatusBadRequest, "Invalid by")
			return
		}

		stats, err := handler.StatRepository.GetStat(from, to, by)
		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		resp.Json(w, http.StatusOK, stats)
	}
}
