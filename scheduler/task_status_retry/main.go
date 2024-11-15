package task_status_retry

import "time"

// TaskStatus представляет состояние задачи.
type TaskStatus int

const (
	Pending TaskStatus = iota
	Running
	Completed
	Failed
)

// Обновите структуру Task, добавив поле для количества попыток.
type Task struct {
	ID         int
	RunAt      time.Time
	Function   func() error
	Status     TaskStatus
	RetryCount int
	MaxRetries int
	// TODO: Добавьте необходимые поля
}

// Добавьте метод для получения состояния задачи.
func (s *Scheduler) GetTaskStatus(id int) (TaskStatus, error) {
	// TODO: Реализуйте функцию
	return Pending, nil
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

// Обновите метод Start для обработки повторных попыток.
func (s *Scheduler) Start() {
	// TODO: Реализуйте функцию
}
