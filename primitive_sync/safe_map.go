package primitive_sync

import "sync"

// SafeMap представляет безопасную для конкурентного доступа мапу.
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

// NewSafeMap создает новую SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

// Get возвращает значение по ключу.
func (sm *SafeMap) Get(key string) (int, bool) {
	// TODO: Реализуйте функцию
	return 0, false
}

// Set устанавливает значение по ключу.
func (sm *SafeMap) Set(key string, value int) {
	// TODO: Реализуйте функцию
}
