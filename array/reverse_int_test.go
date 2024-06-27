package array

import (
	"fmt"
	"testing"
)

func ReverseInt(arr []int) []int {
	var result []int

	for i := 0; i < len(arr); i++ {
		result = append(result, arr[len(arr)-1-i])
	}

	return result
}

func TestReverseInt(t *testing.T) {
	arr := []int{1, 4, 3, 2}
	fmt.Println("result: ", ReverseInt(arr))
}
