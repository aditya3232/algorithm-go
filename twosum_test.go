package algorithmgo

import (
	"fmt"
	"testing"
)

/*
- menjumlahkan deret data yang ada pada slice dan sesuai dengan target yang diinginkan
- jika ada 2 data yang jumlahnya sama dengan target, maka kembalikan index dari kedua data tersebut
*/

func TwoSum(arr []int, target int) []int {
	var result []int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum == target {
				result = append(result, i, j)
			}
		}
	}

	return result

}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 3, 1, 4}
	target := 6

	result := TwoSum(arr, target)
	fmt.Println("result: ", result)
}
