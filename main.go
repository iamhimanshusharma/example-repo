package main

import (
	"fmt"

	b "learnGithub/maths"
)

func main() {
	fmt.Println("Hello Himanshu!")
	fmt.Println("Adding newline")
	fmt.Println("This line will show itself in dev branch.")

	newFunction("hello")

	added := b.Add(5, 6)

	fmt.Println("Added value is: ", added)
}
