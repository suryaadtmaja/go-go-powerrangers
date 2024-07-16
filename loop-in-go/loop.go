package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Printf("%v is even number \n", i)
		} else {
			fmt.Printf("%v is odd number \n", i)
		}
		i++
	}

	// other type loop in go that similar like javascript

	for x := 1; x <= 100; x++ {
		fmt.Println(fizzBuzzing(x))
	}

	for z := 1; z <= 100; z++ {
		fmt.Println(challengeFive(z))
	}
}

// https://www.golang-book.com/books/intro/5
// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
func challengeFive(n int) string {
	switch {
	case n%3 == 0:
		return fmt.Sprintf("%d", n)
	default:
		return ""
	}
}

func fizzBuzzing(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzzing"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzzing"
	default:
		return fmt.Sprintf("%d", n)
	}
}
