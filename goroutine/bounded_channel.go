package goroutine

// BoundedChannel представляет ограниченный канал.
type BoundedChannel struct {
	// TODO: Определите необходимые поля
}

// NewBoundedChannel создает новый BoundedChannel заданного размера.
func NewBoundedChannel(size int) *BoundedChannel {
	// TODO: Реализуйте функцию
	return &BoundedChannel{}
}

// Send отправляет сообщение в канал, блокируясь при необходимости.
func (bc *BoundedChannel) Send(value interface{}) {
	// TODO: Реализуйте функцию
}

// Receive получает сообщение из канала, блокируясь при необходимости.
func (bc *BoundedChannel) Receive() interface{} {
	// TODO: Реализуйте функцию
	return nil
}
