package main

import (
	"fmt"
)

func main() {
	arr1 := []int{2, 4, 6, 2, 5}
	arr2 := []int{5, 1, 1, 5}
	arr3 := []int{-1, 0, -1, 0, -1}

	sum1 := highestNonadjacentSum(arr1)
	sum2 := highestNonadjacentSum(arr2)
	sum3 := highestNonadjacentSum(arr3)

	if sum1 == 13 && sum2 == 10 && sum3 == 0 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func highestNonadjacentSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}

	prev, current := arr[0], max(arr[0], arr[1])	

	for i := 2; i < len(arr); i++ {
		prev, current = current, max(current, prev, prev + arr[i])	
	}

	return current
}

func max(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}