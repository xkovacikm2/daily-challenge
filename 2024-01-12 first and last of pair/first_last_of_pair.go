package main

import "fmt"

func main() {
	pair := cons(3, 4)
	first := car(pair)
	last := cdr(pair)
	
	if first == 3 && last == 4 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

// cons takes two integers and returns a function that takes a function and applies it to the two integers.
func cons(a int, b int) func(func(int, int) int) int {
	return func(f func(int, int) int) int {
		return f(a, b)
	}
}

// car takes a pair (represented as a function) and returns the first element.
func car(pair func(func(int, int) int) int) int {
	return pair(func(a int, b int) int {
		return a
	})
}

// cdr takes a pair (represented as a function) and returns the last element.
func cdr(pair func(func(int, int) int) int) int {
	return pair(func(a int, b int) int {
		return b
	})
}