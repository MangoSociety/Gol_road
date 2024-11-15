package task_status

import "time"

// TaskStatus представляет состояние задачи.
type TaskStatus int

const (
	Pending TaskStatus = iota
	Running
	Completed
	Failed
)

// Обновите структуру Task, добавив поле для состояния.
type Task struct {
	ID       int
	RunAt    time.Time
	Function func() error // Предполагается, что функция может возвращать ошибку
	Status   TaskStatus
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

// Start запускает выполнение задач в отдельной горутине.
func (s *Scheduler) Start() {
	// TODO: Реализуйте функцию
}
