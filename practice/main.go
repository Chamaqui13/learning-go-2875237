package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Pointer", p)
	fmt.Println("Value of p", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 13
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)

}
