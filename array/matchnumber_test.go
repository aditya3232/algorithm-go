package array

import (
	"fmt"
	"testing"
)

func MatchNumber(arr int) []int {
	var evenNumber []int
	var oddNumber []int

	for i := 0; i <= arr; i++ {
		if i%2 == 0 {
			evenNumber = append(evenNumber, i)
		} else {
			oddNumber = append(oddNumber, i)
		}
	}

	switch {
	case arr%2 == 0:
		return evenNumber
	case arr%2 != 0:
		return oddNumber
	default:
		return []int{0}
	}
}

func TestMatchNumber(t *testing.T) {
	fmt.Println("this is even number: ", MatchNumber(20))
	fmt.Println("this is odd number: ", MatchNumber(25))
}
