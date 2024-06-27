package array

import (
	"fmt"
	"testing"
)

func MinMax(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]

	for i := 0; i < len(arr); i++ {
		if min >= arr[i] {
			min = arr[i]
		}

		if max <= arr[i] {
			max = arr[i]
		}
	}

	return min, max
}

func TestMinMax(t *testing.T) {
	arr := []int{1, 2, 5, 3, 4, 6}
	min, max := MinMax(arr)

	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
