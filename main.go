package main

import (
	"fmt"
)

func main() {
	elements := []int{1, 2, 3, 4, 5}
	result := leftRotation(elements, 2)
	fmt.Println(result)
}

func leftRotation(array []int, n int) []int {
	var temp []int
	for n >= 1 {
		temp = array[1:]
		temp = append(temp, array[0])
		array = temp
		n--
	}

	return array
}

/*

intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", intSlice)

	last := intSlice[len(intSlice)-1]
	fmt.Printf("Last element: %v\n", last)

	first := intSlice[:0]
	fmt.Printf("First element: %d\n", first)

	remove := intSlice[:len(intSlice)-1]
	fmt.Printf("Remove Last: %v\n", remove)

*/
