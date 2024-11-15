package goroutine

// Job представляет задание для воркера.
type Job struct {
	ID      int
	Payload int
}

// Result представляет результат обработки задания.
type Result struct {
	JobID  int
	Output int
}

// WorkerPool обрабатывает задания с помощью пула воркеров.
func WorkerPool(jobs <-chan Job, numWorkers int) <-chan Result {
	// TODO: Реализуйте функцию
	return make(chan Result)
}
