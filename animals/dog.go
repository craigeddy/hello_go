package animals

import "fmt"

type Dog struct {
	name string
}

func NewDog(name string) Dog {
	return Dog{name: name}
}

func (d Dog) Name() string { return d.name }

func (d Dog) Talk() {
	fmt.Printf("%s says woof!\n", d.name)
}
