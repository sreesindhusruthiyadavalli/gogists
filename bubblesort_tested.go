package main

import (
	"fmt"
)

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	//numbers := []int{3, 98, 46, 32, 1, 0}
	numbers := []int{-3, -98, -46, -32, -1, 0}
	BubbleSort(numbers)
	fmt.Print(numbers)
}
