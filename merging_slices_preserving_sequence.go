package main

import fmt

Implement the MergeNumberLists(numberLists ...[]int) []int function,
which takes a variadic list of number slices and merges them into 1, preserving the sequence.

func main() {
	fmt.Println(MergeNumberLists([]int{1, 2, 3}, []int{66, 66, 66}, []int{0, -256, -512}))
}

func MergeNumberLists(numberLists ...[]int) []int {

	res = []int{}

	for _, item = range numberLists {

		res = append(res, item...)

	}

	return res
}
