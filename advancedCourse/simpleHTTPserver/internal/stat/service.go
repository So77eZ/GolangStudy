package stat

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/event"
	"log"
)

// StatService обрабатывает события переходов по ссылкам.
type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}
type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

// AddClick читает события с шины и передаёт ID посещённых ссылок в репозиторий.
func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data: ", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
