package main

import "fmt"

// Go' structs are typed collections of fields
// they're useful for grouping data together to form records

// this person struct type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// you can safely return a pointer to a local variable
	// as a local variable will survive the scope of the function
	return &p
}

func main() {
	// this syntax creates a new struct
	fmt.Println(person{"Bob", 20})

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// an & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// it's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can also use dots with struct pointers
	// the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)
	fmt.Println(s.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type is only used of a single value
	// we don't have to give it a name
	// the value can have an anonymous struct type
	// this technique is commonly used for table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: true,
	}

	fmt.Println(dog)
}
