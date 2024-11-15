package primitive_sync

// ResettableOnce представляет одноразовый инициализатор с возможностью сброса.
type ResettableOnce struct {
	// TODO: Определите необходимые поля
}

// Do выполняет функцию f только один раз, до сброса.
func (ro *ResettableOnce) Do(f func()) {
	// TODO: Реализуйте функцию
}

// Reset сбрасывает состояние, позволяя повторно выполнить функцию.
func (ro *ResettableOnce) Reset() {
	// TODO: Реализуйте функцию
}
