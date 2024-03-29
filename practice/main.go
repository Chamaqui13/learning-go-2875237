package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[0:len(colors)-1])
	fmt.Println(colors)

	numbers := append(make([]int,5))
	numbers[0] = 10
	numbers[1] = 8
	numbers[2] = 34234
	numbers[3] = 96
	numbers[4] = 786

	fmt.Println(numbers)

	fmt.Println(append(numbers,1))
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
