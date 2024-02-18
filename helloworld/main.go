package main

import "fmt"

func main() {
	firstName := "Freddy"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "Shaikh"
}
