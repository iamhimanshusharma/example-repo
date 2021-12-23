package main

import (
	"fmt"

	b "learnGithub/maths"
)

func main() {
	fmt.Println("Hello Himanshu!")
	fmt.Println("Adding newline")
	fmt.Println("This line will itself in dev branch.")

	newlife("Himanshu")

	// fmt.Println("Your addition is: ", prt)

	added := b.Add(5, 6)

	fmt.Println("Added value is: ", added)
}
