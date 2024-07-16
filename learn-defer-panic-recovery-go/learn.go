package main

import (
	"fmt"
)

// defer implementation for panic
// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("PANIC BOSQUE")
// }

func main() {
	nums := []int{1, 2, 3, 4, 5}
	someRandomShit()
	result := sum(nums)
	fmt.Println("the result is:", result)
	fmt.Println(randomHalf(3))
	randomNumberz := []int{1, 22, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(findTheGreatest(randomNumberz...))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(findTheFibonacci(8), "fibonaceh")
	learnGoRoutines()
}

// sum is a function which takes a slice of numbers and adds them together.
// What would its function signature look like in Go?

func sum(val []int) int {
	total := 0
	for _, v := range val {
		total += v
	}
	return total
}

// Write a function which takes an integer and halves it and
// returns true if it was even or false if it was odd.
// For example half(1) should return (0, false) and half(2) should return (1, true).

func randomHalf(val int) (int, bool) {
	halved := val / 2
	isEven := val%2 == 0
	return halved, isEven
}

// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

func findTheGreatest(args ...int) int {
	greatest := args[0]
	for _, v := range args {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func findTheFibonacci(val int) int {
	if val <= 1 {
		return val
	}
	return findTheFibonacci(val-1) + findTheFibonacci(val-2)
}

// What are defer, panic and recover? How do you recover from a run-time panic?
// to recover from run time panic, we need to define defer so everytime run-time panic was happening
// we still get the value from the APP
