package primitive_sync

import "sync"

// Counter представляет безопасный для конкурентного доступа счетчик.
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment увеличивает счетчик на 1.
func (c *Counter) Increment() {
	// TODO: Реализуйте функцию
}

// Value возвращает текущее значение счетчика.
func (c *Counter) Value() int {
	// TODO: Реализуйте функцию
	return 0
}
