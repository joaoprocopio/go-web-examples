package main

import "fmt"

// Go supports embedding of structs  and interfaces to express a more seamless composition of types
// this is not to be confused with //go:embed which is a god directive introduced in Go version 1.16+
// to embed files and folders into the application binary

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// a container embeds a base
// an embedding looks like a filed without a name
type container struct {
	base
	str string
}

func main() {
	// when creating structs with literals
	// we have to initialize the embedding explicitly
	// here the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// we can access the base's fields directly on co, e.g. co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// alternatively, we can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	// since containers embeds base
	// the methods of a base also become a methods of container
	// here we invoke a method that was embedded from base directly on co
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// embedding structs with methods may be used to bestow interface implementations onto other structs
	// here we see that a continare now implements the describer interface because it embeds base
	var d describer = co
	fmt.Println("describer:", d.describe())
}
