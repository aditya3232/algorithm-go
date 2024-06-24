package algorithmgo

import (
	"fmt"
	"strings"
	"testing"
)

/*
- jumlah data slice sesuai input
- input baris = n + 3
- ada 2 bintang muncul setelah index 0, [1, 2]
- dan lanjut ke index berikutnya, [2, 3]
- sampai sebelum index terakhir

- contoh input = 5
- [1,*,*,4,5,6,7,8]
- [1,2,*,*,5,6,7,8]
- [1,2,3,*,*,6,7,8]
- [1,2,3,4,*,*,7,8]
- [1,2,3,4,5,*,*,8]

*/

func StarAppear(arr int) []string {
	var results []string
	totalLen := arr + 3

	for i := 1; i <= arr; i++ {
		var sb strings.Builder

		for j := 1; j <= totalLen; j++ {
			if j == i+1 || j == i+2 {
				sb.WriteString("*")
			} else {
				sb.WriteString(fmt.Sprintf("%d", j))
			}

			if j < totalLen {
				sb.WriteString(",")
			}
		}
		results = append(results, sb.String())
	}

	return results
}

func TestStarAppear(t *testing.T) {
	input := 5

	results := StarAppear(input)

	fmt.Println(results)

	for i, result := range results {
		fmt.Printf("result %d: %s\n", i+1, result)
	}

}
