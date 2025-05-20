package multiply

import "fmt"

func Mul(values ...int) int {
	var result int
	result = 1
	for _, value := range values {
		result *= value
		fmt.Println("result=", result)
	}
	return result
}
