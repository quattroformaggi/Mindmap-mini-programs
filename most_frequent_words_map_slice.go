package main

import (
	"fmt"
)

/*Implement the MostPopularWord(words []string) string function, which returns the most frequently occurring word in the slice.

If there are several such words, then the first one is returned.*/

func main() {
	words := []string{"one", "two", "three", "four", "five"}
	fmt.Println(MostPopularWord(words))
}

func MostPopularWord(words []string) string {
	words_map := map[string]int{}
	vmax := 0
	for i := 0; i < len(words); i++ {
		w, found := words_map[words[i]]
		if found {
			words_map[words[i]] = w + 1
		} else {
			words_map[words[i]] = 1
		}
	}

	///

	for _, v := range words_map {
		if v > vmax {
			vmax = v
		}
	}

	///
	if vmax > 0 {
		for key, val := range words_map {
			if val == vmax {
				return key
			}
		}
	}
	return "rip"
}

/*
1. generate a new map "string - number"
2. if the word has not yet been found, add it to the map.
3. if found - find it by key in the map and increment the val.
4. search for the highest val & return it's key

the order of the keys in the map is randomized...*/
