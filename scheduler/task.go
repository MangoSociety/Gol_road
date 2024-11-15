package scheduler

import "time"

// Task представляет запланированную задачу.
type Task struct {
	ID       int
	RunAt    time.Time
	Function func()
}

// Scheduler управляет задачами.
type Scheduler struct {
	// TODO: Определите необходимые поля
}

// NewScheduler создает новый планировщик.
func NewScheduler() *Scheduler {
	// TODO: Реализуйте функцию
	return &Scheduler{}
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
