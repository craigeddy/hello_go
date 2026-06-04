package animals

import "fmt"

type Cat struct {
	name string
}

func NewCat(name string) Cat {
	return Cat{name: name}
}

func (c Cat) Name() string { return c.name }

func (c Cat) Talk() {
	fmt.Printf("%s says meow!\n", c.name)
}
