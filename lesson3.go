package main

// AN important concept of slice that we should clear is that
//

import "fmt"

// Arrays
// length is fix at the time of creation
func array() {
	var numbs [5]int = [5]int{12, 3, 45, 66, 24}
	copyNumbs := numbs
	copyNumbs[2] = 6
	fmt.Println(numbs)
	fmt.Println(copyNumbs)
}

// Slice
// A slice is a dynamically-sized, flexible view into the elements of an array. In Go, slices are more common than arrays because they provide more functionality and flexibility. Slices are built on top of arrays and provide a more convenient way to work with sequences of data.
// In reality it's small descriptor pointing to an array with a length and capacity. The length is the number of elements in the slice, while the capacity is the number of elements in the underlying array, counting from the first element in the slice. When you append to a slice and it exceeds its capacity, a new underlying array is allocated, and the slice now points to this new array.
// we don't define the length in slice
func slice() {
	var sliceNum []int = []int{34, 45, 45}
	sliceNum = append(sliceNum, 100)
	fmt.Println(sliceNum)
	fmt.Println(len(sliceNum))
	fmt.Println(sliceNum[len(sliceNum)-1])
	fmt.Println(cap(sliceNum))

	//  making a slice with make function
	var percentage = make([]float64, 50, 100)
	percentage[0] = 12.5
	percentage[1] = 13.5
	percentage[2] = 14.5
	percentage[3] = 15.5
	percentage[4] = 16.5
	percentage[5] = 17.5
	percentage[6] = 18.5
	percentage[7] = 19.5
	// slicing a slice
	// it will include the first index and exclude the last index
	// it will return a new slice
	subPercentage := percentage[2:5]
	fmt.Println(subPercentage)

	// slicing a slice with only the first index
	subPercentage2 := percentage[2:]
	fmt.Println(subPercentage2)

	// slicing a slice with only the last index
	subPercentage3 := percentage[:5]
	fmt.Println(subPercentage3) // Output: [12.5 13.5 14.5 15.5 16.5]

	fmt.Println(cap(percentage)) // Output: 100

	// ⚠️ Slices share underlying storage
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:4]
	slice2[0] = 20

	fmt.Println(slice1) // Output: [1, 20, 3, 4, 5]
	fmt.Println(slice2) // Output: [20, 3, 4]

	// Copying a slice to another slice

	numbers := []int{1, 2, 3, 4, 5}
	copyNumbers := make([]int, len(numbers))
	copy(copyNumbers, numbers)
	fmt.Println(copyNumbers) // Output: [1, 2, 3, 4, 5]

}

// Range

func rangeLoop() {
	numbers := []int{100, 22, 35, 44, 50}

	// Using range to iterate over a slice
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value) // Output: Index: 0 Value: 100
	}

	// Using range to iterate over a slice without index
	for _, value := range numbers {
		fmt.Println("Value:", value) // Output: Value: 100
	}
}

func arrayAndValues() {
	rangeLoop()
}

// Maps
// A map is an unordered collection of key-value pairs.
func maps() {
	// Creating a map
	marks := make(map[string]float64)
	// Adding key-value pairs to the map
	marks["Math"] = 95.5
	marks["Science"] = 89.0
	marks["English"] = 92.5

	// Accessing values from the map
	fmt.Println("Math Marks:", marks["Math"])
	fmt.Println("Science Marks:", marks["Science"])
	fmt.Println("English Marks:", marks["English"])

	// Creating a map with data directly
	grades := map[string]string{
		"Sad":   "A",
		"Ali":   "B",
		"Ahmed": "C",
	}
	fmt.Println("Grades:", grades)

	// Updating a value in the map
	grades["Ali"] = "A+"
	fmt.Println("Updated Grades:", grades)

	// Deleting a key-value pair from the map
	delete(grades, "Ali")
	fmt.Println("Grades after deletion:", grades)

	// Checking if a key exists in the map
	grade := grades["Ali"]
	fmt.Println("Ali's Grade:", grade) // Output: Ali's Grade:
	grade, exists := grades["Ali"]
	fmt.Println("Ali's Grade:", grade) // Output: Ali's Grade:
	fmt.Println("Exists:", exists)     // Output: Exists: false

	if grade, exists := grades["Ali"]; exists {
		fmt.Println("Ali's Grade:", grade)
	} else {
		fmt.Println("Ali's Grade not found")
	}

	// Iterating over maps
	for subject, grade := range grades {
		fmt.Println("Subject:", subject, "Grade:", grade) // Output: Subject: Sad Grade: A
	}

	// Map with make
	studentMarks := make(map[string]float64, 5)
	fmt.Println("Student Marks:", studentMarks) // Output: Student Marks: map[]

	// Nil Maps
	var studentGrades map[string]string

	fmt.Println("Student Grades:", studentGrades) // Output: Student Grades: map[]
	// we can read any value from a nil map, but we cannot write to it.
	fmt.Println("Ali's Grade:", studentGrades["Ali"]) // Output: Ali's Grade:
	// studentGrades["Ali"] = "A+"                       // This will cause a runtime panic: assignment to entry in nil map

	// so we need to initialize the map before using it using make.
	fmt.Println(studentMarks["Ali"])
	studentMarks["Ali"] = 86

	fmt.Println("Ali's Marks:", studentMarks["Ali"]) // Output: Ali's Marks: 86

	type Student struct {
		Name  string
		Grade string
	}

	users := []Student{
		{Name: "Sad", Grade: "A"},
		{Name: "Ali", Grade: "B"},
		{Name: "Ahmed", Grade: "C"},
	}

	fmt.Println(users[0].Name) // Output: Sad

}
