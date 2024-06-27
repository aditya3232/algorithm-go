package array

import (
	"fmt"
	"testing"
)

func GenerateFactorial(input int) int {
	var result int = 1

	for i := input; i > 0; i-- {
		result *= i
	}

	return result
}

func TestFactorial(t *testing.T) {
	fmt.Println("result: ", GenerateFactorial(4))
}
