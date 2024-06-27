package array

import (
	"fmt"
	"testing"
)

func IsPalindrome(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-1-i] {
			return false
		}
	}

	return true
}

func ReverseString(kata string) (result string) {
	for _, char := range kata {
		result = string(char) + result
	}

	return result
}

func IsPalindrome2(kata string) bool {
	if kata == ReverseString(kata) {
		return true
	} else {
		return false
	}
}

func TestIsPalindrome(t *testing.T) {
	kata := "malam"

	fmt.Println("result: ", IsPalindrome(kata))
	fmt.Println("result: ", IsPalindrome2(kata))
}
