package main

import "hello_go/animals"

func main() {
	dog := animals.NewDog("Rex")
	cat := animals.NewCat("Whiskers")
	dog.Talk()
	cat.Talk()
}
