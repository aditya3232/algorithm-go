package array

import (
	"fmt"
	"testing"
)

func TestRightTriangle(t *testing.T) {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}

		fmt.Println("")
	}
}

func TestEquilateralTriangle(t *testing.T) {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Printf(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Printf("* ")
		}

		fmt.Println("")
	}
}
