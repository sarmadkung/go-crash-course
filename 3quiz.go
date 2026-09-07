package main

import "fmt"

// Slice manipulation and properties

func sliceManipulation() {
	numbers := []int{10, 20, 30, 40, 50}
	// length of the slice
	fmt.Println(len(numbers))
	capacity := cap(numbers)
	//  capacity of the slice
	fmt.Println(capacity)
	// appending to the slice
	numbers = append(numbers, 60, 70, 80, 90, 100)

	// length of the slice after appending
	sliceNumbs := numbers[1:4]
	fmt.Println(sliceNumbs) // Output: [20 30 40]
	// length of the slice after appending
	sliceNumbs[1] = 35
	fmt.Println(numbers)
	fmt.Println(sliceNumbs)
}
func studentSystem() {
	type Student struct {
		Name  string
		Age   int
		Grade string
	}

	students := []Student{
		{Name: "Sad", Age: 20, Grade: "A"},
		{Name: "Ali", Age: 21, Grade: "B"},
		{Name: "Ahmed", Age: 22, Grade: "C"},
	}

	for _, student := range students {
		fmt.Println("Name:", student.Name, "Age:", student.Age, "Grade:", student.Grade)
	}

}

// Frequency Counter

func frequencyCounter() {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	frequency := make(map[int]int)

	for _, number := range numbers {
		frequency[number]++
	}

	for number, count := range frequency {
		fmt.Printf("Number: %d, Frequency: %d\n", number, count)
	}
}

// User lookup

type User struct {
	Name  string
	Email string
	Age   int
}

var users = []User{
	{Name: "Sad", Email: "sad@example.com", Age: 20},
	{Name: "Ali", Email: "ali@example.com", Age: 21},
	{Name: "Ahmed", Email: "ahmed@example.com", Age: 22},
}

func findUserByEmail(email string) (User, bool) {
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return User{}, false // Return zero value and false if user not found
}

func quiz3() {
	sliceManipulation()
	studentSystem()
	user, found := findUserByEmail("e@gmail.com")
	if found {
		fmt.Println("Finding user by email:", user)
	} else {
		fmt.Println("User not found")
	}
	frequencyCounter()
}
