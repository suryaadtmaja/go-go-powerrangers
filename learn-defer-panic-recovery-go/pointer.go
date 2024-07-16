package main

import "fmt"

func someRandomShit() {
	x := 5
	y := 1
	swapInteger(&x, &y)
	// changeToZero(&x)
	fmt.Println(x, "x should be 1")
	fmt.Println(y, "y should be 5")
}

// func changeToZero(val *int) {
// 	*val = 0
// }

func swapInteger(x *int, y *int) {
	temp := *y
	*x = temp
	*y = *x
}

// How do you get the memory address of a variable?

// to get memory address of a variable in go is through
// p := &i // pointing variable p into address of i

// How do you assign a value to a pointer?
// to assign a value to pointer in go is
// *p = 21 // using * to set p through the pointer

// How do you create a new pointer?
// creating new pointer in go could be achieve by using
// newPtr := new(int)

// What is the value of x after running this program:
// func square(x *float64) {
//   *x = *x * *x
// }
// func main() {
//   x := 1.5
//   square(&x)
// }
// the x value is going to be 2.25
