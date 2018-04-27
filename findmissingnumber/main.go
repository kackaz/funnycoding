package main

import (
	"fmt"
	"log"
)

func main() {
	input := []int{1, 1, 2, 3, 2, 3, 9, 7, 7, 9,
		21, 22, 22, 22, 22, 22, 23, 23, 23, 23,
		24, 24, 24, 24, 26, 26, 26, 26, 26, 26,
		26, 26, 28, 28, 28, 28, 28, 28, 28, 30}

	max := 0
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	log.Printf("max = %d\n", max)

	result := make([]int, max)
	for i := 0; i < len(input); i++ {
		result[input[i]-1] += 1
	}
	log.Printf("result.length = %d\n", len(result))
	fmt.Println("Missing numbers:")
	for i := 0; i < len(result); i++ {
		if result[i] == 0 {
			fmt.Printf("%d\n", (i + 1))
		}
	}
}
