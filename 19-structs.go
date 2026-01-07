package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	var p person = person{name: name}
	p.age = age

	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon", 25))

	s := person{"Sean", 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{name: "Lily", isGood: true}
	fmt.Printf("Dog %s is very %s\n", dog.name, func() string {
		if dog.isGood == true {
			return "good"
		} else {
			return "bad"
		}
	}())
}
