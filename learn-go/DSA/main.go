package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 3, 0, 4}
	sort.Ints(numbers)

	fmt.Println(findMissingNumber(numbers))
}

func findMissingNumber(numbers []int) int {

	length := len(numbers)
	expectedSum := length * (length + 1) / 2
	actualSum := 0

	for _, v := range numbers {
		actualSum += v
	}

	fmt.Println(expectedSum, actualSum)
	return expectedSum - actualSum

}
