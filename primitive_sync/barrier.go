package primitive_sync

// Barrier представляет барьер для синхронизации горутин.
type Barrier struct {
	// TODO: Определите необходимые поля
}

// NewBarrier создает новый барьер для заданного числа горутин.
func NewBarrier(n int) *Barrier {
	// TODO: Реализуйте функцию
	return &Barrier{}
}

// Wait блокирует горутину до тех пор, пока все горутины не достигнут барьера.
func (b *Barrier) Wait() {
	// TODO: Реализуйте функцию
}
