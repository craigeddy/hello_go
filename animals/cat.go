package animals

import "fmt"

// Cat represents a cat with a name.
type Cat struct {
	name string
}

// NewCat creates a new Cat with the given name.
func NewCat(name string) Cat {
	return Cat{name: name}
}

// Name returns the cat's name.
func (c Cat) Name() string { return c.name }

// Talk makes the cat meow.
func (c Cat) Talk() {
	fmt.Printf("%s says meow!\n", c.name)
}
