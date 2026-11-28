package stat

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/configs"
	res "GolangCourse/advancedCourse/simpleHTTPserver/pkg"
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/middleware"
	"log"
	"net/http"
	"time"

	"gorm.io/datatypes"
)

// StatHandlerDeps содержит репозиторий статистики и настройки проверки JWT.
type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

// StatHandler обрабатывает запросы на чтение статистики переходов.
type StatHandler struct {
	StatRepository *StatRepository
}

// NewStatHandler регистрирует GET /stat с проверкой JWT.
func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

// GetStat возвращает обработчик статистики по дням или месяцам.
func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Обе границы диапазона передаются в формате YYYY-MM-DD.
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			res.JSONRes(w, "Invalid date from", http.StatusBadRequest)
			return
		}
		fromDate := datatypes.Date(from)

		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			res.JSONRes(w, "Invalid date to", http.StatusBadRequest)
			return
		}
		toDate := datatypes.Date(to)

		// Равные границы дают выборку за один день.
		if from.After(to) {
			res.JSONRes(w, "from must be on or before to", http.StatusBadRequest)
			return
		}
		// by задаёт группировку внутри выбранного диапазона.
		period := r.URL.Query().Get("by")
		if period != PeriodForDay && period != PeriodForMonth {
			res.JSONRes(w, "Invalid by: expected day or month", http.StatusBadRequest)
			return
		}

		Statistics, err := handler.StatRepository.GetBy(fromDate, toDate, period)
		if err != nil {
			log.Printf("get links: %v", err)
			res.JSONRes(w, "Failed to get links", http.StatusInternalServerError)
			return
		}

		if err := res.JSONRes(w, GetStatResponse{
			Stats: Statistics,
		},
			http.StatusOK); err != nil {
			log.Printf("write stat response: %v", err)
			return
		}
	}
}
