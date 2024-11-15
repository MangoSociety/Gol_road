package task_with_limit

import "time"

// Task представляет запланированную задачу.
type Task struct {
	ID       int
	RunAt    time.Time
	Function func()
}

// AddTask добавляет задачу в планировщик.
func (s *Scheduler) AddTask(t Task) {
	// TODO: Реализуйте функцию
}

// RemoveTask удаляет задачу из планировщика по ID.
func (s *Scheduler) RemoveTask(id int) {
	// TODO: Реализуйте функцию
}

// Start запускает выполнение задач в отдельной горутине.
func (s *Scheduler) Start() {
	// TODO: Реализуйте функцию
}

// Scheduler управляет задачами с ограничением на количество параллельных задач.
type Scheduler struct {
	// TODO: Добавьте необходимые поля для ограничения
}

// NewScheduler создает новый планировщик с ограничением на maxConcurrentTasks.
func NewScheduler(maxConcurrentTasks int) *Scheduler {
	// TODO: Реализуйте функцию
	return &Scheduler{}
}
