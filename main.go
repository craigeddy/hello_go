// Package main demonstrates the animals package by creating and interacting with Dog and Cat instances.
package main

import (
	"fmt"
	"hello_go/animals"
)

func main() {
	dog := animals.NewDog("Rex")
	fmt.Printf("Create a new dog named %s\n", dog.Name())
	cat := animals.NewCat("Whiskers")
	fmt.Printf("Create a new cat named %s\n", cat.Name())
	dog.Talk()
	cat.Talk()
}
