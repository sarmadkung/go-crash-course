package main

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

	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value) // Output: Index: 0 Value: 100
	}

	for _, value := range numbers {
		fmt.Println("Value:", value) // Output: Value: 100
	}
}

func arrayAndValues() {
	rangeLoop()
}
