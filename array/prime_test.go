package array

import (
	"fmt"
	"testing"
)

func IsPrime(input int) bool {
	if input <= 1 {
		return false
	}

	for i := 2; i < input; i++ {
		if input%i == 0 {
			return false
		}
	}

	return true
}

func GeneratePrime(input int) []int {
	var result []int

	for i := 1; i <= input; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}

	return result
}

func TestIsPrime(t *testing.T) {
	fmt.Println("result: ", GeneratePrime(20))
}
