package main

import (
	"errors"
	"fmt"
)

// ===== Variables =====

func variables() {

	// Variable declaration

	// Go is a statically typed language. typed means that the type of the variable is known at compile time.
	var age int = 30
	fmt.Printf("My age is %d\n", age)

	// Variable declaration with type inference
	var name = "Sarmad"
	fmt.Printf("My name is %s\n", name)

	// Variable declaration with type inference and short declaration operator
	// is can be only used inside a function.
	country := "Pakistan"
	fmt.Printf("My country is %s\n", country)

	// boolean
	var isStudent bool = true
	fmt.Printf("I am a student: %t\n", isStudent)

	// float
	var pi float64 = 3.14
	fmt.Printf("Pi is %f\n", pi)

	// string
	var message string = "Hello, World!"
	fmt.Printf("Message is %s\n", message)

	// array
	// array is a fixed size collection of elements of the same type.
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers are %v\n", numbers)

	// slice
	// slice is a dynamic size collection of elements of the same type.
	var fruits []string = []string{"apple", "banana", "cherry"}
	fmt.Printf("Fruits are %v\n", fruits)

	// map
	// map is a collection of key-value pairs.
	var person map[string]int = map[string]int{"age": 30, "height": 180}
	fmt.Printf("Person is %v\n", person)

	// struct
	// struct is a collection of fields of the same type.
	var user struct {
		name  string
		age   int
		email string
	}
	fmt.Printf("User is %v\n", user)

}

// ===== Data types =====

func dataTypes() {
	// we have basic data types in Go:
	// boolean
	// float64, float32
	// string
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64

	// complex64, complex128
	// array
	// slice
	// map
	// struct

	// for normal application development we use:
	// string, int, float, boolean, float64, int64

	// Special data types:
	// byte
	// rune
	// error

	// boolean
	var isStudent bool = true
	fmt.Printf("I am a student: %t\n", isStudent)

	// float
	var pi float64 = 3.14
	fmt.Printf("Pi is %f\n", pi)

	// string
	var message string = "Hello, World!"
	fmt.Printf("Message is %s\n", message)

	// int
	var age int = 30
	fmt.Printf("Age is %d\n", age)

	// uint
	var count uint = 100
	fmt.Printf("Count is %d\n", count)

	// int8
	var num int8 = 127
	fmt.Printf("Num is %d\n", num)

	// int16
	var num16 int16 = 32767
	fmt.Printf("Num16 is %d\n", num16)

	// int32
	var num32 int32 = 2147483647
	fmt.Printf("Num32 is %d\n", num32)

	// int64
	var num64 int64 = 9223372036854775807
	fmt.Printf("Num64 is %d\n", num64)

	// uint8
	var num8 uint8 = 255
	fmt.Printf("Num8 is %d\n", num8)

	// uint16
	var unum16 uint16 = 65535
	fmt.Printf("Unum16 is %d\n", unum16)

	// uint32
	var unum32 uint32 = 4294967295
	fmt.Printf("Unum32 is %d\n", unum32)

	// uint64
	var unum64 uint64 = 18446744073709551615
	fmt.Printf("Unum64 is %d\n", unum64)

	// byte
	var b byte = 255
	fmt.Printf("B is %d\n", b)

	// rune
	var r rune = 1234567890
	fmt.Printf("R is %d\n", r)

	// error
	var err error = errors.New("error")
	fmt.Printf("Err is %v\n", err)

	// complex64
	var c complex64 = 1 + 2i
	fmt.Printf("C is %v\n", c)

	// complex128
	var c128 complex128 = 1 + 2i
	fmt.Printf("C128 is %v\n", c128)

	// array
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers is %v\n", numbers)

	// slice
	var fruits []string = []string{"apple", "banana", "cherry"}
	fmt.Printf("Fruits is %v\n", fruits)

	// map
	var person map[string]int = map[string]int{"age": 30, "height": 180}
	fmt.Printf("Person is %v\n", person)

	// struct
	var user struct {
		name  string
		age   int
		email string
	} = struct {
		name  string
		age   int
		email string
	}{
		name:  "Sarmad",
		age:   30,
		email: "sarmad@example.com",
	}
	fmt.Printf("User is %v\n", user)
}

// ===== Functions =====

// A function is declared with `func`, a name, a parameter list, and a return type.
// Parameters sharing a type can be shortened: (a, b int)
func add(a int, b int) int {
	return a + b
}

// A function can return multiple values.
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Statements like `result := add(1, 2)` can only live inside a function body,
// so the calls go here rather than at the top level of the file.
func functions() {
	result := add(1, 2)
	fmt.Printf("Result is %d\n", result)

	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("DivResult is %d\n", divResult)
	}
}

// lesson1 runs every example in this lesson.
func lesson1() {
	variables()
	dataTypes()
	functions()
}
