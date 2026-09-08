package main

// Every program in Go starts with the package declaration.
// Every file belong to a package.
// main package is the entry point of the program.
// A package is a collection of files that are used to build a executable program.
// It's like a library in other programming languages and modules, namespaces in other programming languages.

// import is used to import packages that are used in the program.
// fmt is a package that is used to print output to the console.
import "fmt"

func main() {
	// main function is the entry point of the program.
	// It's the first function that is executed when the program starts.
	// It's the main function that is used to print output to the console.
	fmt.Println("Hello, World!")

	// Uncomment a lesson to run it.
	// lesson1()
	// lesson2()
	// loop100()
	arrayAndValues()
	maps()
	quiz3()
	lesson4()
}
