package dependency_scheduler

// Task представляет задачу с зависимостями.
type Task struct {
	ID           int
	Function     func() error
	Dependencies []int // Список ID задач, от которых зависит текущая
	// TODO: Добавьте необходимые поля
}

// DependencyScheduler управляет задачами с зависимостями.
type DependencyScheduler struct {
	// TODO: Определите необходимые поля
}

// NewDependencyScheduler создает новый планировщик с зависимостями.
func NewDependencyScheduler() *DependencyScheduler {
	// TODO: Реализуйте функцию
	return &DependencyScheduler{}
}

// AddTask добавляет задачу в планировщик.
func (ds *DependencyScheduler) AddTask(t Task) {
	// TODO: Реализуйте функцию
}

// Start запускает выполнение задач с учетом зависимостей.
func (ds *DependencyScheduler) Start() {
	// TODO: Реализуйте функцию
}
