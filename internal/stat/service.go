package stat

import (
	"go/links-shorter/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) ListenEvents() {
	// Фактически это чтение из канала (цикл по каналу событий)
	for e := range s.EventBus.Subscribe() {
		if e.Type == event.EventLinkVisited {
			linkId, ok := e.Data.(uint)

			if !ok {
				log.Println("Invalid event data: ", e.Data)
				continue
			}

			s.StatRepository.AddClick(linkId)
			log.Println("Click added for link: ", linkId)
		}
	}
}
