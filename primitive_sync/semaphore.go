package primitive_sync

// Semaphore представляет семафор с заданной емкостью.
type Semaphore struct {
	// TODO: Определите необходимые поля
}

// NewSemaphore создает новый семафор с заданной емкостью.
func NewSemaphore(capacity int) *Semaphore {
	// TODO: Реализуйте функцию
	return &Semaphore{}
}

// Acquire захватывает семафор, блокируясь при необходимости.
func (s *Semaphore) Acquire() {
	// TODO: Реализуйте функцию
}

// Release освобождает семафор.
func (s *Semaphore) Release() {
	// TODO: Реализуйте функцию
}
