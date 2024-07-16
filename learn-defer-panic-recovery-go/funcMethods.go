package main

import "fmt"

// embedded Types
type Person struct {
	name string
}

type Employee struct {
	Person
}

// https://www.golang-book.com/books/intro/9
// What's the difference between a method and a function?
// In Go, the main difference between a method and a function lies in how they are associated with types.

// Definition: Functions are defined using the func keyword followed by the function name, parameters, return type, and the function body.
// Usage: Functions can be called directly by their name.

// Why would you use an embedded anonymous field instead of a normal named field?
// In Go, embedding is a way to compose types to achieve code reuse and polymorphic behavior. Instead of using inheritance (as in some other languages), Go uses composition to allow a type to include the fields and methods of another type.

// A method is similar to a function but is associated with a specific type. Methods are defined with a receiver argument, which indicates the type to which the method belongs.
// Definition: Methods are defined using the func keyword followed by a receiver (the type that the method is associated with), the method name, parameters, return type, and the method body.
// Usage: Methods are called on instances of the type to which they are associated.

func embeddedType() {
	workers := new(Employee)
	workers.name = "Saleh"
	fmt.Println(workers.Person.name) // we could access it like this
	fmt.Println(workers.name)        // or like this is going to be fine
}

type Shape struct {
	width, height int
}

func (r *Shape) area() int {
	return r.width * r.height
}

func tester() {
	rect := Shape{width: 50, height: 50}
	fmt.Println(rect.area(), "area")
}
