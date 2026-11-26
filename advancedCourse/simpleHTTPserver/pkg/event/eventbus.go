package event

const (
	// EventLinkVisited сообщает о переходе и передаёт ID ссылки как uint в Data.
	EventLinkVisited = "link.visited"
)

type Event struct {
	Type string
	Data any
}

// EventBus передаёт события внутри процесса через один канал без буфера.
type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

// Publish ждёт, пока читатель примет событие из канала.
func (e *EventBus) Publish(event Event) {
	e.bus <- event
}

// Subscribe возвращает общий канал для чтения. Читатели делят события между собой.
func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
