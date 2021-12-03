package main

import "fmt"

func mirror(arr []int) {
	fmt.Println("До:", arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[len(arr)-1-i]
	}
	fmt.Println("После:", arr, "\n")
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arris := []int{100, 7, 353, 66, 9, 8, 7, 8, 9, 44, -545, -1234567}
	mirror(array)
	mirror(arris)
}
