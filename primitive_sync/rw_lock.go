package primitive_sync

// RWLock представляет Reader-Writer Lock с приоритетом для читателей.
type RWLock struct {
	// TODO: Определите необходимые поля
}

// LockReader блокирует для чтения.
func (rw *RWLock) LockReader() {
	// TODO: Реализуйте функцию
}

// UnlockReader разблокирует после чтения.
func (rw *RWLock) UnlockReader() {
	// TODO: Реализуйте функцию
}

// LockWriter блокирует для записи.
func (rw *RWLock) LockWriter() {
	// TODO: Реализуйте функцию
}

// UnlockWriter разблокирует после записи.
func (rw *RWLock) UnlockWriter() {
	// TODO: Реализуйте функцию
}
