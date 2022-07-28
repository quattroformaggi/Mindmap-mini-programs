package main

import "fmt"

/*Let's imagine that there is a Person structure containing a person's age:

type Person struct {
     age uint8
}

Implement a PersonList type (a slice of Person structures), with a method

(pl PersonList) GetAgePopularity() map[uint8]int

that returns a map, where the key is the age and the value is the number of such ages.*/

func main() {
	var g PersonList = PersonList{{Age: 18}, {Age: 44}, {Age: 88}, {Age: 2}, {Age: 14}, {Age: 18}, {Age: 88}, {Age: 18}}
	fmt.Println(g.GetAgePopularity())
}

type (
	Person     struct{ Age uint8 }
	PersonList []Person
)

func (pl PersonList) GetAgePopularity() map[uint8]int {

	ageCounter := make(map[uint8]int)
	fmt.Println(pl, len(pl))
	for i := 0; i < len(pl); i++ {
		if _, ok := ageCounter[pl[i].Age]; ok {
			ageCounter[pl[i].Age] += 1
		} else {
			ageCounter[pl[i].Age] = 1
		}
	}

	return ageCounter
}
