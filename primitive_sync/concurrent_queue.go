package primitive_sync

// ConcurrentQueue представляет безопасную для конкурентного доступа очередь.
type ConcurrentQueue struct {
	// TODO: Определите необходимые поля
}

// NewConcurrentQueue создает новую очередь.
func NewConcurrentQueue() *ConcurrentQueue {
	// TODO: Реализуйте функцию
	return &ConcurrentQueue{}
}

// Enqueue добавляет элемент в очередь.
func (cq *ConcurrentQueue) Enqueue(value interface{}) {
	// TODO: Реализуйте функцию
}

// Dequeue удаляет и возвращает первый элемент из очереди.
func (cq *ConcurrentQueue) Dequeue() (interface{}, bool) {
	// TODO: Реализуйте функцию
	return nil, false
}
