package main

import "fmt"

/*Go does not have a built-in function to remove an element from a slice.

Implement the Remove(nums []int, i int) []int function, which removes the element at index i from the nums slice.
If a non-existent index arrives, then the original slice is returned from the function.

The order of elements can change after an element is removed.*/

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Remove(a, 3))
}

func Remove(nums []int, i int) []int {
	if i <= len(nums)-1 && i >= 0 {
		return append(nums[:i], nums[i+1:]...)
	} else {
		return nums
	}
}
