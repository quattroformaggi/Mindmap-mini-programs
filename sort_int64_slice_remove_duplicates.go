package main

import (
	"fmt"
	"sort"
)

/*Implement the UniqueSortedUserIDs(userIDs []int64) []int64 function, which returns a sorted slice of unique userIDs.

Ideally, processing should be "in-place" (without the allocation of additional memory).*/

func main() {
	a := []int64{9999, 55, 2, 88, 33, 2, 2, 55, 103, 33, 88, 222, 666, 666, -1}
	fmt.Println(UniqueSortedUserIDs(a))
}

func removeDuplicateInt(intSlice []int64) []int64 {
	allKeys := make(map[int64]bool)
	list := []int64{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	UIDs := removeDuplicateInt(userIDs)
	sort.Slice(UIDs, func(i, j int) bool { return UIDs[i] < UIDs[j] })
	return UIDs
}
