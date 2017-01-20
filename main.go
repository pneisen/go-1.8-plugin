package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./plugin.so")

	if err != nil {
		panic(err)
	}

	// Add
	add, err := p.Lookup("Add")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum: %v\n", add.(func(int, int) int)(5, 3))

	// Subtract
	sub, err := p.Lookup("Sub")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Difference: %v\n", sub.(func(int, int) int)(5, 3))
}
