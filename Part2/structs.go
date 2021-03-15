package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	person := Person{name: name, age: age}

	return &person
}

func main() {
	fmt.Println(Person{"Rock", 40})
	fmt.Println(Person{age: 30, name: "Dom"})
	fmt.Println(Person{name: "AgeUnknown"})
	fmt.Println(Person{age: 30})

	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon", 30))

	person := Person{name: "Aaman", age: 24}
	fmt.Println(person.name, person.age)

	sp := &person
	fmt.Println(sp.name)

	sp.name = "Amanur"
	fmt.Println(sp)
}
