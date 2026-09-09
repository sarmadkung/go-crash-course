package main

import "fmt"

// Lesson 5 — Interfaces & Composition ⭐⭐⭐

// This is one of the most important lessons in Go.
// Coming from TypeScript, Java, C#, or traditional OOP, you need to slightly change your mental model.
// In Go:
// We don't primarily use inheritance to share behavior.
// We use interfaces and composition.

// Interface

// An interface defines a set of behaviors.that a type can implement. An interface defines a contract that a type must fulfill in order to be considered as implementing that interface. In Go, interfaces are satisfied implicitly, meaning that a type does not need to explicitly declare that it implements an interface. Instead, if a type has the methods required by an interface, it is considered to implement that interface.
type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

// Using interface in function
func MakeAnimalSpeak(animal Animal) {
	animal.Speak()
}

var newDog = Dog{
	Name: "Puppy",
}

var cat = Cat{
	Name: "Ish",
}

func lesson5() {
	MakeAnimalSpeak(newDog)
	MakeAnimalSpeak(cat)
}
