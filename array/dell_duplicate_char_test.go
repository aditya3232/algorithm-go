package array

import (
	"fmt"
	"testing"
)

func DelDuplicateChar(input string) string {
	see := make(map[rune]bool)
	var result string

	for _, char := range input {
		if !see[char] {
			see[char] = true
			result += string(char)
		}
	}

	return result

}

func TestDelDuplicateChar(t *testing.T) {
	kata := "andika"

	fmt.Println("result: ", DelDuplicateChar(kata))
}
