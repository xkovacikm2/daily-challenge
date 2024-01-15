package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputs := []string{"111", "312", "121", "123123", "122", "911", "123456789"}
	outputs := []int{3, 2, 3, 9, 3, 2, 3}
	for i, input := range inputs {
		if decodeWays(input) != outputs[i] {
			fmt.Println("Failed for input", input)
			return
		}
	}
	fmt.Println("Success")
}

func decodeWays(input string) int {
	n_input := len(input)

	if n_input == 0 {
		return 0
	} else if n_input == 1 {
		return 1
	} else if n_input == 2 {
		bigram, _ := strconv.Atoi(input[0:2])
		if isValidChar(bigram) {
			return 2
		}
		return 1
	}

	count := decodeWays(input[1:])
	bigram, _ := strconv.Atoi(input[0:2])
	if isValidChar(bigram){
		count += decodeWays(input[2:])
	}

	return count
}

func isValidChar(val int) bool {
	return val >= 1 && val <= 26
}

