package animals

import "fmt"

type Ferret struct {
	name  string
	breed string
}

func NewFerret(name string, breed string) Ferret {
	return Ferret{
		name:  name,
		breed: breed,
	}
}

func (f Ferret) Name() string { return f.name }

func (f Ferret) Breed() string { return f.breed }

func (f Ferret) Talk() {
	fmt.Printf("This %s ferret named %s says nothing...\n",
		f.Breed(), f.Name())
}
