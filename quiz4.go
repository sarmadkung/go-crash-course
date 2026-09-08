package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return "Hello, " + p.Name
}

func (p *Person) HaveBirthdayToday() {
	p.Age++
}

func (p Person) Introduce() string {
	return "My name is" + p.Name + " and I am " + fmt.Sprint(p.Age) + " years old."
}

// Shopping cart

type Cart struct {
	Items []Item
}
type Item struct {
	Name  string
	Price float64
}

func AddItem(cart *Cart, item Item) {
	cart.Items = append(cart.Items, Item{Name: item.Name, Price: item.Price})

}

func RemoveItem(cart *Cart, itemName string) {
	for i, item := range cart.Items {
		if item.Name == itemName {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
		}
	}
}

func quiz4() {
	person := Person{Name: "Alice", Age: 25}
	person.HaveBirthdayToday()
	fmt.Println(person.Greeting())                 // Output: Hello, Alice
	fmt.Println("Age after birthday:", person.Age) // Output: Age after birthday: 26
	fmt.Println(person.Introduce())                // Output: My name is Alice and I am 26 years old.
}
