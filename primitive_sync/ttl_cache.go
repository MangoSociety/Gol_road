package primitive_sync

import "time"

// CacheItem представляет элемент кэша с временем жизни.
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// TTLCache представляет потокобезопасный кэш с TTL.
type TTLCache struct {
	// TODO: Определите необходимые поля
}

// NewTTLCache создает новый кэш с заданным интервалом очистки.
func NewTTLCache(cleanupInterval time.Duration) *TTLCache {
	// TODO: Реализуйте функцию
	return &TTLCache{}
}

// Set устанавливает значение в кэше с заданным TTL.
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	// TODO: Реализуйте функцию
}

// Get получает значение из кэша по ключу.
func (c *TTLCache) Get(key string) (interface{}, bool) {
	// TODO: Реализуйте функцию
	return nil, false
}

// Delete удаляет значение из кэша по ключу.
func (c *TTLCache) Delete(key string) {
	// TODO: Реализуйте функцию
}
