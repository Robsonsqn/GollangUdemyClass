package main

import "fmt"

func main() {
	meetShortDeclaration()
}

func meetShortDeclaration() {
	// short declaration
	x := 42

	// normal declaration
	var y = 48

	// using package fmt to emit a message
	fmt.Println("y => ", y, ", x => ", x)
}
