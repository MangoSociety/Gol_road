package sample

// Reverse принимает слайс и возвращает новый слайс с элементами в обратном порядке.
func Reverse(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
