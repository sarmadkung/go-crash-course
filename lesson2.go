package main

import "fmt"

// Lesson 2 — Control Flow, Loops & Scope

// 1. if / else
// 2. for
// 3. switch
// 4. break and continue
// 5. Variable scope
// 6. Shadowing
// 7. A small practical exercise

// 1. if / else

// 1.1. if / else
func isReadyToMarry(age int) bool {
	if age >= 16 {
		fmt.Println("You are ready to get married")
		return true
	} else {
		fmt.Println("You are not ready to get married")
		return false
	}
}

// 1.2. if / else
func isReadyToRetire(age int) bool {
	if age >= 60 {
		fmt.Println("You are ready to retire")
		return true
	} else {
		fmt.Println("You are not ready to retire")
		return false
	}
}

// 1.3. if / else if / else
func whatIsMyGrade(score int) string {
	if score >= 85 {
		return "A+"
	} else if score >= 80 {
		return "A"
	} else if score >= 75 {
		return "B+"
	} else if score >= 70 {
		return "B"
	} else if score >= 65 {
		return "C+"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

// 1.4 if initialization
func isPassedTheMarriageAge() bool {

	if age := 45; age >= 35 {
		fmt.Println("You are passed the marriage age")
		return true
	} else {
		fmt.Println("You are not passed the marriage age")
		return false
	}
}

//1.4.1 if initialization with error handling

func isPassedTheMarriageAgeWithErrorHandling() bool {

	if error := doSomething(); error != nil {
		return false
	}
	return true
}

func doSomething() error {
	return nil
}

// 2. For Loop
// In go, we have two types of for loops: and we don't have while or do while loop like other languages.
// 2.1. for loop
// 2.2. for range loop

// 2.1. for loop
func printTenNumbers(numbers int) {
	for i := 0; i < numbers; i++ {
		fmt.Println(i)
	}
}

// 2.2. for range loop
func printTenNumbersUsingForRange(numbers int) {
	for i := range numbers {
		fmt.Println(i)
	}
}

// 2.3 Go's while loop equivalent

func loopWhile(while int) {
	i := 0
	for i < while {
		fmt.Println(i)
		i++
	}
}

// 3. Infinite loop
func InfiniteLoop() {
	for {
		fmt.Println("لا حول ولا قوة إلا بالله")
	}
}

// 4. Break and continue
// break: break the loop
// continue: skip the current iteration
func breakAndContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
	}
}
func continueLoop() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

// 5. Switch

// 5.1. Switch simple
func finWeekDay(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

// 5.2. Switch with multiple cases
func isWeekday(day int) string {
	switch day {
	case 1, 2, 3, 4, 5:
		return "Weekday"
	case 6, 7:
		return "Weekend"
	default:
		return "Invalid day"
	}
}

// 5.3. Switch without an expression
func isWeekdayWithoutExpression(day int) string {
	switch {
	case day >= 1 && day <= 5:
		return "Weekday"
	case day == 6 || day == 7:
		return "Weekend"
	default:
		return "Invalid day"
	}
}

// 6. Variable scope
// A block is any group of statements between { and }.
// A variable is visible only inside the block where it is declared and any nested blocks within it.
// It cannot be accessed outside that block.
// If an inner block declares a variable with the same name, it shadows the outer one within that inner block.
func variableScope() {
	var x = 10
	{
		var x = 20
		fmt.Println(x)
		var y = 30
		fmt.Println(y)
	}
	// x is still 10 here
	fmt.Println(x)
	// y is not defined here
	// fmt.Println(y)

}

//7. shadowing
// Shadowing occurs when an inner block declares a variable with the same name as a variable in an outer block.
// The inner variable shadows the outer one within that inner block.
// The outer variable is still accessible outside the inner block, but the inner variable is used within the inner block.

func shadowing() {
	outerX := 10
	{
		innerX := 20
		outerX = innerX
		fmt.Println(outerX)
	}
	fmt.Println(outerX)
}

func lesson2() {
	isReadyToMarry(25)
	isReadyToRetire(65)
	fmt.Println(whatIsMyGrade(85))
	isPassedTheMarriageAge()
	printTenNumbers(10)
	printTenNumbersUsingForRange(10)
	loopWhile(10)
	// InfiniteLoop() // never returns — uncomment only if you want to stop it with Ctrl+C
	breakAndContinue()
	continueLoop()
	fmt.Println(finWeekDay(1))
	fmt.Println(isWeekday(1))
	fmt.Println(isWeekdayWithoutExpression(1))
	variableScope()
	shadowing()
}
