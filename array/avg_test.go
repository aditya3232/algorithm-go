package array

import (
	"fmt"
	"testing"
)

func GenerateAvg(arr []float64) float64 {
	var result float64

	for i := 0; i < len(arr); i++ {
		result += float64(arr[i])
	}

	return result / float64(len(arr))
}

func TestAvg(t *testing.T) {
	arr := []float64{1, 2, 3, 4, 5}

	fmt.Println("result: ", GenerateAvg(arr))
}
