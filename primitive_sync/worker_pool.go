package primitive_sync

// Task представляет задание для воркера.
type Task func()

// WorkerPool представляет пул горутин для обработки задач.
type WorkerPool struct {
	// TODO: Определите необходимые поля
}

// NewWorkerPool создает новый пул воркеров.
func NewWorkerPool(numWorkers int) *WorkerPool {
	// TODO: Реализуйте функцию
	return &WorkerPool{}
}

// Run запускает пул воркеров для обработки задач из канала tasks.
func (wp *WorkerPool) Run(tasks <-chan Task) {
	// TODO: Реализуйте функцию
}
