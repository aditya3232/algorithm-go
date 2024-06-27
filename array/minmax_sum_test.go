package array

import (
	"fmt"
	"testing"
)

func MinMaxSum(arr []int) (summin, summax int) {
	min := arr[0]
	max := arr[0]
	var sum int

	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}

		if max < arr[i] {
			max = arr[i]
		}

		sum += i

		summin = sum - min
		summax = sum - max
	}

	return summin, summax
}

func TestSumMinMax(t *testing.T) {
	arr := []int{1, 2, 4, 5, 6}

	summin, summax := MinMaxSum(arr)

	fmt.Println("result summin: ", summin)
	fmt.Println("result summax: ", summax)
}
