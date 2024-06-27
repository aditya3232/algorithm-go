package array

import (
	"fmt"
	"testing"
)

func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func TestSort(t *testing.T) {
	arr := []int{1, 4, 2, 5, 3}

	fmt.Println("before sort", arr)
	fmt.Println("after sort", Sort(arr))
}
