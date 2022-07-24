package main

import "fmt"

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
