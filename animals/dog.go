package animals

import "fmt"

// Dog represents a dog with a name.
type Dog struct {
	name string
}

// NewDog creates a new Dog with the given name.
func NewDog(name string) Dog {
	return Dog{name: name}
}

// Name returns the dog's name.
func (d Dog) Name() string { return d.name }

// Talk makes the dog bark.
func (d Dog) Talk() {
	fmt.Printf("%s says woof!\n", d.name)
}
