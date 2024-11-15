package primitive_sync

// EventLoop представляет цикл событий.
type EventLoop struct {
	// TODO: Определите необходимые поля
}

// NewEventLoop создает новый цикл событий.
func NewEventLoop() *EventLoop {
	// TODO: Реализуйте функцию
	return &EventLoop{}
}

// RegisterChannel регистрирует новый канал для обработки событий.
func (el *EventLoop) RegisterChannel(ch <-chan interface{}) {
	// TODO: Реализуйте функцию
}

// Start запускает цикл событий.
func (el *EventLoop) Start() {
	// TODO: Реализуйте функцию
}

// Stop останавливает цикл событий.
func (el *EventLoop) Stop() {
	// TODO: Реализуйте функцию
}
