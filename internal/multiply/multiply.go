package multiply

func Multiply(multipliers ...int) int {
	var result int
	for _, v := range multipliers {
		result *= v
	}
	return result
}
