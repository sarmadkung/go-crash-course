package main

// Question: Write a program that loops from 1 to 100.
// Rules:
// divisible by 3     → Fizz
// divisible by 5     → Buzz
// divisible by 3 & 5 → FizzBuzz
// otherwise          → number

import "fmt"

func loop100() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func calculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func calculate(a, b float64, operator string) {
	result, err := calculator(a, b, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func numAnalysis(num int) {
	var even, odd int
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)
	// sum of even and odd in string like 50 odd 50 even, total should be 5050
	fmt.Println("Sum of even and odd: ", fmt.Sprintf("%d odd %d even, total should be %d%d", odd, even, odd, even))
}

func scope() {
	x := 10

	if true {
		x := 20
		fmt.Println(x)
	}

	fmt.Println(x)
}

func quiz2() {
	loop100()
	calculate(12, 34, "*")
	numAnalysis(100)
	scope()
}
