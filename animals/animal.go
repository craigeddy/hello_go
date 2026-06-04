// Package animals provides types that represent different animals and their behaviors.
package animals

// Animal is an interface that represents an animal with a name and the ability to make sounds.
type Animal interface {
	// Name returns the animal's name.
	Name() string
	// Talk makes the animal produce its characteristic sound.
	Talk()
}
