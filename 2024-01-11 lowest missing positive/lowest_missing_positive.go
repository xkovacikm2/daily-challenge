package main

import "fmt"

func main() {
	arr1 := []int{3, 4, -1, 1}
	arr2 := []int{1, 2, 0}
	arr3 := []int{2, 1, 3}

	result1 := findLowestMissingPositive(arr1)
	result2 := findLowestMissingPositive(arr2)
	result3 := findLowestMissingPositive(arr3)

	fmt.Println(result1, result2, result3)

	if result1 == 2 && result2 == 3 && result3 == 4 {
		println("Success")
	} else {
		println("Failure")
	}
}

// findLowestMissingPositive finds the smallest positive integer that does not exist in the array.
// The array can contain duplicates and negative numbers as well.
// The algorithm runs in O(n) time and O(1) space.
// Uses pidgeonhole principle.
func findLowestMissingPositive(nums []int) int {
	start := 1
	arr_size := len(nums)
	current := 0
	cur_val := nums[current]

	for start < arr_size {
		if cur_val > arr_size || cur_val <= 0 {
			nums[current] = 0
			current = start
			start++
			cur_val = nums[current]
			continue
		}	

		if cur_val - 1 == current {
			nums[current] = cur_val
			current = start
			start++
			cur_val = nums[current]
			continue
		}

		tmp := nums[cur_val-1]
		nums[cur_val-1] = cur_val
		cur_val = tmp
	}

	for i := 0; i < arr_size; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	return arr_size + 1
}
