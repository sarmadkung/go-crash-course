package main

import (
	"errors"
	"fmt"
)

// Lesson 4 — Structs, Methods & Pointers

// This lesson is critical because it teaches you how Go models data and behavior.

// By the end, you should understand:

// Structs
// Struct initialization
// Value semantics
// Methods
// Value receivers
// Pointer receivers
// Pointers
// & and *
// When Go copies data
// How to modify structs correctly

// Structs are a way to group together related data. They are similar to classes in other programming languages, but they do not have methods or inheritance. Structs are used to create complex data types that can be used to model real-world entities.

func structs() {
	type User struct {
		Name   string
		Age    int
		Height float64
	}
	// Struct initialization
	var user = User{Name: "John", Age: 30, Height: 5.11} // Struct initialization with field names
	// Struct with zero values
	var user2 User
	fmt.Println("User2 values", user2)
	fmt.Println("User Name:", user.Name) // Output: User Name: John

	// Structs as values
	// Structs are value types, which means that when you assign a struct to a new variable, a copy of the struct is created. This means that if you modify the new variable, the original struct will not be affected.
	user3 := user
	user3.Name = "Jane"
	fmt.Println("User Name:", user.Name)   // Output: User Name: John
	fmt.Println("User3 Name:", user3.Name) // Output: User3 Name: Jane
}

// Methods
// Methods are functions that are associated with a specific type. They are defined using the func keyword, followed by the receiver type, and then the method name. Methods can be used to define behavior for structs and other types.

// This type is package-level because a method needs a named package-level type
// as its receiver. It is called Person rather than User because 3quiz.go
// already declares a package-level User — two package-level types in the same
// package cannot share a name.
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

func methods() {
	person := Person{Name: "John", Age: 30}
	fmt.Println(person.Greet()) // Output: Hello, John
}

// Pointers
// Pointers are variables that store the memory address of another variable. They are used to pass data by reference, which means that when you pass a pointer to a function, the function can modify the original variable. Pointers are declared using the * operator, and the & operator is used to get the memory address of a variable.

func pointers() {
	age := 30

	pointerToAge := &age // Get the memory address of age

	fmt.Println("Age:", age)
	fmt.Println("Pointer to Age:", pointerToAge) // Output: Pointer to Age: 0xc0000140b8

	// Dereferencing the pointer to get the value
	fmt.Println("Dereferenced Pointer:", *pointerToAge) // Output: Dereferenced Pointer: 30

	// Modifying the value using the pointer
	*pointerToAge = 31
	fmt.Println("Modified Age:", age) // Output: Modified Age: 31
}

// Pointers in functions

func increase(age int) {
	// age is being passed by value, so modifying it here does not affect the original variable
	age++
}

func increasePointer(age *int) {
	// age is being passed by reference, so modifying it here affects the original variable
	*age++
}

// Pointer Receivers
// Pointer receivers are used to modify the value of the struct that the method is called on. When a method has a pointer receiver, it can modify the original struct, rather than a copy of it.

func (p *Person) HaveBirthday() {
	p.Age++
}

// BankAccount struct represents a bank account with an owner and a balance. The Deposit and Withdraw functions allow you to deposit and withdraw money from the account, respectively. They return an error if the operation is invalid (e.g., depositing a negative amount or withdrawing more than the balance).
type BankAccount struct {
	Owner   string
	Balance float64
}

func Deposit(account *BankAccount, amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	} else {
		return nil
	}

}

func Withdraw(account *BankAccount, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be positive")
	} else if amount > account.Balance {
		return errors.New("Insufficient funds")
	} else {
		account.Balance -= amount
		return nil
	}
}

// Read only Methods
// using account BankAccount as a value receiver, which means that the method cannot modify the original struct. This is useful for methods that only need to read data from the struct, rather than modify it.
func (account BankAccount) Info() string {
	return fmt.Sprintf("Owner: %s, Balance: %.2f", account.Owner, account.Balance)
}

// New new()
// New is a built-in function that allocates memory for a new value of a given type and returns a pointer to it. It is often used to create new instances of structs or other types. The new function takes a type as an argument and returns a pointer to a newly allocated zero value of that type.
var newUser = new(User) // Create a new User instance using new()
var newUser2 = &User{}  // Create a new User instance using & and a composite literal
// in modern go code you will often see the second form used more often than the first form

// 🧠 The Most Important Mental Model
// when we pass or assign a value its copied
// when we pass or assign a pointer its copied but the pointer points to the same value

// ⚠️ Important Clarification

// You'll sometimes hear:
// "Everything in Go is passed by value."
// This is technically true.
// Even pointers are passed by value.

func lesson4() {
	structs()
	methods()
	pointers()
	age := 30
	increase(age)
	fmt.Println("Age after increase function:", age) // Output: Age after increase function: 30 because age is passed by value, not by reference
	increasePointer(&age)
	fmt.Println("Age after increasePointer function:", age) // Output: Age after increasePointer function: 31 because age is passed by reference
	person := Person{Name: "John", Age: 30}
	fmt.Println("Age before birthday:", person.Age) // Output: Age before birthday: 30
	person.HaveBirthday()
	fmt.Println("Age after birthday:", person.Age) // Output: Age after birthday: 31

	// Example of a BankAccount struct with methods to deposit and withdraw money
	account := BankAccount{Owner: "John", Balance: 1000}
	err := Deposit(&account, 500)
	if err != nil {
		fmt.Println("Deposit error:", err)
	} else {
		fmt.Println("Deposit successful. New balance:", account.Balance) // Output: Deposit successful. New balance: 1500
	}

	err = Withdraw(&account, 2000)
	if err != nil {
		fmt.Println("Withdraw error:", err) // Output: Withdraw error: Insufficient funds
	} else {
		fmt.Println("Withdraw successful. New balance:", account.Balance)
	}
	err = Withdraw(&account, 500)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		fmt.Println("Withdraw successful. New balance:", account.Balance) // Output: Withdraw successful. New balance: 1000
	}
	account.Info() // Output: Owner: John, Balance: 1000.00

}
